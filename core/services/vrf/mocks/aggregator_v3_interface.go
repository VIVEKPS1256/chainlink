// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	aggregator_v3_interface "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/aggregator_v3_interface"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// AggregatorV3Interface is an autogenerated mock type for the AggregatorV3InterfaceInterface type
type AggregatorV3Interface struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *AggregatorV3Interface) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// Decimals provides a mock function with given fields: opts
func (_m *AggregatorV3Interface) Decimals(opts *bind.CallOpts) (uint8, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Decimals")
	}

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (uint8, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint8); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Description provides a mock function with given fields: opts
func (_m *AggregatorV3Interface) Description(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (string, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoundData provides a mock function with given fields: opts, _roundId
func (_m *AggregatorV3Interface) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (aggregator_v3_interface.GetRoundData, error) {
	ret := _m.Called(opts, _roundId)

	if len(ret) == 0 {
		panic("no return value specified for GetRoundData")
	}

	var r0 aggregator_v3_interface.GetRoundData
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) (aggregator_v3_interface.GetRoundData, error)); ok {
		return rf(opts, _roundId)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) aggregator_v3_interface.GetRoundData); ok {
		r0 = rf(opts, _roundId)
	} else {
		r0 = ret.Get(0).(aggregator_v3_interface.GetRoundData)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, _roundId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestRoundData provides a mock function with given fields: opts
func (_m *AggregatorV3Interface) LatestRoundData(opts *bind.CallOpts) (aggregator_v3_interface.LatestRoundData, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for LatestRoundData")
	}

	var r0 aggregator_v3_interface.LatestRoundData
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (aggregator_v3_interface.LatestRoundData, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) aggregator_v3_interface.LatestRoundData); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(aggregator_v3_interface.LatestRoundData)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields: opts
func (_m *AggregatorV3Interface) Version(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (*big.Int, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAggregatorV3Interface creates a new instance of AggregatorV3Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAggregatorV3Interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AggregatorV3Interface {
	mock := &AggregatorV3Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
