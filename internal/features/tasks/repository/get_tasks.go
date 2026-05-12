package task_repository_postgres

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (repo *RepositoryPostgres) GetTasks(ctx context.Context, id string, limit *int, offset *int) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, row_version, title, description, completed, completed_at, user_id, group_id, created_at, updated_at, deleted_at
		FROM todo.tasks
		WHERE user_id = $1
		LIMIT $2
		OFFSET $3;
	`

	rows, err := repo.Pool.Query(ctx, query, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при выполнении запроса в бд: %w", err)
	}
	defer rows.Close()

	var tasksModelSlice []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(
			&taskModel.ID, &taskModel.Version, &taskModel.Title, &taskModel.Description,
			&taskModel.Completed, &taskModel.CompletedAt, &taskModel.AuthorID, &taskModel.GroupID,
			&taskModel.CreatedAt, &taskModel.UpdatedAt, &taskModel.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("Ошибка при валидации модели: %w", err)
		}

		tasksModelSlice = append(tasksModelSlice, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows next: %w", err)
	}

	return DomainFromModelSlice(tasksModelSlice), nil
}
