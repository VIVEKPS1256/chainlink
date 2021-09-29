// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/satori/go.uuid"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// TxStrategy is an autogenerated mock type for the TxStrategy type
type TxStrategy struct {
	mock.Mock
}

// PruneQueue provides a mock function with given fields: tx
func (_m *TxStrategy) PruneQueue(tx *gorm.DB) (int64, error) {
	ret := _m.Called(tx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*gorm.DB) int64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Simulate provides a mock function with given fields:
func (_m *TxStrategy) Simulate() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Subject provides a mock function with given fields:
func (_m *TxStrategy) Subject() uuid.NullUUID {
	ret := _m.Called()

	var r0 uuid.NullUUID
	if rf, ok := ret.Get(0).(func() uuid.NullUUID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uuid.NullUUID)
	}

	return r0
}
