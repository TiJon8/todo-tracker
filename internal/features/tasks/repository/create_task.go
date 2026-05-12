package task_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

func (repo *RepositoryPostgres) CreateTask(ctx context.Context, user domain.Task) (domain.Task, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		INSERT INTO todo.tasks (title, description, user_id, group_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, row_version, title, description, completed, completed_at, user_id, group_id, created_at, updated_at, deleted_at;
	`
	row := repo.Pool.QueryRow(context, query, user.Title, user.Description, user.AuthorID, user.GroupID)

	var taskModel TaskModel
	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CompletedAt,
		&taskModel.AuthorID,
		&taskModel.GroupID,
		&taskModel.CreatedAt,
		&taskModel.UpdatedAt,
		&taskModel.DeletedAt,
	); err != nil {
		if errors.Is(err, core_infra_postgres.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf("foreign key %s error; %v; %w", user.AuthorID, err, exceptions.NotFoundException)
		}
		return domain.Task{}, fmt.Errorf("scan error; %w", err)
	}

	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CompletedAt,
		taskModel.AuthorID,
		taskModel.GroupID,
		taskModel.CreatedAt,
		taskModel.UpdatedAt,
		taskModel.DeletedAt,
	), nil

}
