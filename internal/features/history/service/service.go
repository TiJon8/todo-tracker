package history_service

import (
	"context"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)


type HistoryRepository interface {
	CreateSnapshot(ctx context.Context, task domain.History) (domain.History, error)
	GetSnapshot(ctx context.Context, taskId string) (domain.History, error)
}

type HistoryService struct {
	Repository HistoryRepository
}

func NewHistoryService(repo HistoryRepository) *HistoryService {
	return &HistoryService{
		Repository: repo,
	}
}
