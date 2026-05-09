package core_transport_http_middleware

import (
	"context"
	"net/http"
	"time"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get("X-Request-ID")
			if requestId == "" {
				requestId = uuid.NewString()
			}
			r.Header.Set("X-Request-ID", requestId)
			w.Header().Set("X-Request-ID", requestId)
			next.ServeHTTP(w, r)
		})
	}
}

func ConnectLogger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get("X-Request-ID")

			l := log.With(
				zap.String("request-id", requestId),
				zap.String("url", r.URL.String()),
			)
			ctx := context.WithValue(r.Context(), logger.LoggerContextKey, l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CatchPanic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := logger.LoggerFromContext(ctx)
			hrh := response.NewHTTPResponseHandler(logger, w)

			defer func() {
				if p := recover(); p != nil {
					hrh.PanicResponse(p, "При обработке запроса возникла паника!")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			rw := response.NewResponseWriter(w)

			logger := logger.LoggerFromContext(ctx)
			now := time.Now()

			logger.Debug(">> Incoming HTTP request", zap.String("method", r.Method), zap.Time("time", now.UTC()))
			next.ServeHTTP(rw, r)
			logger.Debug("<< Outcoming HTTP response", zap.Int("statusCode", rw.GetStatusCode()), zap.Duration("latency", time.Since(now)))
		})
	}
}
