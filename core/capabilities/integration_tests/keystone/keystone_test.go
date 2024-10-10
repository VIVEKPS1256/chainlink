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
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/integration_tests/framework"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/feeds_consumer"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	reporttypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/types"
)

func Test_AllAtOnceTransmissionSchedule(t *testing.T) {
	ctx := testutils.Context(t)

	// The don IDs set in the below calls are inferred from the order in which the dons are added to the capabilities registry
	// in the setupCapabilitiesRegistryContract function, should this order change the don IDs will need updating.
	workflowDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Workflow", ID: 2, NumNodes: 7, F: 2, AcceptsWorkflows: true})
	triggerDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Trigger", ID: 1, NumNodes: 7, F: 2})
	targetDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Target", ID: 3, NumNodes: 4, F: 1})

	feedCount := 3
	var feedIDs []string
	for i := 0; i < feedCount; i++ {
		feedIDs = append(feedIDs, newFeedID(t))
	}

	createKeystoneWorkflowJob := func(t *testing.T,
		workflowName string,
		workflowOwner string,
		consumerContractAddr common.Address) job.Job {

		return createKeystoneWorkflowJob(t, workflowName, workflowOwner, feedIDs, consumerContractAddr, "2s", "allAtOnce")

	}

	consumer, triggerSink := framework.SetupDons(ctx, t, workflowDonInfo, triggerDonInfo, targetDonInfo, createKeystoneWorkflowJob)

	reports := []*datastreams.FeedReport{
		createFeedReport(t, big.NewInt(1), 5, feedIDs[0], triggerDonInfo.KeyBundles),
		createFeedReport(t, big.NewInt(3), 7, feedIDs[1], triggerDonInfo.KeyBundles),
		createFeedReport(t, big.NewInt(2), 6, feedIDs[2], triggerDonInfo.KeyBundles),
	}

	triggerSink.SendReports(reports)

	waitForConsumerReports(ctx, t, consumer, reports)
}

/*
func Test_OneAtATimeTransmissionSchedule(t *testing.T) {
	ctx := testutils.Context(t)

	// The don IDs set in the below calls are inferred from the order in which the dons are added to the capabilities registry
	// in the setupCapabilitiesRegistryContract function, should this order change the don IDs will need updating.
	workflowDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Workflow", ID: 1, NumNodes: 7, F: 2})
	triggerDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Trigger", ID: 2, NumNodes: 7, F: 2})
	targetDonInfo := framework.CreateDonInfo(t, framework.Don{Name: "Target", ID: 3, NumNodes: 4, F: 1})

	feedCount := 3
	var feedIDs []string
	for i := 0; i < feedCount; i++ {
		feedIDs = append(feedIDs, newFeedID(t))
	}

	createKeystoneWorkflowJob := func(t *testing.T,
		workflowName string,
		workflowOwner string,
		consumerContractAddr common.Address) job.Job {

		return createKeystoneWorkflowJob(t, workflowName, workflowOwner, feedIDs, consumerContractAddr, "2s", "oneAtATime")

	}

	consumer, triggerSink := framework.SetupDons(ctx, t, workflowDonInfo, triggerDonInfo, targetDonInfo, createKeystoneWorkflowJob)

	reports := []*datastreams.FeedReport{
		createFeedReport(t, big.NewInt(1), 5, feedIDs[0], triggerDonInfo.KeyBundles),
		createFeedReport(t, big.NewInt(3), 7, feedIDs[1], triggerDonInfo.KeyBundles),
		createFeedReport(t, big.NewInt(2), 6, feedIDs[2], triggerDonInfo.KeyBundles),
	}

	triggerSink.SendReports(reports)

	waitForConsumerReports(ctx, t, consumer, reports)
}*/

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
