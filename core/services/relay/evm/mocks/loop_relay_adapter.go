// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	evm "github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink-relay/pkg/types"
)

// LoopRelayAdapter is an autogenerated mock type for the LoopRelayAdapter type
type LoopRelayAdapter struct {
	mock.Mock
}

// Chain provides a mock function with given fields:
func (_m *LoopRelayAdapter) Chain() evm.Chain {
	ret := _m.Called()

	var r0 evm.Chain
	if rf, ok := ret.Get(0).(func() evm.Chain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.Chain)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *LoopRelayAdapter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Default provides a mock function with given fields:
func (_m *LoopRelayAdapter) Default() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetChainStatus provides a mock function with given fields: ctx
func (_m *LoopRelayAdapter) GetChainStatus(ctx context.Context) (types.ChainStatus, error) {
	ret := _m.Called(ctx)

	var r0 types.ChainStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.ChainStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.ChainStatus); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.ChainStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthReport provides a mock function with given fields:
func (_m *LoopRelayAdapter) HealthReport() map[string]error {
	ret := _m.Called()

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

// ListNodeStatuses provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *LoopRelayAdapter) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) ([]types.NodeStatus, string, int, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	var r0 []types.NodeStatus
	var r1 string
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) []types.NodeStatus); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NodeStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, string) string); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int32, string) int); ok {
		r2 = rf(ctx, pageSize, pageToken)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(context.Context, int32, string) error); ok {
		r3 = rf(ctx, pageSize, pageToken)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Name provides a mock function with given fields:
func (_m *LoopRelayAdapter) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewConfigProvider provides a mock function with given fields: _a0, _a1
func (_m *LoopRelayAdapter) NewConfigProvider(_a0 context.Context, _a1 types.RelayArgs) (types.ConfigProvider, error) {
	ret := _m.Called(_a0, _a1)

	var r0 types.ConfigProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs) (types.ConfigProvider, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs) types.ConfigProvider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ConfigProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RelayArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPluginProvider provides a mock function with given fields: _a0, _a1, _a2
func (_m *LoopRelayAdapter) NewPluginProvider(_a0 context.Context, _a1 types.RelayArgs, _a2 types.PluginArgs) (types.PluginProvider, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.PluginProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) (types.PluginProvider, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) types.PluginProvider); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PluginProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RelayArgs, types.PluginArgs) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ready provides a mock function with given fields:
func (_m *LoopRelayAdapter) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *LoopRelayAdapter) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transact provides a mock function with given fields: ctx, from, to, amount, balanceCheck
func (_m *LoopRelayAdapter) Transact(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool) error {
	ret := _m.Called(ctx, from, to, amount, balanceCheck)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int, bool) error); ok {
		r0 = rf(ctx, from, to, amount, balanceCheck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLoopRelayAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoopRelayAdapter creates a new instance of LoopRelayAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoopRelayAdapter(t mockConstructorTestingTNewLoopRelayAdapter) *LoopRelayAdapter {
	mock := &LoopRelayAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
