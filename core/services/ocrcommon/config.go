package ocrcommon

import (
	"time"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/libocr/commontypes"
)

type Config interface {
	LogSQL() bool
	EvmGasLimitDefault() uint64
	JobPipelineResultWriteQueueDepth() uint64
	OCRBlockchainTimeout() time.Duration
	OCRContractConfirmations() uint16
	OCRContractPollInterval() time.Duration
	OCRContractSubscribeInterval() time.Duration
	OCRContractTransmitterTransmitTimeout() time.Duration
	OCRDatabaseTimeout() time.Duration
	OCRDefaultTransactionQueueDepth() uint32
	OCRKeyBundleID() (string, error)
	OCRObservationGracePeriod() time.Duration
	OCRObservationTimeout() time.Duration
	OCRTraceLogging() bool
	OCRTransmitterAddress() (ethkey.EIP55Address, error)
	P2PBootstrapPeers() ([]string, error)
	P2PPeerID() p2pkey.PeerID
	P2PV2Bootstrappers() []commontypes.BootstrapperLocator
	FlagsContractAddress() string
	ChainType() chains.ChainType
}

func ParseBootstrapPeers(peers []string) (bootstrapPeers []commontypes.BootstrapperLocator, err error) {
	for _, bs := range peers {
		var bsl commontypes.BootstrapperLocator
		err = bsl.UnmarshalText([]byte(bs))
		if err != nil {
			return nil, err
		}
		bootstrapPeers = append(bootstrapPeers, bsl)
	}
	return
}

// Will error unless at least one valid bootstrap peer is found
func GetValidatedBootstrapPeers(specPeers []string, configPeers []commontypes.BootstrapperLocator) ([]commontypes.BootstrapperLocator, error) {
	bootstrapPeers, err := ParseBootstrapPeers(specPeers)
	if err != nil {
		return nil, err
	}
	if len(bootstrapPeers) == 0 {
		if len(configPeers) == 0 {
			return nil, errors.New("no bootstrappers found")
		}
		return configPeers, nil
	}
	return bootstrapPeers, nil
}
