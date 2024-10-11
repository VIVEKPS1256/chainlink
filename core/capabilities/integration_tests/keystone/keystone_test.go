package keystone

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/datastreams"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/integration_tests/framework"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/feeds_consumer"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	reporttypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/types"
)

func Test_AllAtOnceTransmissionSchedule(t *testing.T) {
	testTransmissionSchedule(t, "2s", "allAtOnce")
}

func Test_OneAtATimeTransmissionSchedule(t *testing.T) {
	testTransmissionSchedule(t, "2s", "oneAtATime")
}

func testTransmissionSchedule(t *testing.T, deltaStage string, schedule string) {
	ctx := testutils.Context(t)
	workflowDonConfiguration, err := framework.NewDonConfiguration(framework.NewDonConfigurationParams{Name: "Workflow", NumNodes: 7, F: 2, AcceptsWorkflows: true})
	require.NoError(t, err)
	triggerDonConfiguration, err := framework.NewDonConfiguration(framework.NewDonConfigurationParams{Name: "Trigger", NumNodes: 7, F: 2})
	require.NoError(t, err)
	targetDonConfiguration, err := framework.NewDonConfiguration(framework.NewDonConfigurationParams{Name: "Target", NumNodes: 4, F: 1})
	require.NoError(t, err)

	feedCount := 3
	var feedIDs []string
	for i := 0; i < feedCount; i++ {
		feedIDs = append(feedIDs, newFeedID(t))
	}

	keystoneWorkflowJobFactory := func(t *testing.T,
		workflowName string,
		workflowOwner string,
		consumerContractAddr common.Address) job.Job {

		return createKeystoneWorkflowJob(t, workflowName, workflowOwner, feedIDs, consumerContractAddr, deltaStage, schedule)

	}

	reportsSink := framework.NewReportsSink()
	consumer := setupKeystoneDons(ctx, t, workflowDonConfiguration, triggerDonConfiguration, targetDonConfiguration,
		keystoneWorkflowJobFactory, reportsSink)

	reports := []*datastreams.FeedReport{
		createFeedReport(t, big.NewInt(1), 5, feedIDs[0], triggerDonConfiguration.KeyBundles),
		createFeedReport(t, big.NewInt(3), 7, feedIDs[1], triggerDonConfiguration.KeyBundles),
		createFeedReport(t, big.NewInt(2), 6, feedIDs[2], triggerDonConfiguration.KeyBundles),
	}

	servicetest.Run(t, reportsSink)
	reportsSink.SendReports(reports)

	waitForConsumerReports(ctx, t, consumer, reports)
}

func newFeedID(t *testing.T) string {
	buf := [32]byte{}
	_, err := rand.Read(buf[:])
	require.NoError(t, err)
	return "0x" + hex.EncodeToString(buf[:])
}

func waitForConsumerReports(ctx context.Context, t *testing.T, consumer *feeds_consumer.KeystoneFeedsConsumer, triggerFeedReports []*datastreams.FeedReport) {
	feedsReceived := make(chan *feeds_consumer.KeystoneFeedsConsumerFeedReceived, 1000)
	feedsSub, err := consumer.WatchFeedReceived(&bind.WatchOpts{}, feedsReceived, nil)
	require.NoError(t, err)

	feedToReport := map[string]*datastreams.FeedReport{}
	for _, report := range triggerFeedReports {
		feedToReport[report.FeedID] = report
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	feedCount := 0
	for {
		select {
		case <-ctxWithTimeout.Done():
			t.Fatalf("timed out waiting for feed reports, expected %d, received %d", len(triggerFeedReports), feedCount)
		case err := <-feedsSub.Err():
			require.NoError(t, err)
		case feed := <-feedsReceived:
			feedID := "0x" + hex.EncodeToString(feed.FeedId[:])
			report := feedToReport[feedID]
			decodedReport, err := reporttypes.Decode(report.FullReport)
			require.NoError(t, err)
			assert.Equal(t, decodedReport.BenchmarkPrice, feed.Price)
			assert.Equal(t, decodedReport.ObservationsTimestamp, feed.Timestamp)

			feedCount++
			if feedCount == len(triggerFeedReports) {
				return
			}
		}
	}
}
