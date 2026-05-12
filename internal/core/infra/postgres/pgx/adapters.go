package core_infra_postgres_pgx

import (
	"errors"
	"fmt"

	core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRow struct {
	pgx.Row
}

type pgxRows struct {
	pgx.Rows
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_infra_postgres.ErrNoRows
		}
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23503" {
				return fmt.Errorf("%v; %w", err, core_infra_postgres.ErrViolatesForeignKey )
			}
		}
		return fmt.Errorf("%v; %w", err, core_infra_postgres.ErrUnknown)
	}
	return nil
}
