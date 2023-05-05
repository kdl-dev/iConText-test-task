package service

import (
	"context"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
)

type MathOperation interface {
	Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error)
}

type User interface {
	CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error)
}

type Signature interface {
	SHA512Sign(sha512Input *entity.HMACSHA512DTO) (*entity.Signature, error)
}

type Service struct {
	MathOperation
	User
	Signature
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		MathOperation: NewMathOperation(repo),
		User:          NewUser(repo),
		Signature:     NewSignature(),
	}
}
