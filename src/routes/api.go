package routes

import (
	"github.com/go-chi/chi/v5"
	"loggingapi/src/pkg/common"
	"net/http"
)

func Api(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					common.ResponseSuccess(w, make(map[string]interface{}))
				})
			})
		})
	})

}
