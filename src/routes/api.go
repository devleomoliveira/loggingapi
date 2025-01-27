package routes

import (
	"github.com/go-chi/chi/v5"
	"loggingapi/src/pkg/common"
	"loggingapi/src/pkg/config"
	"net/http"
)

func Api(r chi.Router, c config.Config) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
			})

			if c.Server.ENV == "development" {
				r.Group(func(r chi.Router) {
				})
			}
		})
	})

}
