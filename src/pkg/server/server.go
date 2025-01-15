package server

import (
	"loggingapi/src/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	engine *chi.Context
	config *config.Config
	logger *logrus.Logger

	//repository repository.Repository
	//controllers []controller.Controller
}

func new(config *config.Config, logger *logrus.Logger) (*Server, error) {
	return nil, nil
}
