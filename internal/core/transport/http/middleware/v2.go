package core_transport_http_middleware

import (
	"fmt"
	"net/http"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
)

func Dumb(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := logger.LoggerFromContext(ctx)

			logger.Debug(fmt.Sprintf("-> before %s", s))
			next.ServeHTTP(w, r)
			logger.Debug(fmt.Sprintf("<- after %s", s))
		})
	}
}
