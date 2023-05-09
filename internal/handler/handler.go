package handler

import (
	"context"
	"net/http"

	"github.com/kdl-dev/iConText-test-task/internal/middleware"
	"github.com/kdl-dev/iConText-test-task/internal/service"
	"github.com/kdl-dev/iConText-test-task/logger"
	"github.com/kdl-dev/iConText-test-task/pkg/utils"
)

var (
	jsonBindErr = "json bind error"
	jsonSendErr = "json send error"

	Context = context.Background()
)

type Handler struct {
	ctx      context.Context
	services *service.Service
	Logger   *logger.Logger
	GoId     uint64
}

var (
	Handlers *Handler
	Router   *http.ServeMux
)

func init() {
	Handlers = NewHandler(Context, service.Services, logger.Log)
	Router = http.NewServeMux()
	Handlers.InitRoutes(Router)
}

func (h *Handler) handleError(w http.ResponseWriter, status int, err error) error {
	h.Logger.Error(err.Error())

	err = utils.SendJSON(w, status, utils.JSON{"err": err.Error()})
	if err != nil {
		h.Logger.Error(err.Error())
		return err
	}

	return nil
}

func NewHandler(ctx context.Context, services *service.Service, log *logger.Logger) *Handler {
	return &Handler{
		ctx:      ctx,
		services: services,
		Logger:   log,
	}
}

func (h *Handler) InitRoutes(r *http.ServeMux) {
	var (
		redisPrefix    = "/redis"
		postgresPrefix = "/postgres"
		signPrefix     = "/sign"

		incrementPath  = redisPrefix + "/incr"
		usersPath      = postgresPrefix + "/users"
		sha512SignPath = signPrefix + "/hmacsha512"
	)

	allowedMethod := http.MethodPost

	r.Handle(incrementPath, middleware.OnlyAllowedMethod(allowedMethod, h.increment))
	r.Handle(usersPath, middleware.OnlyAllowedMethod(allowedMethod, h.createUser))
	r.HandleFunc(sha512SignPath, middleware.OnlyAllowedMethod(allowedMethod, h.sha512Sign))
}
