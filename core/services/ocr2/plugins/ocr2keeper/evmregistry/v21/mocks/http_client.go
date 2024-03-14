// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Do provides a mock function with given fields: req
func (_m *HttpClient) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*http.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHttpClient creates a new instance of HttpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHttpClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *HttpClient {
	mock := &HttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
