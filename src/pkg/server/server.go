package server

import (
	"context"
	"errors"
	"fmt"
	"loggingapi/src/pkg/common"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"loggingapi/src/pkg/config"
	"loggingapi/src/pkg/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	engine *chi.Mux
	config *config.Config
	logger *logrus.Logger
}

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	rateLimitMiddleware, err := middleware.RateLimitMiddleware(conf.Server.LimitConfigs)
	if err != nil {
		return nil, err
	}

	//init database (need for repository)

	//init redis?

	//init repository

	c := chi.NewMux()
	c.Use(
		rateLimitMiddleware,
	)

	return &Server{
		engine: c,
		config: conf,
		logger: logger,
	}, nil
}

func (s *Server) Close() {
	//close repository
}

func (s *Server) initRoutes() {
	root := s.engine

	root.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		common.ResponseSuccess(w, make(map[string]interface{}))
	})
}

func (s *Server) Run() error {
	defer s.Close()

	s.initRoutes()

	addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)
	s.logger.Infof("starting server at %s", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("failed to start server: %s", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(s.config.Server.GracefulShutdownPeriod)*time.Second)
	defer cancel()

	ch := <-sig
	s.logger.Infof("received signal: %s", ch)

	return server.Shutdown(ctx)
}
