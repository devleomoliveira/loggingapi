package database

import (
	"context"

	"loggingapi/src/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(conf *config.DBConfig) (*mongo.Database, error) {
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

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(conf.Name), nil
}
