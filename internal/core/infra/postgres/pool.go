package core_infra_postgres

import (
	"context"
	"time"
)

type Pool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Close()
	GetTimeout() time.Duration
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
	Values() ([]any, error)
	RawValues() [][]byte
}

type Row interface {
	Scan(dest ...any) error
}

type CommandTag interface{
	RowsAffected() int64
}
