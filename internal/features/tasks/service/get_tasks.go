package task_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *TaskService) GetTasks(ctx context.Context, userId string, limit *int, offset *int) ([]domain.Task, error) {
	tasks, err := us.Repository.GetTasks(ctx, userId, limit, offset)
	fmt.Println(tasks)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получения задач из бд: %w", err)
	}
	return tasks, nil
}
