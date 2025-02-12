package server

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"loggingapi/src/pkg/database"
	"loggingapi/src/pkg/model"
	"loggingapi/src/pkg/repository"
	"loggingapi/src/routes"
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
	engine     *chi.Mux
	config     *config.Config
	logger     *logrus.Logger
	repository repository.Repository
}

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	rateLimitMiddleware, err := middleware.RateLimitMiddleware(conf.Server.LimitConfigs)
	if err != nil {
		return nil, err
	}

	db, err := database.NewMongoDB(&conf.DB)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to MongoDB")

	repository := repository.New(db)

	repository.Log().Store(model.Log{
		UUID:     "e1597e3c-1b96-4a96-b42c-45872a8b887t",
		Message:  "teste",
		Name:     "teste",
		Date:     primitive.NewDateTimeFromTime(time.Now()),
		Category: "teste",
		Level:    "info",
		Tags:     []string{"teste", "teste"},
		Trace:    []string{"1 teste", "2 teste"},
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
	})
	//repository.Log().Get("")

	c := chi.NewMux()
	c.Use(
		rateLimitMiddleware,
	)

	return &Server{
		engine:     c,
		config:     conf,
		logger:     logger,
		repository: repository,
	}, nil
}

func (s *Server) Close() {
	//close repository
}

func (s *Server) initRoutes() {
	root := s.engine
	conf := s.config
	routes.Web(root, *conf)
	routes.Api(root, *conf)
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
