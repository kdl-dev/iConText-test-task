package service

import (
	"context"
	"errors"
	"testing"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
	"github.com/kdl-dev/iConText-test-task/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	var age int64 = 21

	user := &entity.UserDTO{
		Name: "Alex",
		Age:  &age,
	}

	tests := []struct {
		name       string
		input      *entity.UserDTO
		expections func(*mocks.User)
		err        error
	}{
		{
			name:  "success create user test",
			input: user,
			expections: func(mo *mocks.User) {
				mo.On("CreateUser", context.Background(), user).Return(&entity.User{}, nil)
			},
		},
		{
			name:  "storage error test",
			input: user,
			expections: func(mo *mocks.User) {
				mo.On("CreateUser", context.Background(), user).Return(nil, errors.New("some error"))
			},
			err: errors.New("service.user.CreateUser error: some error"),
		},
	}

	for _, test := range tests {
		t.Logf("running: %s", test.name)
		userRepo := mocks.NewUser(t)
		test.expections(userRepo)
		svc := NewUserService(&repository.Repository{User: userRepo})

		_, err := svc.CreateUser(context.Background(), test.input)
		if err != nil {
			if test.err != nil {
				assert.EqualError(t, test.err, err.Error())
			} else {
				t.Errorf("expected no error, found: %s", err.Error())
			}
		}

		userRepo.AssertExpectations(t)
	}
}
