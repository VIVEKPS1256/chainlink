package framework

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	CapabilityTypeTrigger   = 0
	CapabilityTypeAction    = 1
	CapabilityTypeConsensus = 2
	CapabilityTypeTarget    = 3
)

type capabilitiesRegistry struct {
	t              *testing.T
	backend        *ethBlockchain
	contract       *kcr.CapabilitiesRegistry
	addr           common.Address
	nodeOperatorID uint32

	ocrid [32]byte
}

func NewCapabilitiesRegistry(ctx context.Context, t *testing.T, backend *ethBlockchain) *capabilitiesRegistry {
	addr, _, contract, err := kcr.DeployCapabilitiesRegistry(backend.transactionOpts, backend)
	require.NoError(t, err)
	backend.Commit()

	_, err = contract.AddNodeOperators(backend.transactionOpts, []kcr.CapabilitiesRegistryNodeOperator{
		{
			Admin: backend.transactionOpts.From,
			Name:  "TEST_NODE_OPERATOR",
		},
	})
	require.NoError(t, err)
	blockHash := backend.Commit()

	logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: nil,
		Topics:    nil,
	})

	require.NoError(t, err)

	recLog, err := contract.ParseNodeOperatorAdded(logs[0])
	require.NoError(t, err)

	nopID := recLog.NodeOperatorId

	return &capabilitiesRegistry{t: t, addr: addr, contract: contract, backend: backend, nodeOperatorID: nopID}
}

func (r *capabilitiesRegistry) getAddress() common.Address {
	return r.addr
}

type capability struct {
	*pb.CapabilityConfig
	kcr.CapabilitiesRegistryCapability
}

func (r *capabilitiesRegistry) setupTriggerDON(triggerDon DonInfo) {

	triggerCapabilityConfig := newCapabilityConfig()
	triggerCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTriggerConfig{
		RemoteTriggerConfig: &pb.RemoteTriggerConfig{
			RegistrationRefresh: durationpb.New(1000 * time.Millisecond),
			RegistrationExpiry:  durationpb.New(60000 * time.Millisecond),
			// F + 1
			MinResponsesToAggregate: uint32(triggerDon.F) + 1,
		},
	}

	streamsTriggerCapability := capability{
		CapabilityConfig: triggerCapabilityConfig,
		CapabilitiesRegistryCapability: kcr.CapabilitiesRegistryCapability{
			LabelledName:   "streams-trigger",
			Version:        "1.0.0",
			CapabilityType: CapabilityTypeTrigger,
		},
	}

	r.setupDON(triggerDon, []capability{streamsTriggerCapability})
}

func (r *capabilitiesRegistry) setupDON(donInfo DonInfo, capabilites []capability) {

	var hashedCapabilityIDs [][32]byte

	for _, c := range capabilites {
		id, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, c.LabelledName, c.Version)
		require.NoError(r.t, err)
		hashedCapabilityIDs = append(hashedCapabilityIDs, id)

		_, err = r.contract.AddCapabilities(r.backend.transactionOpts, []kcr.CapabilitiesRegistryCapability{
			c.CapabilitiesRegistryCapability,
		})
		require.NoError(r.t, err)
	}

	r.backend.Commit()

	nodes := []kcr.CapabilitiesRegistryNodeParams{}
	for _, triggerPeer := range donInfo.peerIDs {
		n, innerErr := peerToNode(r.nodeOperatorID, triggerPeer)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = hashedCapabilityIDs
		nodes = append(nodes, n)
	}

	_, err := r.contract.AddNodes(r.backend.transactionOpts, nodes)
	require.NoError(r.t, err)
	r.backend.Commit()

	ps, err := peers(donInfo.peerIDs)
	require.NoError(r.t, err)

	var capabilityConfigurations []kcr.CapabilitiesRegistryCapabilityConfiguration
	for i, c := range capabilites {

		configBinary, err := proto.Marshal(c.DefaultConfig)
		require.NoError(r.t, err)

		capabilityConfigurations = append(capabilityConfigurations, kcr.CapabilitiesRegistryCapabilityConfiguration{
			CapabilityId: hashedCapabilityIDs[i],
			Config:       configBinary,
		})
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, capabilityConfigurations, true, false, donInfo.F)
	require.NoError(r.t, err)
	r.backend.Commit()

}

