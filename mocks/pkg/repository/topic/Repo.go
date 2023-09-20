// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	db "github.com/dwarvesf/bookstore-api/pkg/repository/db"
	mock "github.com/stretchr/testify/mock"
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
