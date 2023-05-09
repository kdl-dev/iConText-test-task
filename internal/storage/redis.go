package storage

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kdl-dev/iConText-test-task/internal/config"
	"github.com/kdl-dev/iConText-test-task/logger"
	"github.com/redis/go-redis/v9"
)

type RedisPool struct {
	Pool               *redis.Client
	closeSignal        chan bool
	successCloseSignal chan bool
}

var (
	redisHost *string // flag
	redisPort *string // flag

	RedisOpt *redis.Options
	Redis    *RedisPool
)

func init() {
	redisHost = flag.String("host", config.AppCfg.Redis.Host, "redis host")
	redisPort = flag.String("port", config.AppCfg.Redis.Port, "redis port")
	flag.Parse()

	RedisOpt = &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisHost, *redisPort),
		DB:       config.AppCfg.Redis.DB,
		Username: config.AppCfg.Redis.User,
		Password: config.AppCfg.Redis.Password,
	}

	var err error
	Redis, err = NewRedisPool(context.Background(), RedisOpt)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewRedisPool(ctx context.Context, opt *redis.Options) (*RedisPool, error) {
	cl := redis.NewClient(opt)

	statusCmd := cl.Ping(ctx)
	if err := statusCmd.Err(); err != nil {
		return nil, fmt.Errorf("redis ping error: %w", err)
	}

	logger.Log.Info(fmt.Sprintf("redis run on %s", opt.Addr))

	redisPool := &RedisPool{
		Pool:               redis.NewClient(opt),
		closeSignal:        make(chan bool, 1),
		successCloseSignal: make(chan bool, 1),
	}

	go redisPool.KeepAlive()

	return redisPool, nil
}

func (rds *RedisPool) KeepAlive() {
	var err error
	var isLost bool = false
	var status *redis.StatusCmd
	for {
		select {
		case <-rds.closeSignal:
			if err = rds.Pool.Close(); err != nil {
				logger.Log.Error(fmt.Errorf("redis connection pool close error: %w", err).Error())
				return
			}
			logger.Log.Info("redis connection pool success closed")
			rds.successCloseSignal <- true
			return
		default:
			time.Sleep(time.Second * time.Duration(config.AppCfg.Redis.KeepAlivePoolPeriod))

			status = rds.Pool.Ping(context.Background())
			if status.Err() == nil {
				if isLost {
					logger.Log.Info("redis success reconnected")
					isLost = false
				}
				continue
			}

			isLost = true
			logger.Log.Warn("lost redis connection. Restoring...")
			logger.Log.Error(fmt.Errorf("redis ping error: %w", status.Err()).Error())

			rds.Pool = redis.NewClient(rds.Pool.Options())
		}
	}
}

func (rds *RedisPool) Close() {
	rds.closeSignal <- true
	<-rds.successCloseSignal
}
