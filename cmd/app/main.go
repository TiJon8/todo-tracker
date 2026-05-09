package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_infra_postgres_pgx "github.com/TiJon8/todo-tracker/internal/core/infra/postgres/pgx"
	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	middleware "github.com/TiJon8/todo-tracker/internal/core/transport/http/middleware"
	server "github.com/TiJon8/todo-tracker/internal/core/transport/http/server"
	users_repository_postgres "github.com/TiJon8/todo-tracker/internal/features/users/repository"
	users_http_service "github.com/TiJon8/todo-tracker/internal/features/users/service"
	users_trasport_http "github.com/TiJon8/todo-tracker/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := logger.NewLogger(logger.NewLoggerConfigMust())
	if err != nil {
		fmt.Println("Не удалось создать логер для сервера:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Application is bootstarting... ")

	pool, err := core_infra_postgres_pgx.NewPostgresConnPool(ctx, core_infra_postgres_pgx.NewConfigMust())
	if err != nil {
		logger.Fatal("Не получилось создать postgres pool", zap.Error(err))
	}
	defer pool.Close()

	usersRepository := users_repository_postgres.NewRepositoryPostgres(pool)
	usersService := users_http_service.NewUserService(usersRepository)
	UsersTransport := users_trasport_http.NewUserHandlers(usersService)

	ApiVersionRouter := server.NewApiVersionRouter(server.ApiVersion1)
	ApiVersionRouter.Register(UsersTransport.Routes()...)


	/*
		Возможность вешать middleware для определенной версии
	*/
	// ApiVersionRouter2 := server.NewApiVersionRouter(server.ApiVersion2, middleware.Dumb("/api v2 router/"))
	// ApiVersionRouter2.Register(UsersTransport.Routes()...)

	Server := server.NewHTTPServer(
		server.NewConfigMust(),
		logger,
		middleware.RequestID(),
		middleware.ConnectLogger(logger),
		middleware.Trace(),
		middleware.CatchPanic(),
	)
	Server.Register(ApiVersionRouter)

	if err := Server.Run(ctx); err != nil {
		logger.Error("Ошибка при старте сервера", zap.Error(err))
	}
}
