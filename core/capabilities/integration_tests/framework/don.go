package framework

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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
	remotetypes "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
	"github.com/smartcontractkit/chainlink/v2/core/utils/testutils/heavyweight"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

type capabilityNode struct {
	*cltest.TestApplication
	registry  *capabilities.Registry
	key       ethkey.KeyV2
	KeyBundle ocr2key.KeyBundle
	peerID    peer
	start     func()
}

type DonConfiguration struct {
	commoncap.DON
	name             string
	keys             []ethkey.KeyV2
	KeyBundles       []ocr2key.KeyBundle
	peerIDs          []peer
	acceptsWorkflows bool
}

type DON struct {
	config               DonConfiguration
	lggr                 logger.Logger
	nodes                []*capabilityNode
	jobs                 []*job.Job
	capabilities         []capability
	capabilitiesRegistry *CapabilitiesRegistry

	nodeConfigModifier func(c *chainlink.Config, node *capabilityNode)

	addOCR3NonStandardCapability bool

	triggerFactories []triggerFactory
}

func NewDON(ctx context.Context, t *testing.T, lggr logger.Logger, donConfig DonConfiguration,
	dependentDONs []commoncap.DON, donContext DonContext) *DON {

	don := &DON{lggr: lggr.Named(donConfig.name), config: donConfig, capabilitiesRegistry: donContext.capabilityRegistry}

	for i, member := range donConfig.Members {
		dispatcher := donContext.p2pNetwork.NewDispatcherForNode(member)
		capabilityRegistry := capabilities.NewRegistry(lggr)

		nodeInfo := commoncap.Node{
			PeerID:         &member,
			WorkflowDON:    donConfig.DON,
			CapabilityDONs: dependentDONs,
		}

		cn := &capabilityNode{
			registry:  capabilityRegistry,
			key:       donConfig.keys[i],
			KeyBundle: donConfig.KeyBundles[i],
			peerID:    donConfig.peerIDs[i],
		}
		don.nodes = append(don.nodes, cn)

		cn.start = func() {
			node := startNewNode(ctx, t, lggr.Named(donConfig.name+"-"+strconv.Itoa(i)), nodeInfo, donContext.EthBlockchain,
				donContext.capabilityRegistry.getAddress(), dispatcher,
				peerWrapper{peer: p2pPeer{member}}, capabilityRegistry,
				donConfig.keys[i], func(c *chainlink.Config) {
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
	id := d.capabilitiesRegistry.setupDON(d.config, d.capabilities)
	d.config.DON.ID = uint32(id)
}

func (d *DON) GetID() uint32 {
	if d.config.DON.ID == 0 {
		panic("DON ID not set, call Initialise() first")
	}

	return d.config.ID
}

func (d *DON) GetConfigVersion() uint32 {
	return d.config.ConfigVersion
}

func (d *DON) GetF() uint8 {
	return d.config.F
}

func (d *DON) GetPeerIDs() []peer {
	return d.config.peerIDs
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
		libocr := NewMockLibOCR(t, d.config.F, 1*time.Second)
		servicetest.Run(t, libocr)

		for _, node := range d.nodes {
			addOCR3Capability(ctx, t, d.lggr, node.registry, libocr, d.config.F, node.KeyBundle)
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
			MinResponsesToAggregate: uint32(d.config.F) + 1,
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

type triggerFactory interface {
	GetNewTrigger(t *testing.T) commoncap.TriggerCapability
}

func startNewNode(ctx context.Context,
	t *testing.T, lggr logger.Logger, nodeInfo commoncap.Node,
	ethBlockchain *EthBlockchain, capRegistryAddr common.Address,
	dispatcher remotetypes.Dispatcher,
	peerWrapper p2ptypes.PeerWrapper,
	localCapabilities *capabilities.Registry,
	keyV2 ethkey.KeyV2,
	setupCfg func(c *chainlink.Config),
) *cltest.TestApplication {
	config, _ := heavyweight.FullTestDBV2(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		c.Capabilities.ExternalRegistry.ChainID = ptr(fmt.Sprintf("%d", testutils.SimulatedChainID))
		c.Capabilities.ExternalRegistry.Address = ptr(capRegistryAddr.String())
		c.Capabilities.Peering.V2.Enabled = ptr(true)
		c.Feature.FeedsManager = ptr(false)

		if setupCfg != nil {
			setupCfg(c)
		}
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
	d.nodeConfigModifier = func(c *chainlink.Config, node *capabilityNode) {
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

// NewDonConfigurationParams exists purely to make it obvious in the tests what DON configuration is being used
type NewDonConfigurationParams struct {
	Name             string
	NumNodes         int
	F                uint8
	AcceptsWorkflows bool
}

func NewDonConfiguration(don NewDonConfigurationParams) (DonConfiguration, error) {
	keyBundles, peerIDs, err := getKeyBundlesAndPeerIDs(don.NumNodes)
	if err != nil {
		return DonConfiguration{}, fmt.Errorf("failed to get key bundles and peer IDs: %w", err)
	}

	donPeers := make([]p2ptypes.PeerID, len(peerIDs))
	var donKeys []ethkey.KeyV2
	for i := 0; i < len(peerIDs); i++ {
		peerID := p2ptypes.PeerID{}
		err = peerID.UnmarshalText([]byte(peerIDs[i].PeerID))
		if err != nil {
			return DonConfiguration{}, fmt.Errorf("failed to unmarshal peer ID: %w", err)
		}
		donPeers[i] = peerID
		newKey, err := ethkey.NewV2()
		if err != nil {
			return DonConfiguration{}, fmt.Errorf("failed to create key: %w", err)
		}
		donKeys = append(donKeys, newKey)
	}

	donConfiguration := DonConfiguration{
		DON: commoncap.DON{
			Members:          donPeers,
			F:                don.F,
			ConfigVersion:    1,
			AcceptsWorkflows: don.AcceptsWorkflows,
		},
		name:       don.Name,
		peerIDs:    peerIDs,
		keys:       donKeys,
		KeyBundles: keyBundles,
	}
	return donConfiguration, nil
}

type peer struct {
	PeerID string
	Signer string
}

func peerIDToBytes(peerID string) ([32]byte, error) {
	var peerIDB ragetypes.PeerID
	err := peerIDB.UnmarshalText([]byte(peerID))
	if err != nil {
		return [32]byte{}, err
	}

	return peerIDB, nil
}

func peers(ps []peer) ([][32]byte, error) {
	out := [][32]byte{}
	for _, p := range ps {
		b, err := peerIDToBytes(p.PeerID)
		if err != nil {
			return nil, err
		}

		out = append(out, b)
	}

	return out, nil
}

func peerToNode(nopID uint32, p peer) (kcr.CapabilitiesRegistryNodeParams, error) {
	peerIDB, err := peerIDToBytes(p.PeerID)
	if err != nil {
		return kcr.CapabilitiesRegistryNodeParams{}, fmt.Errorf("failed to convert peerID: %w", err)
	}

	sig := strings.TrimPrefix(p.Signer, "0x")
	signerB, err := hex.DecodeString(sig)
	if err != nil {
		return kcr.CapabilitiesRegistryNodeParams{}, fmt.Errorf("failed to convert signer: %w", err)
	}

	var sigb [32]byte
	copy(sigb[:], signerB)

	return kcr.CapabilitiesRegistryNodeParams{
		NodeOperatorId: nopID,
		P2pId:          peerIDB,
		Signer:         sigb,
	}, nil
}
