// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HTTPRequestHandler is an autogenerated mock type for the HTTPRequestHandler type
type HTTPRequestHandler struct {
	mock.Mock
}

// ProcessRequest provides a mock function with given fields: ctx, rawRequest
func (_m *HTTPRequestHandler) ProcessRequest(ctx context.Context, rawRequest []byte) ([]byte, int) {
	ret := _m.Called(ctx, rawRequest)

	if len(ret) == 0 {
		panic("no return value specified for ProcessRequest")
	}

	var r0 []byte
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, []byte) ([]byte, int)); ok {
		return rf(ctx, rawRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) []byte); ok {
		r0 = rf(ctx, rawRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) int); ok {
		r1 = rf(ctx, rawRequest)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// NewHTTPRequestHandler creates a new instance of HTTPRequestHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHTTPRequestHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *HTTPRequestHandler {
	mock := &HTTPRequestHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
