// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	repositories "github.com/quocbang/data-flow-sync/server/internal/repositories"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

type Repositories_Expecter struct {
	mock *mock.Mock
}

func (_m *Repositories) EXPECT() *Repositories_Expecter {
	return &Repositories_Expecter{mock: &_m.Mock}
}

// Account provides a mock function with given fields:
func (_m *Repositories) Account() repositories.AccountServices {
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

// Repositories_Account_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Account'
type Repositories_Account_Call struct {
	*mock.Call
}

// Account is a helper method to define mock.On call
func (_e *Repositories_Expecter) Account() *Repositories_Account_Call {
	return &Repositories_Account_Call{Call: _e.mock.On("Account")}
}

func (_c *Repositories_Account_Call) Run(run func()) *Repositories_Account_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_Account_Call) Return(_a0 repositories.AccountServices) *Repositories_Account_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_Account_Call) RunAndReturn(run func() repositories.AccountServices) *Repositories_Account_Call {
	_c.Call.Return(run)
	return _c
}

// Begin provides a mock function with given fields: _a0, _a1
func (_m *Repositories) Begin(_a0 context.Context, _a1 ...*sql.TxOptions) (repositories.Repositories, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 repositories.Repositories
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*sql.TxOptions) (repositories.Repositories, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*sql.TxOptions) repositories.Repositories); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repositories.Repositories)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*sql.TxOptions) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repositories_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type Repositories_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...*sql.TxOptions
func (_e *Repositories_Expecter) Begin(_a0 interface{}, _a1 ...interface{}) *Repositories_Begin_Call {
	return &Repositories_Begin_Call{Call: _e.mock.On("Begin",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Repositories_Begin_Call) Run(run func(_a0 context.Context, _a1 ...*sql.TxOptions)) *Repositories_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*sql.TxOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*sql.TxOptions)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Repositories_Begin_Call) Return(_a0 repositories.Repositories, _a1 error) *Repositories_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repositories_Begin_Call) RunAndReturn(run func(context.Context, ...*sql.TxOptions) (repositories.Repositories, error)) *Repositories_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *Repositories) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repositories_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Repositories_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Repositories_Expecter) Close() *Repositories_Close_Call {
	return &Repositories_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Repositories_Close_Call) Run(run func()) *Repositories_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_Close_Call) Return(_a0 error) *Repositories_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_Close_Call) RunAndReturn(run func() error) *Repositories_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields:
func (_m *Repositories) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repositories_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type Repositories_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *Repositories_Expecter) Commit() *Repositories_Commit_Call {
	return &Repositories_Commit_Call{Call: _e.mock.On("Commit")}
}

func (_c *Repositories_Commit_Call) Run(run func()) *Repositories_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_Commit_Call) Return(_a0 error) *Repositories_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_Commit_Call) RunAndReturn(run func() error) *Repositories_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// RollBack provides a mock function with given fields:
func (_m *Repositories) RollBack() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repositories_RollBack_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RollBack'
type Repositories_RollBack_Call struct {
	*mock.Call
}

// RollBack is a helper method to define mock.On call
func (_e *Repositories_Expecter) RollBack() *Repositories_RollBack_Call {
	return &Repositories_RollBack_Call{Call: _e.mock.On("RollBack")}
}

func (_c *Repositories_RollBack_Call) Run(run func()) *Repositories_RollBack_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_RollBack_Call) Return(_a0 error) *Repositories_RollBack_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_RollBack_Call) RunAndReturn(run func() error) *Repositories_RollBack_Call {
	_c.Call.Return(run)
	return _c
}

// Station provides a mock function with given fields:
func (_m *Repositories) Station() repositories.StationServices {
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

// Repositories_Station_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Station'
type Repositories_Station_Call struct {
	*mock.Call
}

// Station is a helper method to define mock.On call
func (_e *Repositories_Expecter) Station() *Repositories_Station_Call {
	return &Repositories_Station_Call{Call: _e.mock.On("Station")}
}

func (_c *Repositories_Station_Call) Run(run func()) *Repositories_Station_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_Station_Call) Return(_a0 repositories.StationServices) *Repositories_Station_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_Station_Call) RunAndReturn(run func() repositories.StationServices) *Repositories_Station_Call {
	_c.Call.Return(run)
	return _c
}

// StationGroup provides a mock function with given fields:
func (_m *Repositories) StationGroup() repositories.StationGroupServices {
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

// Repositories_StationGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StationGroup'
type Repositories_StationGroup_Call struct {
	*mock.Call
}

// StationGroup is a helper method to define mock.On call
func (_e *Repositories_Expecter) StationGroup() *Repositories_StationGroup_Call {
	return &Repositories_StationGroup_Call{Call: _e.mock.On("StationGroup")}
}

func (_c *Repositories_StationGroup_Call) Run(run func()) *Repositories_StationGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repositories_StationGroup_Call) Return(_a0 repositories.StationGroupServices) *Repositories_StationGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repositories_StationGroup_Call) RunAndReturn(run func() repositories.StationGroupServices) *Repositories_StationGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepositories creates a new instance of Repositories. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositories(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repositories {
	mock := &Repositories{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
