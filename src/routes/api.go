package routes

import (
	"github.com/go-chi/chi/v5"
	"loggingapi/src/pkg/common"
	"loggingapi/src/pkg/config"
	"loggingapi/src/pkg/controller"
	"net/http"
)

const currentVersion = "1.0"

func Api(r chi.Router, config config.Config) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v"+currentVersion, func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Get("/logs", controller.LogsController.GetLogs)
				r.Get("/logs/{id}", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
				r.Post("/logs", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
				r.Put("/logs/{id}", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
				r.Patch("/logs/{id}", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
			})

			if config.Server.ENV == "development" {
				r.Group(func(r chi.Router) {
				})
			}
		})
	})

}
