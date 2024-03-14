// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	functions "github.com/smartcontractkit/chainlink/v2/core/services/functions"
	mock "github.com/stretchr/testify/mock"
)

// FunctionsListener is an autogenerated mock type for the FunctionsListener type
type FunctionsListener struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *FunctionsListener) Close() error {
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

// HandleOffchainRequest provides a mock function with given fields: ctx, request
func (_m *FunctionsListener) HandleOffchainRequest(ctx context.Context, request *functions.OffchainRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for HandleOffchainRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.OffchainRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *FunctionsListener) Start(_a0 context.Context) error {
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

// NewFunctionsListener creates a new instance of FunctionsListener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFunctionsListener(t interface {
	mock.TestingT
	Cleanup(func())
}) *FunctionsListener {
	mock := &FunctionsListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
