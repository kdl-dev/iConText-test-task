package repository

import (
	"context"
	"fmt"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
	"github.com/redis/go-redis/v9"
)

type mathOperation struct {
	redisPool *storage.RedisPool
}

func NewMathOperation(redisPool *storage.RedisPool) *mathOperation {
	return &mathOperation{redisPool: redisPool}
}

func (m *mathOperation) Increment(ctx context.Context, incrementInput *entity.IncrementDTO) (*entity.IncrementResult, error) {
	var newValue entity.IncrementResult

	cmd := m.redisPool.Pool.Get(ctx, incrementInput.Key)
	if err := cmd.Scan(&newValue.Value); err != nil {
		if err != redis.Nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
	}

	newValue.Key = incrementInput.Key
	newValue.Value += *incrementInput.Value

	statusCMD := m.redisPool.Pool.Set(ctx, incrementInput.Key, newValue.Value, 0)
	if err := statusCMD.Err(); err != nil {
		return nil, fmt.Errorf("value set error: %w", err)
	}

	return &newValue, nil
}
