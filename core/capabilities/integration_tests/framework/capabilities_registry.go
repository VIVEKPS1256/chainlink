package framework

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/types/known/durationpb"

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
	t        *testing.T
	backend  *ethBlockchain
	contract *kcr.CapabilitiesRegistry
	addr     common.Address
}

func NewCapabilitiesRegistry(t *testing.T, backend *ethBlockchain) *capabilitiesRegistry {
	addr, _, contract, err := kcr.DeployCapabilitiesRegistry(backend.transactionOpts, backend)
	require.NoError(t, err)
	backend.Commit()

	return &capabilitiesRegistry{t: t, addr: addr, contract: contract, backend: backend}
}

func (r *capabilitiesRegistry) getAddress() common.Address {
	return r.addr
}

func (r *capabilitiesRegistry) setupCapabilitiesRegistryContract(ctx context.Context, workflowDon DonInfo, triggerDon DonInfo,
	targetDon DonInfo) {

	streamsTrigger := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "streams-trigger",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeTrigger,
	}
	sid, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, streamsTrigger.LabelledName, streamsTrigger.Version)
	require.NoError(r.t, err)

	writeChain := kcr.CapabilitiesRegistryCapability{
		LabelledName: "write_geth-testnet",
		Version:      "1.0.0",

		CapabilityType: CapabilityTypeTarget,
	}
	wid, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, writeChain.LabelledName, writeChain.Version)
	if err != nil {
		log.Printf("failed to call GetHashedCapabilityId: %s", err)
	}

	ocr := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "offchain_reporting",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeConsensus,
	}
	ocrid, err := r.contract.GetHashedCapabilityId(&bind.CallOpts{}, ocr.LabelledName, ocr.Version)
	require.NoError(r.t, err)

	_, err = r.contract.AddCapabilities(r.backend.transactionOpts, []kcr.CapabilitiesRegistryCapability{
		streamsTrigger,
		writeChain,
		ocr,
	})
	require.NoError(r.t, err)
	r.backend.Commit()

	_, err = r.contract.AddNodeOperators(r.backend.transactionOpts, []kcr.CapabilitiesRegistryNodeOperator{
		{
			Admin: r.backend.transactionOpts.From,
			Name:  "TEST_NODE_OPERATOR",
		},
	})
	require.NoError(r.t, err)
	blockHash := r.backend.Commit()

	logs, err := r.backend.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: nil,
		Topics:    nil,
	})

	require.NoError(r.t, err)

	recLog, err := r.contract.ParseNodeOperatorAdded(logs[0])
	require.NoError(r.t, err)

	nopID := recLog.NodeOperatorId
	nodes := []kcr.CapabilitiesRegistryNodeParams{}
	for _, wfPeer := range workflowDon.peerIDs {
		n, innerErr := peerToNode(nopID, wfPeer)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = [][32]byte{ocrid}
		nodes = append(nodes, n)
	}

	for _, triggerPeer := range triggerDon.peerIDs {
		n, innerErr := peerToNode(nopID, triggerPeer)
		require.NoError(r.t, innerErr)

		n.HashedCapabilityIds = [][32]byte{sid}
		nodes = append(nodes, n)
	}

	for _, targetPeer := range targetDon.peerIDs {
		n, innerErr := peerToNode(nopID, targetPeer)
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
			CapabilityId: ocrid,
			Config:       ccb,
		},
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, cfgs, false, true, workflowDon.F)
	require.NoError(r.t, err)

	// trigger DON
	ps, err = peers(triggerDon.peerIDs)
	require.NoError(r.t, err)

	triggerCapabilityConfig := newCapabilityConfig()
	triggerCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTriggerConfig{
		RemoteTriggerConfig: &pb.RemoteTriggerConfig{
			RegistrationRefresh: durationpb.New(1000 * time.Millisecond),
			RegistrationExpiry:  durationpb.New(60000 * time.Millisecond),
			// F + 1
			MinResponsesToAggregate: uint32(triggerDon.F) + 1,
		},
	}

	configb, err := proto.Marshal(triggerCapabilityConfig)
	require.NoError(r.t, err)

	cfgs = []kcr.CapabilitiesRegistryCapabilityConfiguration{
		{
			CapabilityId: sid,
			Config:       configb,
		},
	}

	_, err = r.contract.AddDON(r.backend.transactionOpts, ps, cfgs, true, false, triggerDon.F)
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