func (r *capabilitiesRegistry) setupWorkflowDon(workflowDon DonInfo) {

	ocr := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "offchain_reporting",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeConsensus,
	}
	ocrid, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, ocr.LabelledName, ocr.Version)
	require.NoError(r.t, err)

	r.ocrid = ocrid

	_, err = r.contract.AddCapabilities(r.backend.transactionOpts, []kcr.CapabilitiesRegistryCapability{
		ocr,
	})
	require.NoError(r.t, err)
	r.backend.Commit()

	nodes := []kcr.CapabilitiesRegistryNodeParams{}
	for _, wfPeer := range workflowDon.peerIDs {
		n, innerErr := peerToNode(r.nodeOperatorID, wfPeer)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = [][32]byte{r.ocrid}
		nodes = append(nodes, n)
	}

	_, err = r.contract.AddNodes(r.backend.transactionOpts, nodes)
	require.NoError(r.t, err)

	r.backend.Commit()

}

func (r *capabilitiesRegistry) setupCapabilitiesRegistryContract(workflowDon DonInfo, triggerDon DonInfo,
	targetDon DonInfo) {

	writeChain := kcr.CapabilitiesRegistryCapability{
		LabelledName: "write_geth-testnet",
		Version:      "1.0.0",

		CapabilityType: CapabilityTypeTarget,
	}
	wid, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, writeChain.LabelledName, writeChain.Version)
	if err != nil {
		log.Printf("failed to call GetHashedCapabilityId: %s", err)
	}

	_, err = r.contract.AddCapabilities(r.backend.transactionOpts, []kcr.CapabilitiesRegistryCapability{
		writeChain,
	})
	require.NoError(r.t, err)
	r.backend.Commit()

	nodes := []kcr.CapabilitiesRegistryNodeParams{}

	for _, targetPeer := range targetDon.peerIDs {
		n, innerErr := peerToNode(r.nodeOperatorID, targetPeer)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = [][32]byte{wid}
		nodes = append(nodes, n)
	}

	_, err = r.contract.AddNodes(r.backend.transactionOpts, nodes)
	require.NoError(r.t, err)

	// workflow DON
	ps, err := peers(workflowDon.peerIDs)
	require.NoError(r.t, err)

	cc := newCapabilityConfig()
	ccb, err := proto.Marshal(cc)
	require.NoError(r.t, err)

	cfgs := []kcr.CapabilitiesRegistryCapabilityConfiguration{
		{
			CapabilityId: r.ocrid,
			Config:       ccb,
		},
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, cfgs, false, true, workflowDon.F)
	require.NoError(r.t, err)

	// target DON
	ps, err = peers(targetDon.peerIDs)
	require.NoError(r.t, err)

	targetCapabilityConfig := newCapabilityConfig()

	configWithLimit, err := values.WrapMap(map[string]any{"gasLimit": 500000})
	require.NoError(r.t, err)

	targetCapabilityConfig.DefaultConfig = values.Proto(configWithLimit).GetMapValue()

	targetCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTargetConfig{
		RemoteTargetConfig: &pb.RemoteTargetConfig{
			RequestHashExcludedAttributes: []string{"signed_report.Signatures"},
		},
	}

	remoteTargetConfigBytes, err := proto.Marshal(targetCapabilityConfig)
	require.NoError(r.t, err)

	cfgs = []kcr.CapabilitiesRegistryCapabilityConfiguration{
		{
			CapabilityId: wid,
			Config:       remoteTargetConfigBytes,
		},
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, cfgs, true, false, targetDon.F)
	require.NoError(r.t, err)

	r.backend.Commit()
}

func newCapabilityConfig() *pb.CapabilityConfig {
	return &pb.CapabilityConfig{
		DefaultConfig: values.Proto(values.EmptyMap()).GetMapValue(),
	}
}
