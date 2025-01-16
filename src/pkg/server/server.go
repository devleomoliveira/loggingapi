package server

import (
	"loggingapi/src/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	engine *chi.Mux
	config *config.Config
	logger *logrus.Logger
}

func New(config *config.Config, logger *logrus.Logger) (*Server, error) {
	return nil, nil
}

func Run(s *Server) error {
	return nil
}
