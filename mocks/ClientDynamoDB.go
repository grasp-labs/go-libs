// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ClientDynamoDB is an autogenerated mock type for the ClientDynamoDB type
type ClientDynamoDB struct {
	mock.Mock
}

type ClientDynamoDB_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientDynamoDB) EXPECT() *ClientDynamoDB_Expecter {
	return &ClientDynamoDB_Expecter{mock: &_m.Mock}
}

// PutItem provides a mock function with given fields: ctx, table, itemToPut
func (_m *ClientDynamoDB) PutItem(ctx context.Context, table string, itemToPut interface{}) error {
	ret := _m.Called(ctx, table, itemToPut)

	if len(ret) == 0 {
		panic("no return value specified for PutItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, table, itemToPut)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientDynamoDB_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type ClientDynamoDB_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - table string
//   - itemToPut interface{}
func (_e *ClientDynamoDB_Expecter) PutItem(ctx interface{}, table interface{}, itemToPut interface{}) *ClientDynamoDB_PutItem_Call {
	return &ClientDynamoDB_PutItem_Call{Call: _e.mock.On("PutItem", ctx, table, itemToPut)}
}

func (_c *ClientDynamoDB_PutItem_Call) Run(run func(ctx context.Context, table string, itemToPut interface{})) *ClientDynamoDB_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *ClientDynamoDB_PutItem_Call) Return(_a0 error) *ClientDynamoDB_PutItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientDynamoDB_PutItem_Call) RunAndReturn(run func(context.Context, string, interface{}) error) *ClientDynamoDB_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, tableName, key, value
func (_m *ClientDynamoDB) Query(ctx context.Context, tableName string, key string, value string) ([]map[string]interface{}, error) {
	ret := _m.Called(ctx, tableName, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 []map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) ([]map[string]interface{}, error)); ok {
		return rf(ctx, tableName, key, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []map[string]interface{}); ok {
		r0 = rf(ctx, tableName, key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tableName, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientDynamoDB_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type ClientDynamoDB_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - tableName string
//   - key string
//   - value string
func (_e *ClientDynamoDB_Expecter) Query(ctx interface{}, tableName interface{}, key interface{}, value interface{}) *ClientDynamoDB_Query_Call {
	return &ClientDynamoDB_Query_Call{Call: _e.mock.On("Query", ctx, tableName, key, value)}
}

func (_c *ClientDynamoDB_Query_Call) Run(run func(ctx context.Context, tableName string, key string, value string)) *ClientDynamoDB_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientDynamoDB_Query_Call) Return(_a0 []map[string]interface{}, _a1 error) *ClientDynamoDB_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientDynamoDB_Query_Call) RunAndReturn(run func(context.Context, string, string, string) ([]map[string]interface{}, error)) *ClientDynamoDB_Query_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientDynamoDB creates a new instance of ClientDynamoDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientDynamoDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientDynamoDB {
	mock := &ClientDynamoDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
