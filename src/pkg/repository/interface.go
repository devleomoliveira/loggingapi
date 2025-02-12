package repository

import "loggingapi/src/pkg/model"

type Repository interface {
	Log() LogRepository
	Close() error
}

type LogRepository interface {
	Get(uuid string) []model.Log
	Store(log model.Log) model.Log
}
