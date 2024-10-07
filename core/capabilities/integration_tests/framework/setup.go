package framework

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/feeds_consumer"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	coretypes "github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities"
	remotetypes "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
	"github.com/smartcontractkit/chainlink/v2/core/utils/testutils/heavyweight"
)

const (
	// As a  default set the logging to info otherwise 10s/100s of MB of logs are created on each test run
	TestLogLevel = zapcore.InfoLevel
)

var (
	workflowName    = "abcdef0123"
	workflowOwnerID = "0100000000000000000000000000000000000001"
)

type DonInfo struct {
	commoncap.DON
	keys       []ethkey.KeyV2
	KeyBundles []ocr2key.KeyBundle
	peerIDs    []peer
}

func SetupDons(ctx context.Context, t *testing.T, workflowDonInfo DonInfo, triggerDonInfo DonInfo, targetDonInfo DonInfo,
	addWorkFlowJob func(t *testing.T, nodes []*cltest.TestApplication,
		workflowName string,
		workflowOwner string,
		consumerAddr common.Address)) (*feeds_consumer.KeystoneFeedsConsumer, *ReportsSink) {
	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(TestLogLevel)

	ethBlockchain := setupBlockchain(t, 1000, 1*time.Second)
	capabilitiesRegistryAddr := setupCapabilitiesRegistryContract(ctx, t, workflowDonInfo, triggerDonInfo, targetDonInfo, ethBlockchain)
	forwarderAddr, _ := setupForwarderContract(t, workflowDonInfo, ethBlockchain)
	consumerAddr, consumer := setupConsumerContract(t, ethBlockchain, forwarderAddr, workflowOwnerID, workflowName)
	libocr := NewMockLibOCR(t, workflowDonInfo.F, 1*time.Second)
	msgBroker := NewTestAsyncMessageBroker(t, 1000)

	servicetest.Run(t, msgBroker)
	servicetest.Run(t, ethBlockchain)
	servicetest.Run(t, libocr)

	sink := NewReportsSink()

	createTriggerDON(ctx, t, lggr, sink, triggerDonInfo, msgBroker, ethBlockchain, capabilitiesRegistryAddr)

	createTargetDON(ctx, t, lggr, targetDonInfo, msgBroker, ethBlockchain, capabilitiesRegistryAddr, forwarderAddr)

	workflowDonNodes := createWorkflowDON(ctx, t, lggr, workflowDonInfo, msgBroker, libocr,
		[]commoncap.DON{triggerDonInfo.DON, targetDonInfo.DON},
		ethBlockchain, capabilitiesRegistryAddr)

	addWorkFlowJob(t, workflowDonNodes, workflowName, workflowOwnerID, consumerAddr)

	servicetest.Run(t, sink)
	return consumer, sink
}

func createWorkflowDON(ctx context.Context, t *testing.T, lggr logger.Logger, workflowDon DonInfo, broker *testAsyncMessageBroker, libocr *MockLibOCR,
	capabilityDONs []commoncap.DON, simulatedEthBlockchain *ethBlockchain, capRegistryAddr common.Address) []*cltest.TestApplication {
	var workflowNodes []*cltest.TestApplication
	for i, workflowPeer := range workflowDon.Members {
		workflowPeerDispatcher := broker.NewDispatcherForNode(workflowPeer)
		capabilityRegistry := capabilities.NewRegistry(lggr)

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
			F: int(workflowDon.F),
		}
		plugin, _, err := pluginFactory.NewReportingPlugin(repConfig)
		require.NoError(t, err)

		transmitter := ocr3.NewContractTransmitter(lggr, capabilityRegistry, "")

		libocr.AddNode(plugin, transmitter, workflowDon.KeyBundles[i])

		nodeInfo := commoncap.Node{
			PeerID:         &workflowPeer,
			WorkflowDON:    workflowDon.DON,
			CapabilityDONs: capabilityDONs,
		}

		workflowNode := startNewNode(ctx, t, lggr.Named("Workflow-"+strconv.Itoa(i)), nodeInfo, simulatedEthBlockchain, capRegistryAddr, workflowPeerDispatcher,
			testPeerWrapper{peer: testPeer{workflowPeer}}, capabilityRegistry, nil,
			workflowDon.keys[i])

		require.NoError(t, workflowNode.Start(testutils.Context(t)))
		workflowNodes = append(workflowNodes, workflowNode)
	}
	return workflowNodes
}

