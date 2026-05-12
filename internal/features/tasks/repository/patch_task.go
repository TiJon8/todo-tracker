package task_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	core_infra_postgres "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

func (repo *RepositoryPostgres) PatchTask(ctx context.Context, authorId string, taskId string, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		UPDATE todo.tasks
		SET row_version=row_version+1, title=$4, description=$5, completed=$6, completed_at=$7
		WHERE user_id=$1 AND id=$2 AND row_version=$3
		RETURNING id, row_version, title, description, completed, completed_at, user_id, group_id, created_at, updated_at, deleted_at;
	`

	row := repo.Pool.QueryRow(ctx, query, authorId, taskId, task.Version, task.Title, task.Description, task.Completed, task.CompletedAt)

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

		if errors.Is(err, core_infra_postgres.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("При обновлении задачи возникла ошибка: %w", exceptions.ConflictException)
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
