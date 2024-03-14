// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	synchronization "github.com/smartcontractkit/chainlink/v2/core/services/synchronization"
	mock "github.com/stretchr/testify/mock"
)

// TelemetryService is an autogenerated mock type for the TelemetryService type
type TelemetryService struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TelemetryService) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *TelemetryService) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *TelemetryService) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *TelemetryService) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: ctx, telemetry, contractID, telemType
func (_m *TelemetryService) Send(ctx context.Context, telemetry []byte, contractID string, telemType synchronization.TelemetryType) {
	_m.Called(ctx, telemetry, contractID, telemType)
}

// Start provides a mock function with given fields: _a0
func (_m *TelemetryService) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTelemetryService creates a new instance of TelemetryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTelemetryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TelemetryService {
	mock := &TelemetryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
