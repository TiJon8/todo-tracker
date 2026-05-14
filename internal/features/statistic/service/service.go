package statistic_http_service

import (
	"context"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

type StatisticRepository interface {
	GetTasks(ctx context.Context, userId string, from *time.Time, to *time.Time) ([]domain.Task, error)
}

type StatisticService struct {
	Repository StatisticRepository
}

func NewStatisticService(repo StatisticRepository) *StatisticService {
	return &StatisticService{
		Repository: repo,
	}
}
