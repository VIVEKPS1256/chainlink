// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	job "github.com/smartcontractkit/chainlink/core/services/job"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// Delegate is an autogenerated mock type for the Delegate type
type Delegate struct {
	mock.Mock
}

// FromDBRow provides a mock function with given fields: spec
func (_m *Delegate) FromDBRow(spec models.JobSpecV2) job.Spec {
	ret := _m.Called(spec)

	var r0 job.Spec
	if rf, ok := ret.Get(0).(func(models.JobSpecV2) job.Spec); ok {
		r0 = rf(spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.Spec)
		}
	}

	return r0
}

// JobType provides a mock function with given fields:
func (_m *Delegate) JobType() job.Type {
	ret := _m.Called()

	var r0 job.Type
	if rf, ok := ret.Get(0).(func() job.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(job.Type)
	}

	return r0
}

// ServicesForSpec provides a mock function with given fields: spec
func (_m *Delegate) ServicesForSpec(spec job.Spec) ([]job.Service, error) {
	ret := _m.Called(spec)

	var r0 []job.Service
	if rf, ok := ret.Get(0).(func(job.Spec) []job.Service); ok {
		r0 = rf(spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(job.Spec) error); ok {
		r1 = rf(spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToDBRow provides a mock function with given fields: spec
func (_m *Delegate) ToDBRow(spec job.Spec) models.JobSpecV2 {
	ret := _m.Called(spec)

	var r0 models.JobSpecV2
	if rf, ok := ret.Get(0).(func(job.Spec) models.JobSpecV2); ok {
		r0 = rf(spec)
	} else {
		r0 = ret.Get(0).(models.JobSpecV2)
	}

	return r0
}
