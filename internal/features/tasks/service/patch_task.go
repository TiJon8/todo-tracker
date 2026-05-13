package task_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *TaskService) PatchTask(ctx context.Context, authorId string, taskId string, patch domain.TaskPatch) (domain.Task, error) {
	task, err := us.GetTask(ctx, authorId, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}
	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("Aplly patch error: %w", err)
	}

	patched, err := us.Repository.PatchTask(ctx, authorId, taskId, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch error: %w", err)
	}
	if err := us.HistoryService.CreateSnapshot(ctx, "patch", patched); err != nil {
		return domain.Task{}, fmt.Errorf("Ошибка при создании записи истории: %w", err)
	}
	return patched, nil
}
