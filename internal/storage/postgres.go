package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kdl-dev/iConText-test-task/internal/config"
	"github.com/kdl-dev/iConText-test-task/logger"
)

type PostgresPool struct {
	Pool              *pgxpool.Pool
	closeSignal       chan bool
	succesCloseSignal chan bool
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

	logger.Log.Info(fmt.Sprintf("postgres run on %s", opt.Host))

	pgPool := &PostgresPool{
		Pool:              pool,
		closeSignal:       make(chan bool, 1),
		succesCloseSignal: make(chan bool, 1),
	}

	go pgPool.KeepAlive()

	return pgPool, nil
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

func (pg *PostgresPool) KeepAlive() {
	var err error
	var isLost bool = false
	for {
		select {
		case <-pg.closeSignal:
			pg.Pool.Close()
			logger.Log.Info("postgres connection pool success closed")
			pg.succesCloseSignal <- true
			return
		default:
			time.Sleep(time.Second * time.Duration(config.AppCfg.Postgres.KeepAlivePoolPeriod))

			if err = pg.Pool.Ping(context.Background()); err == nil {
				if isLost {
					logger.Log.Info("postgres success reconnected")
					isLost = false
				}

				continue
			}

			isLost = true
			logger.Log.Error(fmt.Errorf("postgres ping error: %w", err).Error())
			logger.Log.Warn("lost postgresql connection. Restoring...")

			pg.Pool, err = pgxpool.New(context.Background(), pg.Pool.Config().ConnString())
			if err != nil {
				logger.Log.Error(fmt.Errorf("postgres pool create error: %w", err).Error())
			}
		}
	}
}

func (pg *PostgresPool) Close() {
	pg.closeSignal <- true
	<-pg.succesCloseSignal
}
