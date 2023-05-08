package service

import (
	"context"
	"fmt"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
)

type mathOperationService struct {
	repo *repository.Repository
}

func NewMathOperationService(repo *repository.Repository) *mathOperationService {
	return &mathOperationService{repo: repo}
}

func (m *mathOperationService) Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error) {

	incrementResult, err := m.repo.MathOperation.Increment(ctx, incrementInput)
	if err != nil {
		return nil, fmt.Errorf("service.math.Increment error: %w", err)
	}

	return incrementResult, nil
}
