package routes

import (
	"github.com/go-chi/chi/v5"
	"loggingapi/src/pkg/config"
	"loggingapi/src/pkg/controller"
)

func Web(r chi.Router, config config.Config) {
	r.Group(func(r chi.Router) {
		r.Get("/", controller.Index())
	})

	if config.Server.ENV == "development" {
		r.Group(func(r chi.Router) {
		})
	}
}
