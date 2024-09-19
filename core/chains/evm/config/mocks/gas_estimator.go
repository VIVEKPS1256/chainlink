// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	assets "github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"

	config "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"

	mock "github.com/stretchr/testify/mock"
)

// GasEstimator is an autogenerated mock type for the GasEstimator type
type GasEstimator struct {
	mock.Mock
}

type GasEstimator_Expecter struct {
	mock *mock.Mock
}

func (_m *GasEstimator) EXPECT() *GasEstimator_Expecter {
	return &GasEstimator_Expecter{mock: &_m.Mock}
}

// BlockHistory provides a mock function with given fields:
func (_m *GasEstimator) BlockHistory() config.BlockHistory {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BlockHistory")
	}

	var r0 config.BlockHistory
	if rf, ok := ret.Get(0).(func() config.BlockHistory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.BlockHistory)
		}
	}

	return r0
}

// GasEstimator_BlockHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockHistory'
type GasEstimator_BlockHistory_Call struct {
	*mock.Call
}

// BlockHistory is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) BlockHistory() *GasEstimator_BlockHistory_Call {
	return &GasEstimator_BlockHistory_Call{Call: _e.mock.On("BlockHistory")}
}

func (_c *GasEstimator_BlockHistory_Call) Run(run func()) *GasEstimator_BlockHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_BlockHistory_Call) Return(_a0 config.BlockHistory) *GasEstimator_BlockHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_BlockHistory_Call) RunAndReturn(run func() config.BlockHistory) *GasEstimator_BlockHistory_Call {
	_c.Call.Return(run)
	return _c
}

// BumpMin provides a mock function with given fields:
func (_m *GasEstimator) BumpMin() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BumpMin")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_BumpMin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpMin'
type GasEstimator_BumpMin_Call struct {
	*mock.Call
}

// BumpMin is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) BumpMin() *GasEstimator_BumpMin_Call {
	return &GasEstimator_BumpMin_Call{Call: _e.mock.On("BumpMin")}
}

func (_c *GasEstimator_BumpMin_Call) Run(run func()) *GasEstimator_BumpMin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_BumpMin_Call) Return(_a0 *assets.Wei) *GasEstimator_BumpMin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_BumpMin_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_BumpMin_Call {
	_c.Call.Return(run)
	return _c
}

// BumpPercent provides a mock function with given fields:
func (_m *GasEstimator) BumpPercent() uint16 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BumpPercent")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// GasEstimator_BumpPercent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpPercent'
type GasEstimator_BumpPercent_Call struct {
	*mock.Call
}

// BumpPercent is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) BumpPercent() *GasEstimator_BumpPercent_Call {
	return &GasEstimator_BumpPercent_Call{Call: _e.mock.On("BumpPercent")}
}

func (_c *GasEstimator_BumpPercent_Call) Run(run func()) *GasEstimator_BumpPercent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_BumpPercent_Call) Return(_a0 uint16) *GasEstimator_BumpPercent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_BumpPercent_Call) RunAndReturn(run func() uint16) *GasEstimator_BumpPercent_Call {
	_c.Call.Return(run)
	return _c
}

// BumpThreshold provides a mock function with given fields:
func (_m *GasEstimator) BumpThreshold() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BumpThreshold")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GasEstimator_BumpThreshold_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpThreshold'
type GasEstimator_BumpThreshold_Call struct {
	*mock.Call
}

// BumpThreshold is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) BumpThreshold() *GasEstimator_BumpThreshold_Call {
	return &GasEstimator_BumpThreshold_Call{Call: _e.mock.On("BumpThreshold")}
}

func (_c *GasEstimator_BumpThreshold_Call) Run(run func()) *GasEstimator_BumpThreshold_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_BumpThreshold_Call) Return(_a0 uint64) *GasEstimator_BumpThreshold_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_BumpThreshold_Call) RunAndReturn(run func() uint64) *GasEstimator_BumpThreshold_Call {
	_c.Call.Return(run)
	return _c
}

// BumpTxDepth provides a mock function with given fields:
func (_m *GasEstimator) BumpTxDepth() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BumpTxDepth")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// GasEstimator_BumpTxDepth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpTxDepth'
type GasEstimator_BumpTxDepth_Call struct {
	*mock.Call
}

// BumpTxDepth is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) BumpTxDepth() *GasEstimator_BumpTxDepth_Call {
	return &GasEstimator_BumpTxDepth_Call{Call: _e.mock.On("BumpTxDepth")}
}

