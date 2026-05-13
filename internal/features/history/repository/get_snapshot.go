package history_repository

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (repo *RepositoryPostgres) GetSnapshot(ctx context.Context, taskId string) (domain.History, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, title, description, status, task_id, previous, updated_by, created_at FROM todo.history
		WHERE task_id=$1
		ORDER BY created_at DESC
		LIMIT 1;
	`
	row := repo.Pool.QueryRow(context, query, taskId)
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