func createTargetDON(ctx context.Context, t *testing.T, lggr logger.Logger, targetDon DonInfo, broker *testAsyncMessageBroker,
	ethBlockchain *ethBlockchain, capRegistryAddr common.Address, forwarderAddr common.Address) []*cltest.TestApplication {
	var targetNodes []*cltest.TestApplication
	for i, targetPeer := range targetDon.Members {
		targetPeerDispatcher := broker.NewDispatcherForNode(targetPeer)
		nodeInfo := commoncap.Node{
			PeerID: &targetPeer,
		}

		capabilityRegistry := capabilities.NewRegistry(lggr)

		targetNode := startNewNode(ctx, t, lggr.Named("Target-"+strconv.Itoa(i)), nodeInfo, ethBlockchain, capRegistryAddr, targetPeerDispatcher,
			testPeerWrapper{peer: testPeer{targetPeer}}, capabilityRegistry, &forwarderAddr,
			targetDon.keys[i])

		require.NoError(t, targetNode.Start(testutils.Context(t)))
		targetNodes = append(targetNodes, targetNode)
	}
	return targetNodes
}

func createTriggerDON(ctx context.Context, t *testing.T, lggr logger.Logger, reportsSink *ReportsSink, triggerDon DonInfo,
	broker *testAsyncMessageBroker, ethBackend *ethBlockchain, capRegistryAddr common.Address) []*cltest.TestApplication {
	var triggerNodes []*cltest.TestApplication
	for i, triggerPeer := range triggerDon.Members {
		triggerPeerDispatcher := broker.NewDispatcherForNode(triggerPeer)
		nodeInfo := commoncap.Node{
			PeerID: &triggerPeer,
		}

		capabilityRegistry := capabilities.NewRegistry(lggr)
		trigger := reportsSink.GetNewTrigger(t)
		err := capabilityRegistry.Add(ctx, trigger)
		require.NoError(t, err)

		triggerNode := startNewNode(ctx, t, lggr.Named("Trigger-"+strconv.Itoa(i)), nodeInfo, ethBackend, capRegistryAddr, triggerPeerDispatcher,
			testPeerWrapper{peer: testPeer{triggerPeer}}, capabilityRegistry, nil,
			triggerDon.keys[i])

		require.NoError(t, triggerNode.Start(testutils.Context(t)))
		triggerNodes = append(triggerNodes, triggerNode)
	}
	return triggerNodes
}

func startNewNode(ctx context.Context,
	t *testing.T, lggr logger.Logger, nodeInfo commoncap.Node,
	ethBlockchain *ethBlockchain, capRegistryAddr common.Address,
	dispatcher remotetypes.Dispatcher,
	peerWrapper p2ptypes.PeerWrapper,
	localCapabilities *capabilities.Registry,
	forwarderAddress *common.Address,
	keyV2 ethkey.KeyV2,
) *cltest.TestApplication {
	config, _ := heavyweight.FullTestDBV2(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		c.Capabilities.ExternalRegistry.ChainID = ptr(fmt.Sprintf("%d", testutils.SimulatedChainID))
		c.Capabilities.ExternalRegistry.Address = ptr(capRegistryAddr.String())
		c.Capabilities.Peering.V2.Enabled = ptr(true)

		if forwarderAddress != nil {
			eip55Address := types.EIP55AddressFromAddress(*forwarderAddress)
			c.EVM[0].Chain.Workflow.ForwarderAddress = &eip55Address
			c.EVM[0].Chain.Workflow.FromAddress = &keyV2.EIP55Address
		}

		c.Feature.FeedsManager = ptr(false)
	})

	n, err := ethBlockchain.NonceAt(ctx, ethBlockchain.transactionOpts.From, nil)
	require.NoError(t, err)

	tx := cltest.NewLegacyTransaction(
		n, keyV2.Address,
		assets.Ether(1).ToInt(),
		21000,
		assets.GWei(1).ToInt(),
		nil)
	signedTx, err := ethBlockchain.transactionOpts.Signer(ethBlockchain.transactionOpts.From, tx)
	require.NoError(t, err)
	err = ethBlockchain.SendTransaction(ctx, signedTx)
	require.NoError(t, err)
	ethBlockchain.Commit()

	return cltest.NewApplicationWithConfigV2AndKeyOnSimulatedBlockchain(t, config, ethBlockchain.SimulatedBackend, nodeInfo,
		dispatcher, peerWrapper, localCapabilities, keyV2, lggr)
}

