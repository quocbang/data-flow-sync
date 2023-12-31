// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	repositories "github.com/quocbang/data-flow-sync/server/internal/repositories"
	models "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	mock "github.com/stretchr/testify/mock"
)

// AccountServices is an autogenerated mock type for the AccountServices type
type AccountServices struct {
	mock.Mock
}

type AccountServices_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountServices) EXPECT() *AccountServices_Expecter {
	return &AccountServices_Expecter{mock: &_m.Mock}
}

// DeleteAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) DeleteAccount(_a0 context.Context, _a1 repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 repositories.CommonUpdateAndDeleteReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repositories.DeleteAccountRequest) repositories.CommonUpdateAndDeleteReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repositories.CommonUpdateAndDeleteReply)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repositories.DeleteAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountServices_DeleteAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccount'
type AccountServices_DeleteAccount_Call struct {
	*mock.Call
}

// DeleteAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositories.DeleteAccountRequest
func (_e *AccountServices_Expecter) DeleteAccount(_a0 interface{}, _a1 interface{}) *AccountServices_DeleteAccount_Call {
	return &AccountServices_DeleteAccount_Call{Call: _e.mock.On("DeleteAccount", _a0, _a1)}
}

func (_c *AccountServices_DeleteAccount_Call) Run(run func(_a0 context.Context, _a1 repositories.DeleteAccountRequest)) *AccountServices_DeleteAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositories.DeleteAccountRequest))
	})
	return _c
}

func (_c *AccountServices_DeleteAccount_Call) Return(_a0 repositories.CommonUpdateAndDeleteReply, _a1 error) *AccountServices_DeleteAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountServices_DeleteAccount_Call) RunAndReturn(run func(context.Context, repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error)) *AccountServices_DeleteAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) GetAccount(_a0 context.Context, _a1 string) (models.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountServices_GetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccount'
type AccountServices_GetAccount_Call struct {
	*mock.Call
}

// GetAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountServices_Expecter) GetAccount(_a0 interface{}, _a1 interface{}) *AccountServices_GetAccount_Call {
	return &AccountServices_GetAccount_Call{Call: _e.mock.On("GetAccount", _a0, _a1)}
}

func (_c *AccountServices_GetAccount_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountServices_GetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountServices_GetAccount_Call) Return(_a0 models.Account, _a1 error) *AccountServices_GetAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountServices_GetAccount_Call) RunAndReturn(run func(context.Context, string) (models.Account, error)) *AccountServices_GetAccount_Call {
	_c.Call.Return(run)
	return _c
}

// SignUp provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) SignUp(_a0 context.Context, _a1 repositories.SignUpAccountRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repositories.SignUpAccountRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountServices_SignUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignUp'
type AccountServices_SignUp_Call struct {
	*mock.Call
}

// SignUp is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositories.SignUpAccountRequest
func (_e *AccountServices_Expecter) SignUp(_a0 interface{}, _a1 interface{}) *AccountServices_SignUp_Call {
	return &AccountServices_SignUp_Call{Call: _e.mock.On("SignUp", _a0, _a1)}
}

func (_c *AccountServices_SignUp_Call) Run(run func(_a0 context.Context, _a1 repositories.SignUpAccountRequest)) *AccountServices_SignUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositories.SignUpAccountRequest))
	})
	return _c
}

func (_c *AccountServices_SignUp_Call) Return(_a0 error) *AccountServices_SignUp_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountServices_SignUp_Call) RunAndReturn(run func(context.Context, repositories.SignUpAccountRequest) error) *AccountServices_SignUp_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateToUserRole provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) UpdateToUserRole(_a0 context.Context, _a1 string) (repositories.CommonUpdateAndDeleteReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 repositories.CommonUpdateAndDeleteReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (repositories.CommonUpdateAndDeleteReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) repositories.CommonUpdateAndDeleteReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repositories.CommonUpdateAndDeleteReply)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountServices_UpdateToUserRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateToUserRole'
type AccountServices_UpdateToUserRole_Call struct {
	*mock.Call
}

// UpdateToUserRole is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountServices_Expecter) UpdateToUserRole(_a0 interface{}, _a1 interface{}) *AccountServices_UpdateToUserRole_Call {
	return &AccountServices_UpdateToUserRole_Call{Call: _e.mock.On("UpdateToUserRole", _a0, _a1)}
}

func (_c *AccountServices_UpdateToUserRole_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountServices_UpdateToUserRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountServices_UpdateToUserRole_Call) Return(_a0 repositories.CommonUpdateAndDeleteReply, _a1 error) *AccountServices_UpdateToUserRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountServices_UpdateToUserRole_Call) RunAndReturn(run func(context.Context, string) (repositories.CommonUpdateAndDeleteReply, error)) *AccountServices_UpdateToUserRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountServices creates a new instance of AccountServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountServices {
	mock := &AccountServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
