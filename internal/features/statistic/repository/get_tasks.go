package statistic_repository_postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	task_repository_postgres "github.com/TiJon8/todo-tracker/internal/features/tasks/repository"
)

func (repo *RepositoryPostgres) GetTasks(ctx context.Context, userId string, from *time.Time, to *time.Time) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()
	var strngBuilder strings.Builder

	strngBuilder.WriteString(`
		SELECT id, row_version, title, description, completed, completed_at, user_id, group_id, created_at, updated_at, deleted_at
		FROM todo.tasks
	`)

	conditions := []string{}
	args := []any{}

	conditions = append(conditions, fmt.Sprintf("user_id=$%d", len(args)+1))
	args = append(args, userId)

	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}
	if to != nil {
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)+1))
		args = append(args, to)
	}

	if len(conditions) > 0 {
		strngBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}
	strngBuilder.WriteString(" ORDER BY created_at DESC ")

	rows, err := repo.Pool.Query(ctx, strngBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("query rows err: %w", err)
	}
	defer rows.Close()

	var tasksModelSlice []task_repository_postgres.TaskModel
	for rows.Next() {
		var taskModel task_repository_postgres.TaskModel
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

	return task_repository_postgres.DomainFromModelSlice(tasksModelSlice), nil

}
