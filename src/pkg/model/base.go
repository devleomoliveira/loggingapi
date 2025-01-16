package model

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	collection mongo.Collection
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"-"`
}
