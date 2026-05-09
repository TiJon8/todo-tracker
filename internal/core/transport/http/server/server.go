package core_transport_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	middleware "github.com/TiJon8/todo-tracker/internal/core/transport/http/middleware"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Addr              string        `envconfig:"ADDR" required:"true"`
	ShutdownDucration time.Duration `envconfig:"SHUTDOWN_DURATION" required:"false" default:"30s"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("HTTP", &cfg); err != nil {
		return Config{}, fmt.Errorf("ошибка при валидации env http параметров: %w", err)
	}
	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	logger     *logger.Logger
	middleware []middleware.Middleware
}

func NewHTTPServer(config Config, logger *logger.Logger, middleware ...middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		logger:     logger,
		middleware: middleware,
	}
}

func (s *HTTPServer) Register(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleware()))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := middleware.Chain(s.mux, s.middleware...)
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1)
	go func() {
		s.logger.Warn("Started Server on", zap.String("address", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Ошибка при старте сервера: %w", err)
		}
	case <-ctx.Done():
		s.logger.Warn("shutdown server")
		context, cancel := context.WithTimeout(context.Background(), s.config.ShutdownDucration)
		defer cancel()
		if err := server.Shutdown(context); err != nil {
			_ = server.Close()
			return fmt.Errorf("Не удалось остановить сервер за отведенное время %v, %w", s.config.ShutdownDucration, err)
		}
		s.logger.Warn("Server has stopped")
	}
	return nil
}
