package repository

import (
	"context"
	"fmt"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
)

var tableName = "users"

type user struct {
	pgPool *storage.PostgresPool
}

func NewUser(pgPool *storage.PostgresPool) *user {
	return &user{pgPool: pgPool}
}

func (m *user) CreateUser(ctx context.Context, user *entity.UserDTO) (*entity.User, error) {
	statement := `INSERT INTO ` + tableName + ` (name, age) VALUES($1, $2) RETURNING user_id, name, age;`
	var createUserResult entity.User

	row := m.pgPool.Pool.QueryRow(ctx, statement, user.Name, user.Age)
	if err := row.Scan(&createUserResult.ID, &createUserResult.Name, &createUserResult.Age); err != nil {
		return nil, fmt.Errorf("row scan error: %w", err)
	}

	return &createUserResult, nil
}
