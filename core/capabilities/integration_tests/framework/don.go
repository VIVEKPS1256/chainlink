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
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
)

type DON struct {
	lggr  logger.Logger
	Info  DonInfo
	nodes []*CapabilityNode
	jobs  []*job.Job
}

func (d *DON) Start(ctx context.Context, t *testing.T) {
	for _, node := range d.nodes {
		require.NoError(t, node.Start(testutils.Context(t)))
	}

	for _, node := range d.nodes {
		for _, j := range d.jobs {
			require.NoError(t, node.AddJobV2(ctx, j))
		}
	}
}

func NewDON(ctx context.Context, t *testing.T, lggr logger.Logger, donInfo DonInfo, broker *testAsyncMessageBroker,
	dependentDONs []commoncap.DON, ethBackend *ethBlockchain, capRegistryAddr common.Address) *DON {

	var nodes []*CapabilityNode
	for i, member := range donInfo.Members {
		dispatcher := broker.NewDispatcherForNode(member)
		capabilityRegistry := capabilities.NewRegistry(lggr)

		nodeInfo := commoncap.Node{
			PeerID:         &member,
			WorkflowDON:    donInfo.DON,
			CapabilityDONs: dependentDONs,
		}

		node := startNewNode(ctx, t, lggr.Named(donInfo.name+"-"+strconv.Itoa(i)), nodeInfo, ethBackend, capRegistryAddr, dispatcher,
			testPeerWrapper{peer: testPeer{member}}, capabilityRegistry,
			donInfo.keys[i], nil)

		nodes = append(nodes, node)
	}

	return &DON{lggr: lggr.Named(donInfo.name), nodes: nodes, Info: donInfo}
}

func (d *DON) AddJobV2(j *job.Job) {
	d.jobs = append(d.jobs, j)

}

// Functions for adding non-standard capabilities to a DON
func (d *DON) AddOCR3NonStandardCapability(ctx context.Context, t *testing.T) {
	libocr := NewMockLibOCR(t, d.Info.F, 1*time.Second)
	servicetest.Run(t, libocr)

	for i, node := range d.nodes {
		addOCR3Capability(ctx, t, d.lggr, node.registry, libocr, d.Info.F, d.Info.KeyBundles[i])
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
