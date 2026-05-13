package task_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *TaskService) DeleteTask(ctx context.Context, id string, taskId string) (domain.Task, error) {
	task, err := us.Repository.GetTask(ctx, id, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}
	if task.DeletedAt != nil {
		return task, fmt.Errorf("Expected error: задача по id=%s уже удалена", taskId)
	}
	deletedTask, err := us.Repository.DeleteTask(ctx, id, taskId, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("Ошибка при удалении задачи из бд: %w", err)
	}
	if err := us.HistoryService.CreateSnapshot(ctx, "delete", deletedTask); err != nil {
		return domain.Task{}, fmt.Errorf("Ошибка при создании записи истории: %w", err)
	}
	return deletedTask, nil
}
