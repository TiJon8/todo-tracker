package task_repository_postgres

import infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"

type RepositoryPostgres struct {
	Pool infra_postgres.Pool
}

func NewTaskRepository(pool infra_postgres.Pool) *RepositoryPostgres {
	return &RepositoryPostgres{
		Pool: pool,
	}
}