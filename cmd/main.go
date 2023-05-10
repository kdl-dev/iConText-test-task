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
	"github.com/kdl-dev/iConText-test-task/internal/config"
	"github.com/kdl-dev/iConText-test-task/internal/handler"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
	"github.com/kdl-dev/iConText-test-task/internal/service"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
	"github.com/kdl-dev/iConText-test-task/logger"
	"github.com/kdl-dev/iConText-test-task/server"
	"github.com/redis/go-redis/v9"
)

var (
	envConfigFilePath    = ".env"
	loggerConfigFilePath = "config/logger.yaml"

	httpServer *http.Server
	appServer  *server.Server
	Handlers   *handler.Handler
	Router     *http.ServeMux
	Context    = context.Background()
	Services   *service.Service
	Repo       *repository.Repository

	PostgresOpt *storage.PostgresOptions
	Postgres    *storage.PostgresPool
	redisHost   *string // flag
	redisPort   *string // flag

	RedisOpt *redis.Options
	Redis    *storage.RedisPool
)

func init() {
	var err error
	config.AppCfg = &config.App{}

	if err := cleanenv.ReadConfig(envConfigFilePath, config.AppCfg); err != nil {
		log.Fatal(err)
	}

	if err := cleanenv.ReadConfig(loggerConfigFilePath, &config.AppCfg.Logger); err != nil {
		log.Fatal(err)
	}

	if logger.Log, err = logger.NewLogger(&config.AppCfg.Logger); err != nil {
		log.Fatal(err)
	}

	redisHost = flag.String("host", config.AppCfg.Redis.Host, "redis host")
	redisPort = flag.String("port", config.AppCfg.Redis.Port, "redis port")
	flag.Parse()

	RedisOpt = &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisHost, *redisPort),
		DB:       config.AppCfg.Redis.DB,
		Username: config.AppCfg.Redis.User,
		Password: config.AppCfg.Redis.Password,
	}

	Redis, err = storage.NewRedisPool(context.Background(), RedisOpt)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	PostgresOpt = &storage.PostgresOptions{
		Role:     config.AppCfg.Postgres.Role,
		Password: config.AppCfg.Postgres.Password,
		Host:     fmt.Sprintf("%s:%s", config.AppCfg.Postgres.Host, config.AppCfg.Postgres.Port),
		DBName:   config.AppCfg.Postgres.DBName,
		SSLMode:  config.AppCfg.Postgres.SSLMode,
	}

	Postgres, err = storage.NewPostgresPool(context.Background(), PostgresOpt)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	if err = Postgres.InitTables(os.Getenv("PG_INIT_TABLES_FILE_PATH")); err != nil {
		logger.Log.Fatal(err.Error())
	}

	Repo = repository.NewRepository(Postgres, Redis)
	Services = service.NewService(Repo)
	Handlers = handler.NewHandler(Context, Services, logger.Log)
	Router = http.NewServeMux()
	Handlers.InitRoutes(Router)

	httpServer = new(http.Server)
	httpServer.Addr = fmt.Sprintf("%s:%s", config.AppCfg.Server.Addr, config.AppCfg.Server.Port)
	httpServer.Handler = Router

}

func main() {
	appServer = server.NewServer(httpServer)
	isGracefullShutdown := make(chan bool, 1)

	go func() {
		logger.Log.Info(fmt.Sprintf("http server run on %s", httpServer.Addr))
		if err := appServer.ListenAndServe(); err != nil {
			select {
			case <-isGracefullShutdown:
				logger.Log.Info(err.Error())
				return
			default:
				logger.Log.Fatal(err.Error())
			}
		}
	}()

	signals := make(chan os.Signal, 1)
	defer close(signals)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	isGracefullShutdown <- true
	logger.Log.Info("gracefull shutdown...")

	cleanUpAfterYourself()
}

func cleanUpAfterYourself() {
	appServer.Shutdown(context.Background())
	Postgres.Close()
	Redis.Close()
}
