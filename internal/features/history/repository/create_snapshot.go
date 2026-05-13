package history_repository

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (repo *RepositoryPostgres) CreateSnapshot(ctx context.Context, task domain.History) (domain.History, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		INSERT INTO todo.history (task_id, title, description, status, updated_by, previous)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, status, task_id, previous, updated_by, created_at;
	`
	row := repo.Pool.QueryRow(context, query, task.TaskID, task.Title, task.Description, task.Status, task.UpdatedBy, task.Previous)
	var historyModel HistoryModel
	if err := row.Scan(
		&historyModel.ID,
		&historyModel.Title,
		&historyModel.Description,
		&historyModel.Status,
		&historyModel.TaskID,
		&historyModel.Previous,
		&historyModel.UpdatedBy,
		&historyModel.CreatedAt,
	); err != nil {
		return domain.History{}, fmt.Errorf("scan history error: %w", err)
	}
	// fmt.Println(historyModel)

	return domain.NewHistory(
		historyModel.ID,
		historyModel.Title,
		historyModel.Description,
		historyModel.TaskID,
		domain.ToEnum(historyModel.Status),
		historyModel.UpdatedBy,
		historyModel.Previous,
		historyModel.CreatedAt,
	), nil
}
