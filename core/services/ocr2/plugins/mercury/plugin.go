package mercury

import (
	"encoding/json"

	"github.com/pkg/errors"

	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"

	relaymercury "github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/mercury"
	relaytypes "github.com/smartcontractkit/chainlink-relay/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/mercury/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury"
)

type Config interface {
	MaxSuccessfulRuns() uint64
}

func NewServices(
	jb job.Job,
	ocr2Provider relaytypes.MercuryProvider,
	pipelineRunner pipeline.Runner,
	runResults chan pipeline.Run,
	lggr logger.Logger,
	argsNoPlugin libocr2.MercuryOracleArgs,
	chEnhancedTelem chan ocrcommon.EnhancedTelemetryMercuryData,
	chainHeadTracker mercury.ChainHeadTracker,
	runSaver *ocrcommon.RunResultSaver,
) ([]job.ServiceCtx, error) {
	if jb.PipelineSpec == nil {
		return nil, errors.New("expected job to have a non-nil PipelineSpec")
	}
	var pluginConfig config.PluginConfig
	err := json.Unmarshal(jb.OCR2OracleSpec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = config.ValidatePluginConfig(pluginConfig)
	if err != nil {
		return nil, err
	}
	lggr = lggr.Named("MercuryPlugin").With("jobID", jb.ID, "jobName", jb.Name.ValueOrZero())
	ds := mercury.NewDataSource(
		pipelineRunner,
		jb,
		*jb.PipelineSpec,
		lggr,
		runResults,
		chEnhancedTelem,
		chainHeadTracker,
		ocr2Provider.ContractTransmitter(),
		pluginConfig.InitialBlockNumber.Ptr(),
	)
	argsNoPlugin.MercuryPluginFactory = relaymercury.NewFactory(
		ds,
		lggr,
		ocr2Provider.OnchainConfigCodec(),
		ocr2Provider.ReportCodec(),
	)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return []job.ServiceCtx{ocr2Provider, runSaver,
		job.NewServiceAdapter(oracle)}, nil
}
