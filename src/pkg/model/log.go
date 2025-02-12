package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	CollectionName = "logs"
)

type Log struct {
	UUID      string             `bson:"uuid"`
	Name      string             `bson:"name"`
	Category  string             `bson:"category"`
	Level     string             `bson:"level"`
	Message   string             `bson:"message"`
	Tags      []string           `bson:"tags"`
	Trace     []string           `bson:"trace"`
	Date      primitive.DateTime `bson:"date"`
	BaseModel BaseModel          `bson:",inline"`
}

func (l *Log) CollectionName() string {
	return CollectionName
}
