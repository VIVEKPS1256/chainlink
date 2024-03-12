// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// KVStore is an autogenerated mock type for the KVStore type
type KVStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: key, dest
func (_m *KVStore) Get(key string, dest interface{}) error {
	ret := _m.Called(key, dest)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: key, val
func (_m *KVStore) Store(key string, val interface{}) error {
	ret := _m.Called(key, val)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKVStore creates a new instance of KVStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKVStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *KVStore {
	mock := &KVStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