func (_c *GasEstimator_BumpTxDepth_Call) Run(run func()) *GasEstimator_BumpTxDepth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_BumpTxDepth_Call) Return(_a0 uint32) *GasEstimator_BumpTxDepth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_BumpTxDepth_Call) RunAndReturn(run func() uint32) *GasEstimator_BumpTxDepth_Call {
	_c.Call.Return(run)
	return _c
}

// EIP1559DynamicFees provides a mock function with given fields:
func (_m *GasEstimator) EIP1559DynamicFees() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EIP1559DynamicFees")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GasEstimator_EIP1559DynamicFees_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EIP1559DynamicFees'
type GasEstimator_EIP1559DynamicFees_Call struct {
	*mock.Call
}

// EIP1559DynamicFees is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) EIP1559DynamicFees() *GasEstimator_EIP1559DynamicFees_Call {
	return &GasEstimator_EIP1559DynamicFees_Call{Call: _e.mock.On("EIP1559DynamicFees")}
}

func (_c *GasEstimator_EIP1559DynamicFees_Call) Run(run func()) *GasEstimator_EIP1559DynamicFees_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_EIP1559DynamicFees_Call) Return(_a0 bool) *GasEstimator_EIP1559DynamicFees_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_EIP1559DynamicFees_Call) RunAndReturn(run func() bool) *GasEstimator_EIP1559DynamicFees_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateLimit provides a mock function with given fields:
func (_m *GasEstimator) EstimateLimit() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EstimateLimit")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GasEstimator_EstimateLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateLimit'
type GasEstimator_EstimateLimit_Call struct {
	*mock.Call
}

// EstimateLimit is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) EstimateLimit() *GasEstimator_EstimateLimit_Call {
	return &GasEstimator_EstimateLimit_Call{Call: _e.mock.On("EstimateLimit")}
}

func (_c *GasEstimator_EstimateLimit_Call) Run(run func()) *GasEstimator_EstimateLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_EstimateLimit_Call) Return(_a0 bool) *GasEstimator_EstimateLimit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_EstimateLimit_Call) RunAndReturn(run func() bool) *GasEstimator_EstimateLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FeeCapDefault provides a mock function with given fields:
func (_m *GasEstimator) FeeCapDefault() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FeeCapDefault")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_FeeCapDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeeCapDefault'
type GasEstimator_FeeCapDefault_Call struct {
	*mock.Call
}

// FeeCapDefault is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) FeeCapDefault() *GasEstimator_FeeCapDefault_Call {
	return &GasEstimator_FeeCapDefault_Call{Call: _e.mock.On("FeeCapDefault")}
}

func (_c *GasEstimator_FeeCapDefault_Call) Run(run func()) *GasEstimator_FeeCapDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_FeeCapDefault_Call) Return(_a0 *assets.Wei) *GasEstimator_FeeCapDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_FeeCapDefault_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_FeeCapDefault_Call {
	_c.Call.Return(run)
	return _c
}

// FeeHistory provides a mock function with given fields:
func (_m *GasEstimator) FeeHistory() config.FeeHistory {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FeeHistory")
	}

	var r0 config.FeeHistory
	if rf, ok := ret.Get(0).(func() config.FeeHistory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.FeeHistory)
		}
	}

	return r0
}

// GasEstimator_FeeHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeeHistory'
type GasEstimator_FeeHistory_Call struct {
	*mock.Call
}

// FeeHistory is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) FeeHistory() *GasEstimator_FeeHistory_Call {
	return &GasEstimator_FeeHistory_Call{Call: _e.mock.On("FeeHistory")}
}

func (_c *GasEstimator_FeeHistory_Call) Run(run func()) *GasEstimator_FeeHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_FeeHistory_Call) Return(_a0 config.FeeHistory) *GasEstimator_FeeHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_FeeHistory_Call) RunAndReturn(run func() config.FeeHistory) *GasEstimator_FeeHistory_Call {
	_c.Call.Return(run)
	return _c
}

// LimitDefault provides a mock function with given fields:
func (_m *GasEstimator) LimitDefault() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LimitDefault")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GasEstimator_LimitDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitDefault'
type GasEstimator_LimitDefault_Call struct {
	*mock.Call
}

// LimitDefault is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) LimitDefault() *GasEstimator_LimitDefault_Call {
	return &GasEstimator_LimitDefault_Call{Call: _e.mock.On("LimitDefault")}
}

func (_c *GasEstimator_LimitDefault_Call) Run(run func()) *GasEstimator_LimitDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_LimitDefault_Call) Return(_a0 uint64) *GasEstimator_LimitDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_LimitDefault_Call) RunAndReturn(run func() uint64) *GasEstimator_LimitDefault_Call {
	_c.Call.Return(run)
	return _c
}

