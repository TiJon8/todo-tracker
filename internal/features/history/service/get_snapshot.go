package history_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (s *HistoryService) GetSnapshot(ctx context.Context, taskId string) (domain.History, error) {
	history, err := s.Repository.GetSnapshot(ctx, taskId)
	if err != nil {
		return domain.History{}, fmt.Errorf("history snapshot get: %w", err)
	}
	return history, nil
}
