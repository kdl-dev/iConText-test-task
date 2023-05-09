package repository

import (
	"context"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
)

type MathOperation interface {
	Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error)
}

type User interface {
	CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error)
}

type Repository struct {
	MathOperation
	User
}

var (
	Repo *Repository
)

func init() {
	Repo = NewRepository(storage.Postgres, storage.Redis)
}

func NewRepository(pgPool *storage.PostgresPool, redisPool *storage.RedisPool) *Repository {
	return &Repository{
		MathOperation: NewMathOperationRepo(redisPool),
		User:          NewUserRepo(pgPool),
	}
}
