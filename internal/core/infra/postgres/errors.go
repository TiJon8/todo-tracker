package core_infra_postgres

import "errors"

var (
	ErrNoRows             = errors.New("sql: no rows in result set")
	ErrViolatesForeignKey = errors.New("sql: violates foreign key")
	ErrUnknown = errors.New("sql: error unknown")
)
