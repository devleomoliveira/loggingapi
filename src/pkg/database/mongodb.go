package database

import (
	"context"

	"loggingapi/src/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(conf *config.DBConfig) (*mongo.Client, error) {
	dbURI := "mongodb://" + conf.Host
	clientOptions := options.Client().ApplyURI(dbURI)
	if conf.User != "" && conf.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: conf.User,
			Password: conf.Password,
		})
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
