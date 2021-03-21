// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	context "context"

	logger "github.com/smartcontractkit/chainlink/core/logger"
	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"
)

// Runner is an autogenerated mock type for the Runner type
type Runner struct {
	mock.Mock
}

// AwaitRun provides a mock function with given fields: ctx, runID
func (_m *Runner) AwaitRun(ctx context.Context, runID int64) error {
	ret := _m.Called(ctx, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Runner) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRun provides a mock function with given fields: ctx, jobID, meta
func (_m *Runner) CreateRun(ctx context.Context, jobID int32, meta map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, jobID, meta)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int32, map[string]interface{}) int64); ok {
		r0 = rf(ctx, jobID, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, map[string]interface{}) error); ok {
		r1 = rf(ctx, jobID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteAndInsertNewRun provides a mock function with given fields: ctx, spec, l
func (_m *Runner) ExecuteAndInsertNewRun(ctx context.Context, spec pipeline.Spec, l logger.Logger) (int64, pipeline.FinalResult, error) {
	ret := _m.Called(ctx, spec, l)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Spec, logger.Logger) int64); ok {
		r0 = rf(ctx, spec, l)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 pipeline.FinalResult
	if rf, ok := ret.Get(1).(func(context.Context, pipeline.Spec, logger.Logger) pipeline.FinalResult); ok {
		r1 = rf(ctx, spec, l)
	} else {
		r1 = ret.Get(1).(pipeline.FinalResult)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, pipeline.Spec, logger.Logger) error); ok {
		r2 = rf(ctx, spec, l)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ExecuteRun provides a mock function with given fields: ctx, spec, l
func (_m *Runner) ExecuteRun(ctx context.Context, spec pipeline.Spec, l logger.Logger) (pipeline.TaskRunResults, error) {
	ret := _m.Called(ctx, spec, l)

	var r0 pipeline.TaskRunResults
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Spec, logger.Logger) pipeline.TaskRunResults); ok {
		r0 = rf(ctx, spec, l)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.TaskRunResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pipeline.Spec, logger.Logger) error); ok {
		r1 = rf(ctx, spec, l)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertFinishedRunWithResults provides a mock function with given fields: ctx, run, trrs
func (_m *Runner) InsertFinishedRunWithResults(ctx context.Context, run pipeline.Run, trrs pipeline.TaskRunResults) (int64, error) {
	ret := _m.Called(ctx, run, trrs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Run, pipeline.TaskRunResults) int64); ok {
		r0 = rf(ctx, run, trrs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pipeline.Run, pipeline.TaskRunResults) error); ok {
		r1 = rf(ctx, run, trrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResultsForRun provides a mock function with given fields: ctx, runID
func (_m *Runner) ResultsForRun(ctx context.Context, runID int64) ([]pipeline.Result, error) {
	ret := _m.Called(ctx, runID)

	var r0 []pipeline.Result
	if rf, ok := ret.Get(0).(func(context.Context, int64) []pipeline.Result); ok {
		r0 = rf(ctx, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, runID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Runner) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
