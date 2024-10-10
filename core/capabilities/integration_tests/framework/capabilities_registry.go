package framework

import (
	"context"
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
	donCapabilityConfig *pb.CapabilityConfig
	registryConfig      kcr.CapabilitiesRegistryCapability
}

func (r *capabilitiesRegistry) setupTargetDon(targetDon DonInfo) {

	writeChain := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "write_geth-testnet",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeTarget,
	}

	targetCapabilityConfig := newCapabilityConfig()

	configWithLimit, err := values.WrapMap(map[string]any{"gasLimit": 500000})
	require.NoError(r.t, err)

	targetCapabilityConfig.DefaultConfig = values.Proto(configWithLimit).GetMapValue()

	targetCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTargetConfig{
		RemoteTargetConfig: &pb.RemoteTargetConfig{
			RequestHashExcludedAttributes: []string{"signed_report.Signatures"},
		},
	}

	r.setupDON(targetDon, []capability{{
		donCapabilityConfig: targetCapabilityConfig,
		registryConfig:      writeChain,
	}})
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
		donCapabilityConfig: triggerCapabilityConfig,
		registryConfig: kcr.CapabilitiesRegistryCapability{
			LabelledName:   "streams-trigger",
			Version:        "1.0.0",
			CapabilityType: CapabilityTypeTrigger,
		},
	}

	r.setupDON(triggerDon, []capability{streamsTriggerCapability})
}

func (r *capabilitiesRegistry) setupWorkflowDon(workflowDon DonInfo) {

	ocr := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "offchain_reporting",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeConsensus,
	}

	r.setupDON(workflowDon, []capability{{
		donCapabilityConfig: newCapabilityConfig(),
		registryConfig:      ocr,
	}})
}

func (r *capabilitiesRegistry) setupDON(donInfo DonInfo, capabilites []capability) {

	var hashedCapabilityIDs [][32]byte

	for _, c := range capabilites {
		id, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, c.registryConfig.LabelledName, c.registryConfig.Version)
		require.NoError(r.t, err)
		hashedCapabilityIDs = append(hashedCapabilityIDs, id)
	}

	var registryCapabilities []kcr.CapabilitiesRegistryCapability
	for _, c := range capabilites {
		registryCapabilities = append(registryCapabilities, c.registryConfig)
	}

	_, err := r.contract.AddCapabilities(r.backend.transactionOpts, registryCapabilities)
	require.NoError(r.t, err)

	r.backend.Commit()

	nodes := []kcr.CapabilitiesRegistryNodeParams{}
	for _, peerID := range donInfo.peerIDs {
		n, innerErr := peerToNode(r.nodeOperatorID, peerID)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = hashedCapabilityIDs
		nodes = append(nodes, n)
	}

	_, err = r.contract.AddNodes(r.backend.transactionOpts, nodes)
	require.NoError(r.t, err)
	r.backend.Commit()

	ps, err := peers(donInfo.peerIDs)
	require.NoError(r.t, err)

	var capabilityConfigurations []kcr.CapabilitiesRegistryCapabilityConfiguration
	for i, c := range capabilites {

		configBinary, err := proto.Marshal(c.donCapabilityConfig)
		require.NoError(r.t, err)

		capabilityConfigurations = append(capabilityConfigurations, kcr.CapabilitiesRegistryCapabilityConfiguration{
			CapabilityId: hashedCapabilityIDs[i],
			Config:       configBinary,
		})
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, capabilityConfigurations, true, donInfo.AcceptsWorkflows, donInfo.F)
	require.NoError(r.t, err)
	r.backend.Commit()

}

func newCapabilityConfig() *pb.CapabilityConfig {
	return &pb.CapabilityConfig{
		DefaultConfig: values.Proto(values.EmptyMap()).GetMapValue(),
	}
}
