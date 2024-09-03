package test_env

import (
	"context"
	"github.com/rs/zerolog"
	d "github.com/smartcontractkit/chainlink/integration-tests/docker"
	tc "github.com/testcontainers/testcontainers-go"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func SaveCodeCoverageData(l zerolog.Logger, t *testing.T, nodes []*ClNode, showHTMLCoverageReport bool) error {
	testName := strings.ReplaceAll(t.Name(), "/", "_")
	isCI := os.Getenv("CI") != ""

	l.Info().
		Bool("showCoverageReportFlag", showHTMLCoverageReport).
		Bool("isCI", isCI).
		Bool("show", showHTMLCoverageReport || isCI).
		Msg("Checking if coverage report should be shown")

	var covHelper *d.NodeCoverageHelper

	if showHTMLCoverageReport || isCI {
		// Stop all nodes in the chainlink cluster.
		// This is needed to get go coverage profile from the node containers https://go.dev/doc/build-cover#FAQ
		// TODO: fix this as it results in: ERR LOG AFTER TEST ENDED ... INF 🐳 Stopping container
		cluster := ClCluster{Nodes: nodes}
		err := cluster.Stop()
		if err != nil {
			return err
		}

		clDir, err := getChainlinkDir()
		if err != nil {
			return err
		}

		var coverageRootDir string
		if os.Getenv("GO_COVERAGE_DEST_DIR") != "" {
			coverageRootDir = filepath.Join(os.Getenv("GO_COVERAGE_DEST_DIR"), testName)
		} else {
			coverageRootDir = filepath.Join(clDir, ".covdata", testName)
		}

		var containers []tc.Container
		for _, node := range cluster.Nodes {
			containers = append(containers, node.Container)
		}

		covHelper, err = d.NewNodeCoverageHelper(context.Background(), containers, clDir, coverageRootDir)
		if err != nil {
			return err
		}
	}

	// Show html coverage report when flag is set (local runs)
	if showHTMLCoverageReport {
		path, err := covHelper.SaveMergedHTMLReport()
		if err != nil {
			l.Error().Err(err).Msg("Error saving merged html report")
			return err
		}
		l.Info().Str("testName", testName).Str("filePath", path).Msg("Chainlink node coverage html report saved")
	}

	// Save percentage coverage report when running in CI
	if isCI {
		// Save coverage percentage to a file to show in the CI
		path, err := covHelper.SaveMergedCoveragePercentage()
		if err != nil {
			l.Error().Err(err).Str("testName", testName).Msg("Failed to save coverage percentage for test")
			return err
		} else {
			l.Info().Str("testName", testName).Str("filePath", path).Msg("Chainlink node coverage percentage report saved")
		}
	}

	return nil
}
