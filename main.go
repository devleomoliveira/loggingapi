package main

import (
	"flag"
	"os"

	"loggingapi/src/pkg/config"
	"loggingapi/src/pkg/server"
	"loggingapi/src/pkg/version"

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

	s, err := server.New(conf, logger)
	if err != nil {
		logger.Fatalf("failed to start server %v", err)
	}

	if err := s.Run(); err != nil {
		logger.Fatalf("failed to start server %v", err)
	}
	/*

		client, _ := database.New(&conf.DB)

			err = client.Ping(context.Background(), nil)
			if err != nil {
				logger.Fatalf("failed to connect to database")
			}
			fmt.Println("Successfully connected to MongoDB")
			collection := client.Database("loggingapi").Collection("logs")

			documento := bson.D{
				{Key: "nome", Value: "Exemplo"},
				{Key: "valor", Value: 123},
			}
			resultado, err := collection.InsertOne(context.Background(), documento)
			if err != nil {
				log.Fatalf("Erro ao inserir documento: %v", err)
			}

			fmt.Printf("Documento inserido com ID: %v\n", resultado.InsertedID)

			s, err := server.New(conf, logger)
			if err != nil {
				logger.Fatalf("failed to start server: %v", err)
			}

			if err := s.Run(); err != nil {
				logger.Fatalf("failed to start server: %v", err)
			}*/
}
