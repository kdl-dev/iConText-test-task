// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/kdl-dev/iConText-test-task/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *User) CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserDTO) (*entity.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserDTO) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.UserDTO) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUser(t mockConstructorTestingTNewUser) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
