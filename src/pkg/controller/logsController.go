package controller

import (
	"net/http"
)

var (
	LogsController = &logsController{}
)

type logsController struct {
	ApiBaseController
}

func (l *logsController) GetLogs(w http.ResponseWriter, _ *http.Request) {
	data := make(map[string]interface{})
	data["teste"] = "test2"
	l.ResponseSuccess(w, data)
}
