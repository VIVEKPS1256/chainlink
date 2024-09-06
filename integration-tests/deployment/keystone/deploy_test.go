package keystone_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"
	"go.uber.org/zap/zapcore"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/clo"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/clo/models"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/keystone"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var (
	testNops = []kcr.CapabilitiesRegistryNodeOperator{
		{
			Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f22D6"),
			Name:  "NOP_00",
		},
		{
			Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f2200"),
			Name:  "NOP_01",
		},
		{
			Admin: common.HexToAddress("0x11dfBF967A8ec4C29Fe26aF2a33Eb485d02f22D6"),
			Name:  "NOP_02",
		},
		{
			Admin: common.HexToAddress("0x6CdfBF967A8ec4C29Fe26aF2a33Eb485d02f2222"),
			Name:  "NOP_03",
		},
	}
)

func TestDeploy(t *testing.T) {
	lggr := logger.TestLogger(t)
	t.Run("memory environment", func(t *testing.T) {
		multDonCfg := memory.MemoryEnvironmentMultiDonConfig{
			Configs: make(map[string]memory.MemoryEnvironmentConfig),
		}
		wfEnvCfg := memory.MemoryEnvironmentConfig{
			Bootstraps: 1,
			Chains:     1,
			Nodes:      4,
		}
		multDonCfg.Configs[keystone.WFDonName] = wfEnvCfg

		targetEnvCfg := memory.MemoryEnvironmentConfig{
			Bootstraps: 1,
			Chains:     4,
			Nodes:      4,
		}
		multDonCfg.Configs[keystone.TargetDonName] = targetEnvCfg

		e := memory.NewMultiDonMemoryEnvironment(t, lggr, zapcore.InfoLevel, multDonCfg)

		var nodeToNop = make(map[string]kcr.CapabilitiesRegistryNodeOperator) //node -> nop
		// assign nops to nodes
		for _, env := range e.DonToEnv {
			for i, nodeID := range env.NodeIDs {
				idx := i % len(testNops)
				nop := testNops[idx]
				nodeToNop[nodeID] = nop
			}
		}

		var donsToDeploy = map[string][]kcr.CapabilitiesRegistryCapability{
			keystone.WFDonName:     []kcr.CapabilitiesRegistryCapability{keystone.OCR3Cap},
			keystone.TargetDonName: []kcr.CapabilitiesRegistryCapability{keystone.WriteChainCap},
		}

		ctx := context.Background()
		// Deploy all the Keystone contracts.
		homeChain := e.Get(keystone.WFDonName).AllChainSelectors()[0]
		deployReq := keystone.DeployRequest{
			RegistryChain:     homeChain,
			Menv:              e,
			DonToCapabilities: donsToDeploy,
			NodeIDToNop:       nodeToNop,
		}

		deployResp, err := keystone.Deploy(ctx, lggr, deployReq)
		require.NoError(t, err)
		ad := deployResp.Changeset.AddressBook
		addrs, err := ad.Addresses()
		require.NoError(t, err)
		lggr.Infow("Deployed Keystone contracts", "address book", addrs)

		// all contracts on home chain
		homeChainAddrs, err := ad.AddressesForChain(homeChain)
		require.NoError(t, err)
		require.Len(t, homeChainAddrs, 3)
		// only forwarder on non-home chain
		for _, chain := range e.Get(keystone.TargetDonName).AllChainSelectors() {
			chainAddrs, err := ad.AddressesForChain(chain)
			require.NoError(t, err)
			if chain != homeChain {
				require.Len(t, chainAddrs, 1)
			} else {
				require.Len(t, chainAddrs, 3)
			}
			containsForwarder := false
			for _, tv := range chainAddrs {
				if tv.Type == keystone.KeystoneForwarder {
					containsForwarder = true
					break
				}
			}
			require.True(t, containsForwarder, "no forwarder found in %v on chain %d for target don", chainAddrs, chain)
		}

		req := &keystone.GetContractSetsRequest{
			Chains:      e.Chains(),
			AddressBook: ad,
		}

		contractSetsResp, err := keystone.GetContractSets(lggr, req)
		require.NoError(t, err)
		require.Len(t, contractSetsResp.ContractSets, 4)
		// check the registry
		regChainContracts, ok := contractSetsResp.ContractSets[homeChain]
		require.True(t, ok)
		gotRegistry := regChainContracts.CapabilitiesRegistry
		require.NotNil(t, gotRegistry)
		// contract reads
		gotDons, err := gotRegistry.GetDONs(nil)
		require.NoError(t, err)
		assert.Len(t, gotDons, len(e.DonToEnv))
		for don, id := range deployResp.DonToId {
			// id starts at 1 in the contract
			gdon := gotDons[id-1]
			cfg, ok := multDonCfg.Configs[don]
			require.True(t, ok, "no config for don %s", don)
			assert.Equal(t, cfg.Nodes/3, int(gdon.F))
			assert.Len(t, gdon.NodeP2PIds, cfg.Nodes)
			assert.Equal(t, don == keystone.WFDonName, gdon.AcceptsWorkflows, "don %s, %d has wrong AcceptsWorkflows", don, id)
		}
	})

	t.Run("memory chains clo offchain", func(t *testing.T) {
		wfNops := loadTestNops(t, "testdata/workflow_nodes.json")
		cwNops := loadTestNops(t, "testdata/chain_writer_nodes.json")

		wfDon := clo.NewDonEnvWithMemoryChains(t, clo.DonEnvConfig{
			DonName: keystone.WFDonName,
			Nops:    wfNops,
			Logger:  lggr,
		})

		cwDon := clo.NewDonEnvWithMemoryChains(t, clo.DonEnvConfig{
			DonName: keystone.TargetDonName,
			Nops:    cwNops,
			Logger:  lggr,
		})

		var donToEnv = map[string]*deployment.Environment{
			keystone.WFDonName:     wfDon,
			keystone.TargetDonName: cwDon,
		}

		menv := clo.NewMultiDonEnvironment(lggr, donToEnv)

		var nodeToNop = make(map[string]kcr.CapabilitiesRegistryNodeOperator) //node -> nop
		// assign nops to nodes
		for _, env := range menv.DonToEnv {
			for i, nodeID := range env.NodeIDs {
				idx := i % len(testNops)
				nop := testNops[idx]
				nodeToNop[nodeID] = nop
			}
		}

		var donsToDeploy = map[string][]kcr.CapabilitiesRegistryCapability{
			keystone.WFDonName:     []kcr.CapabilitiesRegistryCapability{keystone.OCR3Cap},
			keystone.TargetDonName: []kcr.CapabilitiesRegistryCapability{keystone.WriteChainCap},
		}

		ctx := context.Background()

		// sepolia
		homeChainSel, err := chainsel.SelectorFromChainId(11155111)
		require.NoError(t, err)
		deployReq := keystone.DeployRequest{
			RegistryChain:     homeChainSel,
			Menv:              menv,
			DonToCapabilities: donsToDeploy,
			NodeIDToNop:       nodeToNop,
		}

		deployResp, err := keystone.Deploy(ctx, lggr, deployReq)
		require.NoError(t, err)
		ad := deployResp.Changeset.AddressBook
		addrs, err := ad.Addresses()
		require.NoError(t, err)
		lggr.Infow("Deployed Keystone contracts", "address book", addrs)

	})
}

func loadTestNops(t *testing.T, pth string) []*models.NodeOperator {
	f, err := os.ReadFile(pth)
	require.NoError(t, err)
	var nops []*models.NodeOperator
	require.NoError(t, json.Unmarshal(f, &nops))
	return nops
}
