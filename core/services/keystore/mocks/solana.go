// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	solkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/solkey"
)

// Solana is an autogenerated mock type for the Solana type
type Solana struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *Solana) Add(key solkey.Key) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(solkey.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *Solana) Create() (solkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() (solkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() solkey.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Solana) Delete(id string) (solkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (solkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) solkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *Solana) EnsureKey() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *Solana) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Solana) Get(id string) (solkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (solkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) solkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Solana) GetAll() ([]solkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]solkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []solkey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]solkey.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *Solana) Import(keyJSON []byte, password string) (solkey.Key, error) {
	ret := _m.Called(keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (solkey.Key, error)); ok {
		return rf(keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) solkey.Key); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sign provides a mock function with given fields: ctx, id, msg
func (_m *Solana) Sign(ctx context.Context, id string, msg []byte) ([]byte, error) {
	ret := _m.Called(ctx, id, msg)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) ([]byte, error)); ok {
		return rf(ctx, id, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) []byte); ok {
		r0 = rf(ctx, id, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, id, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSolana creates a new instance of Solana. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSolana(t interface {
	mock.TestingT
	Cleanup(func())
}) *Solana {
	mock := &Solana{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
