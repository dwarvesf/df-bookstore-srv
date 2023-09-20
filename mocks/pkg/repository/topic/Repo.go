// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	db "github.com/dwarvesf/bookstore-api/pkg/repository/db"
	mock "github.com/stretchr/testify/mock"

	model "github.com/dwarvesf/bookstore-api/pkg/model"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

type Repo_Expecter struct {
	mock *mock.Mock
}

func (_m *Repo) EXPECT() *Repo_Expecter {
	return &Repo_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repo) GetAll(ctx db.Context) ([]model.Topic, error) {
	ret := _m.Called(ctx)

	var r0 []model.Topic
	var r1 error
	if rf, ok := ret.Get(0).(func(db.Context) ([]model.Topic, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(db.Context) []model.Topic); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Topic)
		}
	}

	if rf, ok := ret.Get(1).(func(db.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repo_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Repo_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx db.Context
func (_e *Repo_Expecter) GetAll(ctx interface{}) *Repo_GetAll_Call {
	return &Repo_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *Repo_GetAll_Call) Run(run func(ctx db.Context)) *Repo_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Context))
	})
	return _c
}

func (_c *Repo_GetAll_Call) Return(_a0 []model.Topic, _a1 error) *Repo_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repo_GetAll_Call) RunAndReturn(run func(db.Context) ([]model.Topic, error)) *Repo_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// IsExist provides a mock function with given fields: ctx, ID
func (_m *Repo) IsExist(ctx db.Context, ID int) (bool, error) {
	ret := _m.Called(ctx, ID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(db.Context, int) (bool, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(db.Context, int) bool); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(db.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repo_IsExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsExist'
type Repo_IsExist_Call struct {
	*mock.Call
}

// IsExist is a helper method to define mock.On call
//   - ctx db.Context
//   - ID int
func (_e *Repo_Expecter) IsExist(ctx interface{}, ID interface{}) *Repo_IsExist_Call {
	return &Repo_IsExist_Call{Call: _e.mock.On("IsExist", ctx, ID)}
}

func (_c *Repo_IsExist_Call) Run(run func(ctx db.Context, ID int)) *Repo_IsExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Context), args[1].(int))
	})
	return _c
}

func (_c *Repo_IsExist_Call) Return(_a0 bool, _a1 error) *Repo_IsExist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repo_IsExist_Call) RunAndReturn(run func(db.Context, int) (bool, error)) *Repo_IsExist_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
