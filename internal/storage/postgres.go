package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	Pool *pgxpool.Pool
}

type PostgresOptions struct {
	Role     string
	Password string
	Host     string
	DBName   string
	SSLMode  string
}

func NewPostgresPool(ctx context.Context, opt *PostgresOptions) (*PostgresPool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", opt.Role, opt.Password, opt.Host, opt.DBName, opt.SSLMode)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("postgres pool create error: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres ping error: %w", err)
	}

	return &PostgresPool{
		Pool: pool,
	}, nil
}

func (p *PostgresPool) InitTables(sqlFile string) error {
	statements, err := os.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("file read error: %w", err)
	}

	_, err = p.Pool.Exec(context.Background(), string(statements))
	if err != nil {
		return fmt.Errorf("sql execute error: %w", err)
	}

	return nil
}
