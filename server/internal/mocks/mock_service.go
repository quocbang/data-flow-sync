// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	repositories "github.com/quocbang/data-flow-sync/server/internal/repositories"
	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

type Services_Expecter struct {
	mock *mock.Mock
}

func (_m *Services) EXPECT() *Services_Expecter {
	return &Services_Expecter{mock: &_m.Mock}
}

// Account provides a mock function with given fields:
func (_m *Services) Account() repositories.AccountServices {
	ret := _m.Called()

	var r0 repositories.AccountServices
	if rf, ok := ret.Get(0).(func() repositories.AccountServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.AccountServices)
		}
	}

	return r0
}

// Services_Account_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Account'
type Services_Account_Call struct {
	*mock.Call
}

// Account is a helper method to define mock.On call
func (_e *Services_Expecter) Account() *Services_Account_Call {
	return &Services_Account_Call{Call: _e.mock.On("Account")}
}

func (_c *Services_Account_Call) Run(run func()) *Services_Account_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Services_Account_Call) Return(_a0 repositories.AccountServices) *Services_Account_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Services_Account_Call) RunAndReturn(run func() repositories.AccountServices) *Services_Account_Call {
	_c.Call.Return(run)
	return _c
}

// File provides a mock function with given fields:
func (_m *Services) File() repositories.FileServices {
	ret := _m.Called()

	var r0 repositories.FileServices
	if rf, ok := ret.Get(0).(func() repositories.FileServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.FileServices)
		}
	}

	return r0
}

// Services_File_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'File'
type Services_File_Call struct {
	*mock.Call
}

// File is a helper method to define mock.On call
func (_e *Services_Expecter) File() *Services_File_Call {
	return &Services_File_Call{Call: _e.mock.On("File")}
}

func (_c *Services_File_Call) Run(run func()) *Services_File_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Services_File_Call) Return(_a0 repositories.FileServices) *Services_File_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Services_File_Call) RunAndReturn(run func() repositories.FileServices) *Services_File_Call {
	_c.Call.Return(run)
	return _c
}

// MergeRequest provides a mock function with given fields:
func (_m *Services) MergeRequest() repositories.MergeRequestServices {
	ret := _m.Called()

	var r0 repositories.MergeRequestServices
	if rf, ok := ret.Get(0).(func() repositories.MergeRequestServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.MergeRequestServices)
		}
	}

	return r0
}

// Services_MergeRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeRequest'
type Services_MergeRequest_Call struct {
	*mock.Call
}

// MergeRequest is a helper method to define mock.On call
func (_e *Services_Expecter) MergeRequest() *Services_MergeRequest_Call {
	return &Services_MergeRequest_Call{Call: _e.mock.On("MergeRequest")}
}

func (_c *Services_MergeRequest_Call) Run(run func()) *Services_MergeRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Services_MergeRequest_Call) Return(_a0 repositories.MergeRequestServices) *Services_MergeRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Services_MergeRequest_Call) RunAndReturn(run func() repositories.MergeRequestServices) *Services_MergeRequest_Call {
	_c.Call.Return(run)
	return _c
}

// Station provides a mock function with given fields:
func (_m *Services) Station() repositories.StationServices {
	ret := _m.Called()

	var r0 repositories.StationServices
	if rf, ok := ret.Get(0).(func() repositories.StationServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.StationServices)
		}
	}

	return r0
}

// Services_Station_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Station'
type Services_Station_Call struct {
	*mock.Call
}

// Station is a helper method to define mock.On call
func (_e *Services_Expecter) Station() *Services_Station_Call {
	return &Services_Station_Call{Call: _e.mock.On("Station")}
}

func (_c *Services_Station_Call) Run(run func()) *Services_Station_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Services_Station_Call) Return(_a0 repositories.StationServices) *Services_Station_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Services_Station_Call) RunAndReturn(run func() repositories.StationServices) *Services_Station_Call {
	_c.Call.Return(run)
	return _c
}

// StationGroup provides a mock function with given fields:
func (_m *Services) StationGroup() repositories.StationGroupServices {
	ret := _m.Called()

	var r0 repositories.StationGroupServices
	if rf, ok := ret.Get(0).(func() repositories.StationGroupServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.StationGroupServices)
		}
	}

	return r0
}

// Services_StationGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StationGroup'
type Services_StationGroup_Call struct {
	*mock.Call
}

// StationGroup is a helper method to define mock.On call
func (_e *Services_Expecter) StationGroup() *Services_StationGroup_Call {
	return &Services_StationGroup_Call{Call: _e.mock.On("StationGroup")}
}

func (_c *Services_StationGroup_Call) Run(run func()) *Services_StationGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Services_StationGroup_Call) Return(_a0 repositories.StationGroupServices) *Services_StationGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Services_StationGroup_Call) RunAndReturn(run func() repositories.StationGroupServices) *Services_StationGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
