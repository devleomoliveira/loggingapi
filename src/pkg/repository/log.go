package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"loggingapi/src/pkg/model"
)

type logRepository struct {
	db *mongo.Collection
}

func NewLogRepository(db *mongo.Database) LogRepository {
	log := model.Log{}
	collection := db.Collection(log.CollectionName())
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "uuid", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	collection.Indexes().CreateOne(context.TODO(), indexModel)

	return &logRepository{
		db: collection,
	}
}

func (l logRepository) Get(uuid string) []model.Log {
	filter := bson.M{}
	var logs []model.Log
	if uuid != "" {
		filter = bson.M{"uuid": uuid}
	}

	c, error := l.db.Find(context.Background(), filter)

	if error != nil {
		fmt.Println(error)
	}

	if err := c.All(context.Background(), &logs); err != nil {
		fmt.Println(err)
	}

	for _, logEntry := range logs {
		fmt.Printf("UUID: %s, Name: %s, Category: %s, Level: %s, Message: %s\n",
			logEntry.UUID, logEntry.Name, logEntry.Category, logEntry.Level, logEntry.Message)
	}
	return logs

}

func (l logRepository) Store(log model.Log) model.Log {
	l.db.InsertOne(context.Background(), log)

	return log
}

func (l logRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (l logRepository) Update() (model.Log, error) {
	//TODO implement me
	panic("implement me")
}
