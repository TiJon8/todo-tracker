package core_infra_postgres_pgx

import (
	"context"
	"fmt"
	"time"

	core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type Pool struct {
	*pgxpool.Pool
	timeout time.Duration
}

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Db       string        `envconfig:"NAME" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func (p *Pool) GetTimeout() time.Duration {
	return p.timeout
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

func NewPostgresConnPool(ctx context.Context, config *Config) (*Pool, error) {
	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)
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

	return &Pool{
		Pool:    pool,
		timeout: config.Timeout,
	}, nil
}

func (p *Pool) Exec(ctx context.Context, sql string, arguments ...any) (core_infra_postgres.CommandTag, error) {
	tag, err := p.Pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}
	return pgxCommandTag{tag}, nil
}

func (p *Pool) Query(ctx context.Context, sql string, args ...any) (core_infra_postgres.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgxRows{rows}, nil
}

func (p *Pool) QueryRow(ctx context.Context, sql string, args ...any) core_infra_postgres.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)
	return pgxRow{row}
}
