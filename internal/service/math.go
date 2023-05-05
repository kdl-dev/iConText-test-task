package service

import (
	"context"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
)

type mathOperation struct {
	repo *repository.Repository
}

func NewMathOperation(repo *repository.Repository) *mathOperation {
	return &mathOperation{repo: repo}
}

func (m *mathOperation) Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error) {
	return m.repo.MathOperation.Increment(ctx, incrementInput)
}
