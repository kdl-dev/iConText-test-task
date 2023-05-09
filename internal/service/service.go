package service

import (
	"context"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
)

// go:generate go run github.com/vektra/mockery/v2@v2.26.0 --name MathOperation
type MathOperation interface {
	Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error)
}

// go:generate go run github.com/vektra/mockery/v2@v2.26.0 --name User
type User interface {
	CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error)
}

type Signature interface {
	SHA512Sign(sha512Input *entity.HMACSHA512DTO) *entity.Signature
}

type Service struct {
	MathOperation
	User
	Signature
}

var (
	Services *Service
)

func init() {
	Services = NewService(repository.Repo)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		MathOperation: NewMathOperationService(repo),
		User:          NewUserService(repo),
		Signature:     NewSignatureService(),
	}
}
