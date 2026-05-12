package task_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *TaskService) GetTask(ctx context.Context, authorId string, taskId string) (domain.Task, error) {
	task, err := us.Repository.GetTask(ctx, authorId, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf("Ошибка при получения задач из бд: %w", err)
	}
	return task, nil
}
