package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kdl-dev/iConText-test-task/internal/handler"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
	"github.com/kdl-dev/iConText-test-task/internal/service"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
	"github.com/kdl-dev/iConText-test-task/pkg/config"
	"github.com/kdl-dev/iConText-test-task/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var (
	loggerConfigFilePath = "config/logger.yaml"
	envConfigFilePath    = ".env"
	AppCfg               config.App

	redisHost *string
	redisPort *string

	postgresOpt storage.PostgresOptions
	redisOpt    redis.Options

	pgPool    *storage.PostgresPool
	redisPool *storage.RedisPool
)

func main() {
	Log, err := logger.NewLogger(&AppCfg.Logger)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	repo := repository.NewRepository(pgPool, redisPool)
	svc := service.NewService(repo)
	h := handler.NewHandler(ctx, svc, Log)
	router := http.NewServeMux()
	h.InitRoutes(router)

	server := new(http.Server)
	server.Addr = fmt.Sprintf("%s:%s", AppCfg.Server.Addr, AppCfg.Server.Port)
	server.Handler = router

	isGracefullShutdown := make(chan bool, 1)

	go func() {
		Log.Info(fmt.Sprintf("server run on %s", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			select {
			case <-isGracefullShutdown:
				Log.Info(err.Error())
				return
			default:
				Log.Fatal(err.Error())
			}
		}
	}()

	signals := make(chan os.Signal, 1)
	defer close(signals)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	isGracefullShutdown <- true
	Log.Info("gracefull shutdown...")

	if err := server.Shutdown(context.Background()); err != nil {
		Log.Info(fmt.Errorf("server shutdown error: %w", err).Error())
	}

	pgPool.Pool.Close()
	Log.Info("postgres connection pool closed")

	if err := redisPool.Pool.Close(); err != nil {
		Log.Info(fmt.Errorf("redis pool close error: %w", err).Error())
	} else {
		Log.Info("redis connection pool closed")
	}
}

func init() {
	if err := cleanenv.ReadConfig(envConfigFilePath, &AppCfg); err != nil {
		log.Fatal(err)
	}

	var logger config.Logger
	if err := cleanenv.ReadConfig(loggerConfigFilePath, &logger); err != nil {
		log.Fatal(err)
	}

	AppCfg.Logger = logger

	initStorageOptions()
	initConnPools()
}

func initStorageOptions() {
	redisHost = flag.String("host", AppCfg.Redis.Host, "redis host")
	redisPort = flag.String("port", AppCfg.Redis.Port, "redis port")
	flag.Parse()

	postgresOpt = storage.PostgresOptions{
		Role:     AppCfg.Postgres.Role,
		Password: AppCfg.Postgres.Password,
		Host:     AppCfg.Postgres.Host,
		DBName:   AppCfg.Postgres.DBName,
		SSLMode:  AppCfg.Postgres.SSLMode,
	}

	redisOpt = redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisHost, *redisPort),
		DB:       AppCfg.Redis.DB,
		Username: AppCfg.Redis.User,
		Password: AppCfg.Redis.Password,
	}
}

func initConnPools() {
	var err error
	pgPool, err = storage.NewPostgresPool(context.Background(), &postgresOpt)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = pgPool.InitTables(os.Getenv("PG_INIT_TABLES_FILE_PATH")); err != nil {
		log.Fatal(err.Error())
	}

	redisPool, err = storage.NewRedisPool(context.Background(), &redisOpt)
	if err != nil {
		log.Fatal(err.Error())
	}
}
