package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kdl-dev/iConText-test-task/internal/config"
	"github.com/kdl-dev/iConText-test-task/internal/handler"
	"github.com/kdl-dev/iConText-test-task/internal/storage"
	"github.com/kdl-dev/iConText-test-task/logger"
	"github.com/kdl-dev/iConText-test-task/server"
)

var (
	httpServer *http.Server
	appServer  *server.Server
)

func init() {
	httpServer = new(http.Server)
	httpServer.Addr = fmt.Sprintf("%s:%s", config.AppCfg.Server.Addr, config.AppCfg.Server.Port)
	httpServer.Handler = handler.Router
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
	storage.Postgres.Close()
	storage.Redis.Close()
}
