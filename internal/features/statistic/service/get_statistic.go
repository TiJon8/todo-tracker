package statistic_http_service

import (
	"context"
	"fmt"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

func (s *StatisticService) GetStatistic(ctx context.Context, userId string, from *time.Time, to *time.Time) (domain.Statistic, error) {
	if from != nil && to != nil {
		if to.Before(*from) {
			return domain.Statistic{}, fmt.Errorf("Поле to не должно быть меньше значения from: %w", exceptions.BadRequestException)
		}
	}
	tasks, err := s.Repository.GetTasks(ctx, userId, from, to)
	if err != nil {
		return domain.Statistic{}, fmt.Errorf("Ошибка при получении задач для статистики из бд: %w", err)
	}
	return calcStats(tasks), nil
}

func calcStats(tasks []domain.Task) domain.Statistic {
	if len(tasks) == 0 {
		return domain.NewStatistic(0, 0, nil, nil)
	}
	total := len(tasks)
	completed := 0
	var completionDuration time.Duration
	for _, v := range tasks {
		if v.Completed {
			completed++
		}
		duration := v.GetDuration()
		if duration != nil {
			completionDuration += *duration
		}
	}

	completedRate := float64(completed) / float64(total) * 100

	var averageCompletionTime *time.Duration
	if completed > 0 && completionDuration != 0 {
		averageCompleted := completionDuration / time.Duration(completed)
		averageCompletionTime = &averageCompleted
	}

	return domain.NewStatistic(total, completed, &completedRate, averageCompletionTime)
}
