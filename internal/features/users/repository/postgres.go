package users_repository_postgres

import core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"

type RepositoryPostgres struct {
	Pool core_infra_postgres.Pool
}


func NewRepositoryPostgres(pool core_infra_postgres.Pool) *RepositoryPostgres {
	return &RepositoryPostgres{Pool: pool}
}