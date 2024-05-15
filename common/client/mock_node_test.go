// Code generated by mockery v2.42.2. DO NOT EDIT.

package client

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// mockNode is an autogenerated mock type for the Node type
type mockNode[CHAIN_ID types.ID, HEAD Head, RPC_CLIENT RPCClient[CHAIN_ID, HEAD]] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) Close() error {
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

// ConfiguredChainID provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) ConfiguredChainID() CHAIN_ID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConfiguredChainID")
	}

	var r0 CHAIN_ID
	if rf, ok := ret.Get(0).(func() CHAIN_ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) Name() string {
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

// Order provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) Order() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Order")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// RPC provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) RPC() RPC_CLIENT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RPC")
	}

	var r0 RPC_CLIENT
	if rf, ok := ret.Get(0).(func() RPC_CLIENT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(RPC_CLIENT)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) Start(_a0 context.Context) error {
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

// State provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) State() nodeState {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for State")
	}

	var r0 nodeState
	if rf, ok := ret.Get(0).(func() nodeState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(nodeState)
	}

	return r0
}

// StateAndLatest provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) StateAndLatest() (nodeState, ChainInfo) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StateAndLatest")
	}

	var r0 nodeState
	var r1 ChainInfo
	if rf, ok := ret.Get(0).(func() (nodeState, ChainInfo)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() nodeState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(nodeState)
	}

	if rf, ok := ret.Get(1).(func() ChainInfo); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(ChainInfo)
	}

	return r0, r1
}

// String provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UnsubscribeAll provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, HEAD, RPC_CLIENT]) UnsubscribeAll() {
	_m.Called()
}

// newMockNode creates a new instance of mockNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockNode[CHAIN_ID types.ID, HEAD Head, RPC_CLIENT RPCClient[CHAIN_ID, HEAD]](t interface {
	mock.TestingT
	Cleanup(func())
}) *mockNode[CHAIN_ID, HEAD, RPC_CLIENT] {
	mock := &mockNode[CHAIN_ID, HEAD, RPC_CLIENT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
