package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisPool struct {
	Pool *redis.Client
}

func NewRedisPool(ctx context.Context, opt *redis.Options) (*RedisPool, error) {
	cl := redis.NewClient(opt)

	statusCmd := cl.Ping(ctx)
	if err := statusCmd.Err(); err != nil {
		return nil, fmt.Errorf("redis ping error: %w", err)
	}

	return &RedisPool{
		Pool: redis.NewClient(opt),
	}, nil
}
