// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/dwarvesf/bookstore-api/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

type Controller_Expecter struct {
	mock *mock.Mock
}

func (_m *Controller) EXPECT() *Controller_Expecter {
	return &Controller_Expecter{mock: &_m.Mock}
}

// CreateBook provides a mock function with given fields: ctx, _a1
func (_m *Controller) CreateBook(ctx context.Context, _a1 model.CreateBookRequest) (*model.Book, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateBookRequest) (*model.Book, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateBookRequest) *model.Book); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateBookRequest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_CreateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBook'
type Controller_CreateBook_Call struct {
	*mock.Call
}

// CreateBook is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 model.CreateBookRequest
func (_e *Controller_Expecter) CreateBook(ctx interface{}, _a1 interface{}) *Controller_CreateBook_Call {
	return &Controller_CreateBook_Call{Call: _e.mock.On("CreateBook", ctx, _a1)}
}

func (_c *Controller_CreateBook_Call) Run(run func(ctx context.Context, _a1 model.CreateBookRequest)) *Controller_CreateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateBookRequest))
	})
	return _c
}

func (_c *Controller_CreateBook_Call) Return(_a0 *model.Book, _a1 error) *Controller_CreateBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_CreateBook_Call) RunAndReturn(run func(context.Context, model.CreateBookRequest) (*model.Book, error)) *Controller_CreateBook_Call {
	_c.Call.Return(run)
	return _c
}

// GetBooks provides a mock function with given fields: ctx, q, topicID
func (_m *Controller) GetBooks(ctx context.Context, q model.ListQuery, topicID int) (*model.ListResult[model.Book], error) {
	ret := _m.Called(ctx, q, topicID)

	var r0 *model.ListResult[model.Book]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ListQuery, int) (*model.ListResult[model.Book], error)); ok {
		return rf(ctx, q, topicID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.ListQuery, int) *model.ListResult[model.Book]); ok {
		r0 = rf(ctx, q, topicID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ListResult[model.Book])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.ListQuery, int) error); ok {
		r1 = rf(ctx, q, topicID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_GetBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooks'
type Controller_GetBooks_Call struct {
	*mock.Call
}

// GetBooks is a helper method to define mock.On call
//   - ctx context.Context
//   - q model.ListQuery
//   - topicID int
func (_e *Controller_Expecter) GetBooks(ctx interface{}, q interface{}, topicID interface{}) *Controller_GetBooks_Call {
	return &Controller_GetBooks_Call{Call: _e.mock.On("GetBooks", ctx, q, topicID)}
}

func (_c *Controller_GetBooks_Call) Run(run func(ctx context.Context, q model.ListQuery, topicID int)) *Controller_GetBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.ListQuery), args[2].(int))
	})
	return _c
}

func (_c *Controller_GetBooks_Call) Return(_a0 *model.ListResult[model.Book], _a1 error) *Controller_GetBooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_GetBooks_Call) RunAndReturn(run func(context.Context, model.ListQuery, int) (*model.ListResult[model.Book], error)) *Controller_GetBooks_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBook provides a mock function with given fields: ctx, _a1
func (_m *Controller) UpdateBook(ctx context.Context, _a1 model.UpdateBookRequest) (*model.Book, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateBookRequest) (*model.Book, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateBookRequest) *model.Book); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateBookRequest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_UpdateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBook'
type Controller_UpdateBook_Call struct {
	*mock.Call
}

// UpdateBook is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 model.UpdateBookRequest
func (_e *Controller_Expecter) UpdateBook(ctx interface{}, _a1 interface{}) *Controller_UpdateBook_Call {
	return &Controller_UpdateBook_Call{Call: _e.mock.On("UpdateBook", ctx, _a1)}
}

func (_c *Controller_UpdateBook_Call) Run(run func(ctx context.Context, _a1 model.UpdateBookRequest)) *Controller_UpdateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UpdateBookRequest))
	})
	return _c
}

func (_c *Controller_UpdateBook_Call) Return(_a0 *model.Book, _a1 error) *Controller_UpdateBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_UpdateBook_Call) RunAndReturn(run func(context.Context, model.UpdateBookRequest) (*model.Book, error)) *Controller_UpdateBook_Call {
	_c.Call.Return(run)
	return _c
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
