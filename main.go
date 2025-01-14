package main

import (
	"flag"
	"os"

	"github.com/devleomoliveira/loggingapi/src/pkg/version"
	"loggingapi/src/pkg/config"

	"github.com/sirupsen/logrus"
)

var (
	printVersion = flag.Bool("v", false, "print version")
	appConfig    = flag.String("config", "config/app.yaml", "application config path")
)

func main() {
	flag.Parse()

	if *printVersion {
		version.Print()
		os.Exit(0)
	}

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	conf, err := config.Parse(*appConfig)
	if err != nil {
		logger.Fatalf("failed to parse application config")
	}
	println(conf)
	/*
		s, err := server.New(conf, logger)
		if err != nil {
			logger.Fatalf("failed to start server: %v", err)
		}

		if err := s.Run(); err != nil {
			logger.Fatalf("failed to start server: %v", err)
		}*/
}
