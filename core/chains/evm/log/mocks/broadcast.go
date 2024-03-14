// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Broadcast is an autogenerated mock type for the Broadcast type
type Broadcast struct {
	mock.Mock
}

// DecodedLog provides a mock function with given fields:
func (_m *Broadcast) DecodedLog() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DecodedLog")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// EVMChainID provides a mock function with given fields:
func (_m *Broadcast) EVMChainID() big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EVMChainID")
	}

	var r0 big.Int
	if rf, ok := ret.Get(0).(func() big.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(big.Int)
	}

	return r0
}

// JobID provides a mock function with given fields:
func (_m *Broadcast) JobID() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JobID")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// LatestBlockHash provides a mock function with given fields:
func (_m *Broadcast) LatestBlockHash() common.Hash {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LatestBlockHash")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// LatestBlockNumber provides a mock function with given fields:
func (_m *Broadcast) LatestBlockNumber() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LatestBlockNumber")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// RawLog provides a mock function with given fields:
func (_m *Broadcast) RawLog() types.Log {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RawLog")
	}

	var r0 types.Log
	if rf, ok := ret.Get(0).(func() types.Log); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Log)
	}

	return r0
}

// ReceiptsRoot provides a mock function with given fields:
func (_m *Broadcast) ReceiptsRoot() common.Hash {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReceiptsRoot")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// StateRoot provides a mock function with given fields:
func (_m *Broadcast) StateRoot() common.Hash {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StateRoot")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Broadcast) String() string {
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

// TransactionsRoot provides a mock function with given fields:
func (_m *Broadcast) TransactionsRoot() common.Hash {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TransactionsRoot")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// NewBroadcast creates a new instance of Broadcast. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBroadcast(t interface {
	mock.TestingT
	Cleanup(func())
}) *Broadcast {
	mock := &Broadcast{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
