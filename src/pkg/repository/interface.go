package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository interface {
	Log() LogRepository
	Init() error
	Close() error
}

type LogRepository interface {
	uuid() string
	name() string
	category() string
	level() string
	message() string
	tags() *[]string
	trace() *[]string
	date() primitive.DateTime
}
