// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	config "github.com/smartcontractkit/chainlink-common/pkg/config"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// DefaultHTTPLimit provides a mock function with given fields:
func (_m *Config) DefaultHTTPLimit() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultHTTPLimit")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// DefaultHTTPTimeout provides a mock function with given fields:
func (_m *Config) DefaultHTTPTimeout() config.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultHTTPTimeout")
	}

	var r0 config.Duration
	if rf, ok := ret.Get(0).(func() config.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Duration)
	}

	return r0
}

// MaxRunDuration provides a mock function with given fields:
func (_m *Config) MaxRunDuration() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MaxRunDuration")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// ReaperInterval provides a mock function with given fields:
func (_m *Config) ReaperInterval() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReaperInterval")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// ReaperThreshold provides a mock function with given fields:
func (_m *Config) ReaperThreshold() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReaperThreshold")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
