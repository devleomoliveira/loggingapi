package middleware

import (
	"loggingapi/src/pkg/common"
	"net/http"
	"time"

	"github.com/go-chi/httprate"
	"loggingapi/src/pkg/utils/ratelimit"
)

func RateLimitMiddleware(configs []ratelimit.LimitConfig) (func(http.Handler) http.Handler, error) {
	var limiters []func(http.Handler) http.Handler
	for _, config := range configs {
		limiter := httprate.Limit(
			config.QPS,
			time.Second,
			httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
				common.ResponseFailed(w, r, http.StatusTooManyRequests, nil)
			}),
		)
		limiters = append(limiters, limiter)
	}

	return func(next http.Handler) http.Handler {
		for _, limiter := range limiters {
			next = limiter(next)
		}
		return next
	}, nil
}
