package service

import (
	"context"
	"fmt"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
)

type user struct {
	repo *repository.Repository
}

func NewUser(repo *repository.Repository) *user {
	return &user{repo: repo}
}

func (u *user) CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error) {

	dbUser, err := u.repo.User.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service.user.CreateUser error: %w", err)
	}

	return dbUser, nil
}
