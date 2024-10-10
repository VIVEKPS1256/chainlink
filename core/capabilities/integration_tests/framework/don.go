package framework

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	coretypes "github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
)

type CapabilityNode struct {
	*cltest.TestApplication
	registry  *capabilities.Registry
	key       ethkey.KeyV2
	KeyBundle ocr2key.KeyBundle
	peerID    peer
	start     func()
}

type DON struct {
	info                 DonInfo
	lggr                 logger.Logger
	name                 string
	nodes                []*CapabilityNode
	jobs                 []*job.Job
	capabilities         []capability
	capabilitiesRegistry *CapabilitiesRegistry

	nodeConfigModifier func(c *chainlink.Config, node *CapabilityNode)

	addOCR3NonStandardCapability bool

	triggerFactories []triggerFactory
}

func NewDON(ctx context.Context, t *testing.T, lggr logger.Logger, donInfo DonInfo, broker *testAsyncMessageBroker,
	dependentDONs []commoncap.DON, ethBackend *ethBlockchain, capabilitiesRegistry *CapabilitiesRegistry) *DON {

	don := &DON{lggr: lggr.Named(donInfo.name), name: donInfo.name, info: donInfo, capabilitiesRegistry: capabilitiesRegistry}

	for i, member := range donInfo.Members {
		dispatcher := broker.NewDispatcherForNode(member)
		capabilityRegistry := capabilities.NewRegistry(lggr)

		nodeInfo := commoncap.Node{
			PeerID:         &member,
			WorkflowDON:    donInfo.DON,
			CapabilityDONs: dependentDONs,
		}

		cn := &CapabilityNode{
			registry:  capabilityRegistry,
			key:       donInfo.keys[i],
			KeyBundle: donInfo.KeyBundles[i],
			peerID:    donInfo.peerIDs[i],
		}
		don.nodes = append(don.nodes, cn)

		cn.start = func() {
			node := startNewNode(ctx, t, lggr.Named(donInfo.name+"-"+strconv.Itoa(i)), nodeInfo, ethBackend, capabilitiesRegistry.getAddress(), dispatcher,
				testPeerWrapper{peer: testPeer{member}}, capabilityRegistry,
				donInfo.keys[i], func(c *chainlink.Config) {
					if don.nodeConfigModifier != nil {
						don.nodeConfigModifier(c, cn)
					}
				})

			require.NoError(t, node.Start(testutils.Context(t)))
			cn.TestApplication = node
		}
	}

	return don
}

func (d *DON) Initialise() {
	d.capabilitiesRegistry.setupDON(d.info, d.capabilities)
}

func (d *DON) Start(ctx context.Context, t *testing.T) {
	for _, triggerFactory := range d.triggerFactories {
		for _, node := range d.nodes {
			trigger := triggerFactory.GetNewTrigger(t)
			err := node.registry.Add(ctx, trigger)
			require.NoError(t, err)
		}
	}

	for _, node := range d.nodes {
		node.start()
	}

	if d.addOCR3NonStandardCapability {
		libocr := NewMockLibOCR(t, d.info.F, 1*time.Second)
		servicetest.Run(t, libocr)

		for _, node := range d.nodes {
			addOCR3Capability(ctx, t, d.lggr, node.registry, libocr, d.info.F, node.KeyBundle)
		}
	}

	for _, node := range d.nodes {
		for _, j := range d.jobs {
			require.NoError(t, node.AddJobV2(ctx, j))
		}
	}
}

// Is this streams specific, can it made capabilty type agnostic?
func (d *DON) AddTriggerCapability(triggerFactory triggerFactory) {
	d.triggerFactories = append(d.triggerFactories, triggerFactory)

	triggerCapabilityConfig := newCapabilityConfig()
	triggerCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTriggerConfig{
		RemoteTriggerConfig: &pb.RemoteTriggerConfig{
			RegistrationRefresh: durationpb.New(1000 * time.Millisecond),
			RegistrationExpiry:  durationpb.New(60000 * time.Millisecond),
			// F + 1
			MinResponsesToAggregate: uint32(d.info.F) + 1,
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

	d.capabilities = append(d.capabilities, streamsTriggerCapability)
}

func (d *DON) AddJob(j *job.Job) {
	d.jobs = append(d.jobs, j)
}

// Functions for adding non-standard capabilities to a DON, deliberately verbose
func (d *DON) AddOCR3NonStandardCapability(ctx context.Context, t *testing.T) {
	d.addOCR3NonStandardCapability = true

	ocr := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "offchain_reporting",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeConsensus,
	}

	d.capabilities = append(d.capabilities, capability{
		donCapabilityConfig: newCapabilityConfig(),
		registryConfig:      ocr,
	})
}

// TODO support chaining these modifiers?
func (d *DON) AddEthereumWriteTargetNonStandardCapability(forwarderAddr common.Address) error {
	d.nodeConfigModifier = func(c *chainlink.Config, node *CapabilityNode) {
		eip55Address := types.EIP55AddressFromAddress(forwarderAddr)
		c.EVM[0].Chain.Workflow.ForwarderAddress = &eip55Address
		c.EVM[0].Chain.Workflow.FromAddress = &node.key.EIP55Address
	}

	writeChain := kcr.CapabilitiesRegistryCapability{
		LabelledName:   "write_geth-testnet",
		Version:        "1.0.0",
		CapabilityType: CapabilityTypeTarget,
	}

	targetCapabilityConfig := newCapabilityConfig()

	configWithLimit, err := values.WrapMap(map[string]any{"gasLimit": 500000})
	if err != nil {
		return fmt.Errorf("failed to wrap map: %v", err)
	}

	targetCapabilityConfig.DefaultConfig = values.Proto(configWithLimit).GetMapValue()

	targetCapabilityConfig.RemoteConfig = &pb.CapabilityConfig_RemoteTargetConfig{
		RemoteTargetConfig: &pb.RemoteTargetConfig{
			RequestHashExcludedAttributes: []string{"signed_report.Signatures"},
		},
	}

	d.capabilities = append(d.capabilities, capability{
		donCapabilityConfig: targetCapabilityConfig,
		registryConfig:      writeChain,
	})

	return nil
}

func addOCR3Capability(ctx context.Context, t *testing.T, lggr logger.Logger, capabilityRegistry *capabilities.Registry,
	libocr *MockLibOCR, donF uint8, ocr2KeyBundle ocr2key.KeyBundle) {
	requestTimeout := 10 * time.Minute
	cfg := ocr3.Config{
		Logger:            lggr,
		EncoderFactory:    capabilities.NewEncoder,
		AggregatorFactory: capabilities.NewAggregator,
		RequestTimeout:    &requestTimeout,
	}

	ocr3Capability := ocr3.NewOCR3(cfg)
	servicetest.Run(t, ocr3Capability)

	pluginCfg := coretypes.ReportingPluginServiceConfig{}
	pluginFactory, err := ocr3Capability.NewReportingPluginFactory(ctx, pluginCfg, nil,
		nil, nil, nil, capabilityRegistry, nil, nil)
	require.NoError(t, err)

	repConfig := ocr3types.ReportingPluginConfig{
		F: int(donF),
	}
	plugin, _, err := pluginFactory.NewReportingPlugin(repConfig)
	require.NoError(t, err)

	transmitter := ocr3.NewContractTransmitter(lggr, capabilityRegistry, "")

	libocr.AddNode(plugin, transmitter, ocr2KeyBundle)
}
