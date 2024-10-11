package keystone

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/datastreams"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	v3 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v3"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/integration_tests/framework"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/feeds_consumer"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"

	ocrTypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/reportcodec"
)

var (
	workflowName    = "abcdef0123"
	workflowOwnerID = "0100000000000000000000000000000000000001"
)

func setupDons(ctx context.Context, t *testing.T, workflowDonInfo framework.DonConfiguration, triggerDonInfo framework.DonConfiguration,
	targetDonInfo framework.DonConfiguration,
	getJob func(t *testing.T,
		workflowName string,
		workflowOwner string,
		consumerAddr common.Address) job.Job) (*feeds_consumer.KeystoneFeedsConsumer, *framework.ReportsSink) {
	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(zapcore.InfoLevel)

	donContext := framework.CreateDonContext(ctx, t)

	workflowDon := framework.NewDON(ctx, t, lggr, workflowDonInfo,
		[]commoncap.DON{triggerDonInfo.DON, targetDonInfo.DON},
		donContext)

	workflowDon.AddOCR3NonStandardCapability(ctx, t)
	workflowDon.Initialise()

	forwarderAddr, _ := SetupForwarderContract(t, workflowDon, donContext.EthBlockchain)
	consumerAddr, consumer := SetupConsumerContract(t, donContext.EthBlockchain, forwarderAddr, workflowOwnerID, workflowName)

	job := getJob(t, workflowName, workflowOwnerID, consumerAddr)
	workflowDon.AddJob(&job)

	writeTargetDon := framework.NewDON(ctx, t, lggr, targetDonInfo,
		[]commoncap.DON{}, donContext)
	err := writeTargetDon.AddEthereumWriteTargetNonStandardCapability(forwarderAddr)
	require.NoError(t, err)
	writeTargetDon.Initialise()

	sink := framework.NewReportsSink()

	triggerDon := framework.NewDON(ctx, t, lggr, triggerDonInfo,
		[]commoncap.DON{}, donContext)

	triggerDon.AddTriggerCapability(sink)
	triggerDon.Initialise()

	triggerDon.Start(ctx, t)
	writeTargetDon.Start(ctx, t)
	workflowDon.Start(ctx, t)

	servicetest.Run(t, sink)
	return consumer, sink
}

func createFeedReport(t *testing.T, price *big.Int, observationTimestamp int64,
	feedIDString string,
	keyBundles []ocr2key.KeyBundle) *datastreams.FeedReport {
	reportCtx := ocrTypes.ReportContext{}
	rawCtx := RawReportContext(reportCtx)

	bytes, err := hex.DecodeString(feedIDString[2:])
	require.NoError(t, err)
	var feedIDBytes [32]byte
	copy(feedIDBytes[:], bytes)

	report := &datastreams.FeedReport{
		FeedID:               feedIDString,
		FullReport:           newReport(t, feedIDBytes, price, observationTimestamp),
		BenchmarkPrice:       price.Bytes(),
		ObservationTimestamp: observationTimestamp,
		Signatures:           [][]byte{},
		ReportContext:        rawCtx,
	}

	for _, key := range keyBundles {
		sig, err := key.Sign(reportCtx, report.FullReport)
		require.NoError(t, err)
		report.Signatures = append(report.Signatures, sig)
	}

	return report
}

func RawReportContext(reportCtx ocrTypes.ReportContext) []byte {
	rc := evmutil.RawReportContext(reportCtx)
	flat := []byte{}
	for _, r := range rc {
		flat = append(flat, r[:]...)
	}
	return flat
}

func newReport(t *testing.T, feedID [32]byte, price *big.Int, timestamp int64) []byte {
	v3Codec := reportcodec.NewReportCodec(feedID, logger.TestLogger(t))
	raw, err := v3Codec.BuildReport(v3.ReportFields{
		BenchmarkPrice: price,
		Timestamp:      uint32(timestamp),
		Bid:            big.NewInt(0),
		Ask:            big.NewInt(0),
		LinkFee:        big.NewInt(0),
		NativeFee:      big.NewInt(0),
	})
	require.NoError(t, err)
	return raw
}
