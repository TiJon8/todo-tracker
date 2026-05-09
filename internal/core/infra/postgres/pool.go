package core_infra_postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type InfraPostgresPool struct {
	*pgxpool.Pool
	timeout time.Duration
}

type Pool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
	GetTimeout() time.Duration
}

func (p *InfraPostgresPool) GetTimeout() time.Duration {
	return p.timeout
}

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Db       string        `envconfig:"NAME" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("DATABASE", &cfg); err != nil {
		return nil, fmt.Errorf("Не удалось проанализировать переменные окружения: %w", err)
	}
	return &cfg, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func NewPostgresConnPool(ctx context.Context, config *Config) (*InfraPostgresPool, error) {
	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)
	fmt.Println(connection)
	pgxConfig, err := pgxpool.ParseConfig(connection)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при парсинге конфига: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при создании pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Ошибка пула: %w", err)
	}
	return &InfraPostgresPool{
		Pool:    pool,
		timeout: config.Timeout,
	}, nil
}
