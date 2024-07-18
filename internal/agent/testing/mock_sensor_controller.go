// Code generated by mockery v2.43.2. DO NOT EDIT.

package testing

import (
	context "context"

	sensor "github.com/joshuar/go-hass-agent/internal/hass/sensor"
	mock "github.com/stretchr/testify/mock"
)

// MockSensorController is an autogenerated mock type for the SensorController type
type MockSensorController struct {
	mock.Mock
}

type MockSensorController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSensorController) EXPECT() *MockSensorController_Expecter {
	return &MockSensorController_Expecter{mock: &_m.Mock}
}

// ActiveWorkers provides a mock function with given fields:
func (_m *MockSensorController) ActiveWorkers() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ActiveWorkers")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockSensorController_ActiveWorkers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActiveWorkers'
type MockSensorController_ActiveWorkers_Call struct {
	*mock.Call
}

// ActiveWorkers is a helper method to define mock.On call
func (_e *MockSensorController_Expecter) ActiveWorkers() *MockSensorController_ActiveWorkers_Call {
	return &MockSensorController_ActiveWorkers_Call{Call: _e.mock.On("ActiveWorkers")}
}

func (_c *MockSensorController_ActiveWorkers_Call) Run(run func()) *MockSensorController_ActiveWorkers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorController_ActiveWorkers_Call) Return(_a0 []string) *MockSensorController_ActiveWorkers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorController_ActiveWorkers_Call) RunAndReturn(run func() []string) *MockSensorController_ActiveWorkers_Call {
	_c.Call.Return(run)
	return _c
}

// InactiveWorkers provides a mock function with given fields:
func (_m *MockSensorController) InactiveWorkers() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InactiveWorkers")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockSensorController_InactiveWorkers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InactiveWorkers'
type MockSensorController_InactiveWorkers_Call struct {
	*mock.Call
}

// InactiveWorkers is a helper method to define mock.On call
func (_e *MockSensorController_Expecter) InactiveWorkers() *MockSensorController_InactiveWorkers_Call {
	return &MockSensorController_InactiveWorkers_Call{Call: _e.mock.On("InactiveWorkers")}
}

func (_c *MockSensorController_InactiveWorkers_Call) Run(run func()) *MockSensorController_InactiveWorkers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorController_InactiveWorkers_Call) Return(_a0 []string) *MockSensorController_InactiveWorkers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorController_InactiveWorkers_Call) RunAndReturn(run func() []string) *MockSensorController_InactiveWorkers_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx, name
func (_m *MockSensorController) Start(ctx context.Context, name string) (<-chan sensor.Details, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 <-chan sensor.Details
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (<-chan sensor.Details, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan sensor.Details); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan sensor.Details)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSensorController_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockSensorController_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockSensorController_Expecter) Start(ctx interface{}, name interface{}) *MockSensorController_Start_Call {
	return &MockSensorController_Start_Call{Call: _e.mock.On("Start", ctx, name)}
}

func (_c *MockSensorController_Start_Call) Run(run func(ctx context.Context, name string)) *MockSensorController_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSensorController_Start_Call) Return(_a0 <-chan sensor.Details, _a1 error) *MockSensorController_Start_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSensorController_Start_Call) RunAndReturn(run func(context.Context, string) (<-chan sensor.Details, error)) *MockSensorController_Start_Call {
	_c.Call.Return(run)
	return _c
}

// StartAll provides a mock function with given fields: ctx
func (_m *MockSensorController) StartAll(ctx context.Context) (<-chan sensor.Details, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StartAll")
	}

	var r0 <-chan sensor.Details
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan sensor.Details, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan sensor.Details); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan sensor.Details)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSensorController_StartAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartAll'
type MockSensorController_StartAll_Call struct {
	*mock.Call
}

// StartAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSensorController_Expecter) StartAll(ctx interface{}) *MockSensorController_StartAll_Call {
	return &MockSensorController_StartAll_Call{Call: _e.mock.On("StartAll", ctx)}
}

func (_c *MockSensorController_StartAll_Call) Run(run func(ctx context.Context)) *MockSensorController_StartAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSensorController_StartAll_Call) Return(_a0 <-chan sensor.Details, _a1 error) *MockSensorController_StartAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSensorController_StartAll_Call) RunAndReturn(run func(context.Context) (<-chan sensor.Details, error)) *MockSensorController_StartAll_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: name
func (_m *MockSensorController) Stop(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSensorController_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockSensorController_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - name string
func (_e *MockSensorController_Expecter) Stop(name interface{}) *MockSensorController_Stop_Call {
	return &MockSensorController_Stop_Call{Call: _e.mock.On("Stop", name)}
}

func (_c *MockSensorController_Stop_Call) Run(run func(name string)) *MockSensorController_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSensorController_Stop_Call) Return(_a0 error) *MockSensorController_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorController_Stop_Call) RunAndReturn(run func(string) error) *MockSensorController_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// StopAll provides a mock function with given fields:
func (_m *MockSensorController) StopAll() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StopAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSensorController_StopAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopAll'
type MockSensorController_StopAll_Call struct {
	*mock.Call
}

// StopAll is a helper method to define mock.On call
func (_e *MockSensorController_Expecter) StopAll() *MockSensorController_StopAll_Call {
	return &MockSensorController_StopAll_Call{Call: _e.mock.On("StopAll")}
}

func (_c *MockSensorController_StopAll_Call) Run(run func()) *MockSensorController_StopAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorController_StopAll_Call) Return(_a0 error) *MockSensorController_StopAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorController_StopAll_Call) RunAndReturn(run func() error) *MockSensorController_StopAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSensorController creates a new instance of MockSensorController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSensorController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSensorController {
	mock := &MockSensorController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}