type Don struct {
	ID       uint32
	NumNodes int
	F        uint8
}

func CreateDonInfo(t *testing.T, don Don) DonInfo {
	keyBundles, peerIDs := getKeyBundlesAndPeerIDs(t, don.NumNodes)

	donPeers := make([]p2ptypes.PeerID, len(peerIDs))
	var donKeys []ethkey.KeyV2
	for i := 0; i < len(peerIDs); i++ {
		peerID := p2ptypes.PeerID{}
		require.NoError(t, peerID.UnmarshalText([]byte(peerIDs[i].PeerID)))
		donPeers[i] = peerID
		newKey, err := ethkey.NewV2()
		require.NoError(t, err)
		donKeys = append(donKeys, newKey)
	}

	triggerDonInfo := DonInfo{
		DON: commoncap.DON{
			ID:            don.ID,
			Members:       donPeers,
			F:             don.F,
			ConfigVersion: 1,
		},
		peerIDs:    peerIDs,
		keys:       donKeys,
		KeyBundles: keyBundles,
	}
	return triggerDonInfo
}

func getKeyBundlesAndPeerIDs(t *testing.T, numNodes int) ([]ocr2key.KeyBundle, []peer) {
	var keyBundles []ocr2key.KeyBundle
	var donPeerIDs []peer
	for i := 0; i < numNodes; i++ {
		peerID := NewPeerID()

		keyBundle, err := ocr2key.New(chaintype.EVM)
		require.NoError(t, err)
		keyBundles = append(keyBundles, keyBundle)

		pk := keyBundle.PublicKey()

		p := peer{
			PeerID: peerID,
			Signer: fmt.Sprintf("0x%x", pk),
		}

		donPeerIDs = append(donPeerIDs, p)
	}
	return keyBundles, donPeerIDs
}

type testPeerWrapper struct {
	peer testPeer
}

func (t testPeerWrapper) Start(ctx context.Context) error {
	return nil
}

func (t testPeerWrapper) Close() error {
	return nil
}

func (t testPeerWrapper) Ready() error {
	return nil
}

func (t testPeerWrapper) HealthReport() map[string]error {
	return nil
}

func (t testPeerWrapper) Name() string {
	return "testPeerWrapper"
}

func (t testPeerWrapper) GetPeer() p2ptypes.Peer {
	return t.peer
}

type testPeer struct {
	id p2ptypes.PeerID
}

func (t testPeer) Start(ctx context.Context) error {
	return nil
}

func (t testPeer) Close() error {
	return nil
}

func (t testPeer) Ready() error {
	return nil
}

func (t testPeer) HealthReport() map[string]error {
	return nil
}

func (t testPeer) Name() string {
	return "testPeer"
}

func (t testPeer) ID() p2ptypes.PeerID {
	return t.id
}

func (t testPeer) UpdateConnections(peers map[p2ptypes.PeerID]p2ptypes.StreamConfig) error {
	return nil
}

func (t testPeer) Send(peerID p2ptypes.PeerID, msg []byte) error {
	return nil
}

func (t testPeer) Receive() <-chan p2ptypes.Message {
	return nil
}

func NewPeerID() string {
	var privKey [32]byte
	_, err := rand.Read(privKey[:])
	if err != nil {
		panic(err)
	}

	peerID := append(libp2pMagic(), privKey[:]...)

	return base58.Encode(peerID[:])
}

func libp2pMagic() []byte {
	return []byte{0x00, 0x24, 0x08, 0x01, 0x12, 0x20}
}

func ptr[T any](t T) *T { return &t }
