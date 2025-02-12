package service

import (
	"context"
	"loggingapi/src/pkg/repository"
)

type LogService struct {
	logRepository repository.LogRepository
}

func (l LogService) GetLogList(ctx context.Context, pageNum, pageSize int) (res []map[string]interface{}, err error) {
	return make([]map[string]interface{}, 0), nil
}
