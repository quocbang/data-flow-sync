// Code generated by mockery v2.32.0. DO NOT EDIT.

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

// Authorization provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) Authorization(_a0 context.Context, _a1 string) (*models.JwtCustomClaims, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.JwtCustomClaims
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.JwtCustomClaims, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.JwtCustomClaims); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JwtCustomClaims)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountServices_Authorization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authorization'
type AccountServices_Authorization_Call struct {
	*mock.Call
}

// Authorization is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountServices_Expecter) Authorization(_a0 interface{}, _a1 interface{}) *AccountServices_Authorization_Call {
	return &AccountServices_Authorization_Call{Call: _e.mock.On("Authorization", _a0, _a1)}
}

func (_c *AccountServices_Authorization_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountServices_Authorization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountServices_Authorization_Call) Return(_a0 *models.JwtCustomClaims, _a1 error) *AccountServices_Authorization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountServices_Authorization_Call) RunAndReturn(run func(context.Context, string) (*models.JwtCustomClaims, error)) *AccountServices_Authorization_Call {
	_c.Call.Return(run)
	return _c
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

// SignIn provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) SignIn(_a0 context.Context, _a1 repositories.SignInRequest) (repositories.SignInReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 repositories.SignInReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repositories.SignInRequest) (repositories.SignInReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repositories.SignInRequest) repositories.SignInReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repositories.SignInReply)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repositories.SignInRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountServices_SignIn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignIn'
type AccountServices_SignIn_Call struct {
	*mock.Call
}

// SignIn is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositories.SignInRequest
func (_e *AccountServices_Expecter) SignIn(_a0 interface{}, _a1 interface{}) *AccountServices_SignIn_Call {
	return &AccountServices_SignIn_Call{Call: _e.mock.On("SignIn", _a0, _a1)}
}

func (_c *AccountServices_SignIn_Call) Run(run func(_a0 context.Context, _a1 repositories.SignInRequest)) *AccountServices_SignIn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositories.SignInRequest))
	})
	return _c
}

func (_c *AccountServices_SignIn_Call) Return(_a0 repositories.SignInReply, _a1 error) *AccountServices_SignIn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountServices_SignIn_Call) RunAndReturn(run func(context.Context, repositories.SignInRequest) (repositories.SignInReply, error)) *AccountServices_SignIn_Call {
	_c.Call.Return(run)
	return _c
}

// SignOut provides a mock function with given fields: _a0, _a1
func (_m *AccountServices) SignOut(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountServices_SignOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignOut'
type AccountServices_SignOut_Call struct {
	*mock.Call
}

// SignOut is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountServices_Expecter) SignOut(_a0 interface{}, _a1 interface{}) *AccountServices_SignOut_Call {
	return &AccountServices_SignOut_Call{Call: _e.mock.On("SignOut", _a0, _a1)}
}

func (_c *AccountServices_SignOut_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountServices_SignOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountServices_SignOut_Call) Return(_a0 error) *AccountServices_SignOut_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountServices_SignOut_Call) RunAndReturn(run func(context.Context, string) error) *AccountServices_SignOut_Call {
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