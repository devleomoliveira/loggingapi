package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db  *mongo.Database
	log LogRepository
}

func New(db *mongo.Database) Repository {
	r := &repository{
		db:  db,
		log: NewLogRepository(db),
	}

	return r
}

func (r *repository) Close() error {
	db := r.db

	if db != nil {
		if err := db.Client().Disconnect(context.Background()); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) Log() LogRepository {
	return r.log
}
