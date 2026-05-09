package core_infra_postgres

import "errors"


var ErrNoRows = errors.New("sql: no rows in result set")