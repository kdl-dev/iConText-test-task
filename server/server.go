package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kdl-dev/iConText-test-task/logger"
)

type Server struct {
	pool *http.Server
}

func NewServer(pool *http.Server) *Server {
	return &Server{
		pool: pool,
	}
}

func (s *Server) ListenAndServe() error {
	return s.pool.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.pool.Shutdown(ctx); err != nil {
		logger.Log.Error(fmt.Errorf("http server connection pool close error: %w", err).Error())
	}
}
