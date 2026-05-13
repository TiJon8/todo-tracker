package task_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

func (repo *RepositoryPostgres) GetTask(ctx context.Context, authorId string, taskId string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, row_version, title, description, completed, completed_at, user_id, group_id, created_at, updated_at, deleted_at
		FROM todo.tasks
		WHERE user_id=$1 AND id=$2;
	`

	row := repo.Pool.QueryRow(ctx, query, authorId, taskId)
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
			return domain.Task{}, fmt.Errorf("foreign key %s error; %v; %w", authorId, err, exceptions.BadRequestException)
		}
		if errors.Is(err, core_infra_postgres.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("Задача с id=%s не найдена: %v: %w", taskId, err, exceptions.NotFoundException)
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