// LimitJobType provides a mock function with given fields:
func (_m *GasEstimator) LimitJobType() config.LimitJobType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LimitJobType")
	}

	var r0 config.LimitJobType
	if rf, ok := ret.Get(0).(func() config.LimitJobType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.LimitJobType)
		}
	}

	return r0
}

// GasEstimator_LimitJobType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitJobType'
type GasEstimator_LimitJobType_Call struct {
	*mock.Call
}

// LimitJobType is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) LimitJobType() *GasEstimator_LimitJobType_Call {
	return &GasEstimator_LimitJobType_Call{Call: _e.mock.On("LimitJobType")}
}

func (_c *GasEstimator_LimitJobType_Call) Run(run func()) *GasEstimator_LimitJobType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_LimitJobType_Call) Return(_a0 config.LimitJobType) *GasEstimator_LimitJobType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_LimitJobType_Call) RunAndReturn(run func() config.LimitJobType) *GasEstimator_LimitJobType_Call {
	_c.Call.Return(run)
	return _c
}

// LimitMax provides a mock function with given fields:
func (_m *GasEstimator) LimitMax() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LimitMax")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GasEstimator_LimitMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitMax'
type GasEstimator_LimitMax_Call struct {
	*mock.Call
}

// LimitMax is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) LimitMax() *GasEstimator_LimitMax_Call {
	return &GasEstimator_LimitMax_Call{Call: _e.mock.On("LimitMax")}
}

func (_c *GasEstimator_LimitMax_Call) Run(run func()) *GasEstimator_LimitMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_LimitMax_Call) Return(_a0 uint64) *GasEstimator_LimitMax_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_LimitMax_Call) RunAndReturn(run func() uint64) *GasEstimator_LimitMax_Call {
	_c.Call.Return(run)
	return _c
}

// LimitMultiplier provides a mock function with given fields:
func (_m *GasEstimator) LimitMultiplier() float32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LimitMultiplier")
	}

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

// GasEstimator_LimitMultiplier_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitMultiplier'
type GasEstimator_LimitMultiplier_Call struct {
	*mock.Call
}

// LimitMultiplier is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) LimitMultiplier() *GasEstimator_LimitMultiplier_Call {
	return &GasEstimator_LimitMultiplier_Call{Call: _e.mock.On("LimitMultiplier")}
}

func (_c *GasEstimator_LimitMultiplier_Call) Run(run func()) *GasEstimator_LimitMultiplier_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_LimitMultiplier_Call) Return(_a0 float32) *GasEstimator_LimitMultiplier_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_LimitMultiplier_Call) RunAndReturn(run func() float32) *GasEstimator_LimitMultiplier_Call {
	_c.Call.Return(run)
	return _c
}

// LimitTransfer provides a mock function with given fields:
func (_m *GasEstimator) LimitTransfer() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LimitTransfer")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GasEstimator_LimitTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitTransfer'
type GasEstimator_LimitTransfer_Call struct {
	*mock.Call
}

// LimitTransfer is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) LimitTransfer() *GasEstimator_LimitTransfer_Call {
	return &GasEstimator_LimitTransfer_Call{Call: _e.mock.On("LimitTransfer")}
}

func (_c *GasEstimator_LimitTransfer_Call) Run(run func()) *GasEstimator_LimitTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_LimitTransfer_Call) Return(_a0 uint64) *GasEstimator_LimitTransfer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_LimitTransfer_Call) RunAndReturn(run func() uint64) *GasEstimator_LimitTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// Mode provides a mock function with given fields:
func (_m *GasEstimator) Mode() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Mode")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GasEstimator_Mode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mode'
type GasEstimator_Mode_Call struct {
	*mock.Call
}

// Mode is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) Mode() *GasEstimator_Mode_Call {
	return &GasEstimator_Mode_Call{Call: _e.mock.On("Mode")}
}

func (_c *GasEstimator_Mode_Call) Run(run func()) *GasEstimator_Mode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_Mode_Call) Return(_a0 string) *GasEstimator_Mode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_Mode_Call) RunAndReturn(run func() string) *GasEstimator_Mode_Call {
	_c.Call.Return(run)
	return _c
}

// PriceDefault provides a mock function with given fields:
func (_m *GasEstimator) PriceDefault() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PriceDefault")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_PriceDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PriceDefault'
type GasEstimator_PriceDefault_Call struct {
	*mock.Call
}

// PriceDefault is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) PriceDefault() *GasEstimator_PriceDefault_Call {
	return &GasEstimator_PriceDefault_Call{Call: _e.mock.On("PriceDefault")}
}

