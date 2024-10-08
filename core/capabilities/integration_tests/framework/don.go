package framework

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	coretypes "github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
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
}

type DON struct {
	commoncap.DON
	lggr  logger.Logger
	name  string
	nodes []*CapabilityNode
	jobs  []*job.Job

	nodeConfigModifier func(c *chainlink.Config, node *CapabilityNode)

	addOCR3NonStandardCapability bool

	startNodeFns []func() *CapabilityNode
}

func NewDON(ctx context.Context, t *testing.T, lggr logger.Logger, donInfo DonInfo, broker *testAsyncMessageBroker,
	dependentDONs []commoncap.DON, ethBackend *ethBlockchain, capRegistryAddr common.Address) *DON {

	don := &DON{lggr: lggr.Named(donInfo.name), name: donInfo.name, DON: donInfo.DON}

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

		don.startNodeFns = append(don.startNodeFns, func() *CapabilityNode {

			node := startNewNode(ctx, t, lggr.Named(donInfo.name+"-"+strconv.Itoa(i)), nodeInfo, ethBackend, capRegistryAddr, dispatcher,
				testPeerWrapper{peer: testPeer{member}}, capabilityRegistry,
				donInfo.keys[i], func(c *chainlink.Config) {
					if don.nodeConfigModifier != nil {
						don.nodeConfigModifier(c, cn)
					}
				})

			require.NoError(t, node.Start(testutils.Context(t)))
			cn.TestApplication = node
			return cn
		})
	}

	return don
}

func (d *DON) Start(ctx context.Context, t *testing.T) {
	for _, startFn := range d.startNodeFns {
		d.nodes = append(d.nodes, startFn())
	}

	if d.addOCR3NonStandardCapability {
		libocr := NewMockLibOCR(t, d.F, 1*time.Second)
		servicetest.Run(t, libocr)

		for _, node := range d.nodes {
			addOCR3Capability(ctx, t, d.lggr, node.registry, libocr, d.F, node.KeyBundle)
		}
	}

	for _, node := range d.nodes {
		for _, j := range d.jobs {
			require.NoError(t, node.AddJobV2(ctx, j))
		}
	}
}

func (d *DON) AddJobV2(j *job.Job) {
	d.jobs = append(d.jobs, j)
}

// Functions for adding non-standard capabilities to a DON, deliberately verbose
func (d *DON) AddOCR3NonStandardCapability(ctx context.Context, t *testing.T) {
	d.addOCR3NonStandardCapability = true
}

// TODO support chaining these modifiers?
func (d *DON) AddEthereumWriteTargetNonStandardCapability(forwarderAddr common.Address) {
	d.nodeConfigModifier = func(c *chainlink.Config, node *CapabilityNode) {
		eip55Address := types.EIP55AddressFromAddress(forwarderAddr)
		c.EVM[0].Chain.Workflow.ForwarderAddress = &eip55Address
		c.EVM[0].Chain.Workflow.FromAddress = &node.key.EIP55Address
	}
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
