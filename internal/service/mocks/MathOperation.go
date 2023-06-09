// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/kdl-dev/iConText-test-task/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MathOperation is an autogenerated mock type for the MathOperation type
type MathOperation struct {
	mock.Mock
}

// Increment provides a mock function with given fields: ctx, incrementInput
func (_m *MathOperation) Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error) {
	ret := _m.Called(ctx, incrementInput)

	var r0 *entity.IncrementResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.IncrementDTO) (*entity.IncrementResult, error)); ok {
		return rf(ctx, incrementInput)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.IncrementDTO) *entity.IncrementResult); ok {
		r0 = rf(ctx, incrementInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.IncrementResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.IncrementDTO) error); ok {
		r1 = rf(ctx, incrementInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMathOperation interface {
	mock.TestingT
	Cleanup(func())
}

// NewMathOperation creates a new instance of MathOperation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMathOperation(t mockConstructorTestingTNewMathOperation) *MathOperation {
	mock := &MathOperation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