func (_c *GasEstimator_PriceDefault_Call) Run(run func()) *GasEstimator_PriceDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_PriceDefault_Call) Return(_a0 *assets.Wei) *GasEstimator_PriceDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_PriceDefault_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_PriceDefault_Call {
	_c.Call.Return(run)
	return _c
}

// PriceMax provides a mock function with given fields:
func (_m *GasEstimator) PriceMax() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PriceMax")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_PriceMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PriceMax'
type GasEstimator_PriceMax_Call struct {
	*mock.Call
}

// PriceMax is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) PriceMax() *GasEstimator_PriceMax_Call {
	return &GasEstimator_PriceMax_Call{Call: _e.mock.On("PriceMax")}
}

func (_c *GasEstimator_PriceMax_Call) Run(run func()) *GasEstimator_PriceMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_PriceMax_Call) Return(_a0 *assets.Wei) *GasEstimator_PriceMax_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_PriceMax_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_PriceMax_Call {
	_c.Call.Return(run)
	return _c
}

// PriceMaxKey provides a mock function with given fields: _a0
func (_m *GasEstimator) PriceMaxKey(_a0 common.Address) *assets.Wei {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for PriceMaxKey")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Wei); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_PriceMaxKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PriceMaxKey'
type GasEstimator_PriceMaxKey_Call struct {
	*mock.Call
}

// PriceMaxKey is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *GasEstimator_Expecter) PriceMaxKey(_a0 interface{}) *GasEstimator_PriceMaxKey_Call {
	return &GasEstimator_PriceMaxKey_Call{Call: _e.mock.On("PriceMaxKey", _a0)}
}

func (_c *GasEstimator_PriceMaxKey_Call) Run(run func(_a0 common.Address)) *GasEstimator_PriceMaxKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *GasEstimator_PriceMaxKey_Call) Return(_a0 *assets.Wei) *GasEstimator_PriceMaxKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_PriceMaxKey_Call) RunAndReturn(run func(common.Address) *assets.Wei) *GasEstimator_PriceMaxKey_Call {
	_c.Call.Return(run)
	return _c
}

// PriceMin provides a mock function with given fields:
func (_m *GasEstimator) PriceMin() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PriceMin")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_PriceMin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PriceMin'
type GasEstimator_PriceMin_Call struct {
	*mock.Call
}

// PriceMin is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) PriceMin() *GasEstimator_PriceMin_Call {
	return &GasEstimator_PriceMin_Call{Call: _e.mock.On("PriceMin")}
}

func (_c *GasEstimator_PriceMin_Call) Run(run func()) *GasEstimator_PriceMin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_PriceMin_Call) Return(_a0 *assets.Wei) *GasEstimator_PriceMin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_PriceMin_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_PriceMin_Call {
	_c.Call.Return(run)
	return _c
}

// TipCapDefault provides a mock function with given fields:
func (_m *GasEstimator) TipCapDefault() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TipCapDefault")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_TipCapDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TipCapDefault'
type GasEstimator_TipCapDefault_Call struct {
	*mock.Call
}

// TipCapDefault is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) TipCapDefault() *GasEstimator_TipCapDefault_Call {
	return &GasEstimator_TipCapDefault_Call{Call: _e.mock.On("TipCapDefault")}
}

func (_c *GasEstimator_TipCapDefault_Call) Run(run func()) *GasEstimator_TipCapDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_TipCapDefault_Call) Return(_a0 *assets.Wei) *GasEstimator_TipCapDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_TipCapDefault_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_TipCapDefault_Call {
	_c.Call.Return(run)
	return _c
}

// TipCapMin provides a mock function with given fields:
func (_m *GasEstimator) TipCapMin() *assets.Wei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TipCapMin")
	}

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// GasEstimator_TipCapMin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TipCapMin'
type GasEstimator_TipCapMin_Call struct {
	*mock.Call
}

// TipCapMin is a helper method to define mock.On call
func (_e *GasEstimator_Expecter) TipCapMin() *GasEstimator_TipCapMin_Call {
	return &GasEstimator_TipCapMin_Call{Call: _e.mock.On("TipCapMin")}
}

func (_c *GasEstimator_TipCapMin_Call) Run(run func()) *GasEstimator_TipCapMin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GasEstimator_TipCapMin_Call) Return(_a0 *assets.Wei) *GasEstimator_TipCapMin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GasEstimator_TipCapMin_Call) RunAndReturn(run func() *assets.Wei) *GasEstimator_TipCapMin_Call {
	_c.Call.Return(run)
	return _c
}

// NewGasEstimator creates a new instance of GasEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGasEstimator(t interface {
	mock.TestingT
	Cleanup(func())
}) *GasEstimator {
	mock := &GasEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
