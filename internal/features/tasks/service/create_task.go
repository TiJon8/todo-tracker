package task_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (s *TaskService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("Не удалось провалидировать структуру: %w", err)
	}

	task, err := s.Repository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("task create: %w", err)
	}
	s.HistoryService.CreateSnapshot(ctx, "create", task)

	return task, nil
}
