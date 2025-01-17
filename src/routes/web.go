package routes

import (
	"github.com/go-chi/chi/v5"
	"loggingapi/src/pkg/controller"
)

func Web(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/", controller.Index())
	})
}
