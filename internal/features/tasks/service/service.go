package task_http_service

import (
	"context"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	history_service "github.com/TiJon8/todo-tracker/internal/features/history/service"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, id string, limit *int, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, authorId string, taskId string) (domain.Task, error)
	DeleteTask(ctx context.Context, userId string, taskId string, task domain.Task) (domain.Task, error)
	PatchTask(ctx context.Context, userId string, taskId string, task domain.Task) (domain.Task, error)
}

type TaskService struct {
	Repository     TaskRepository
	HistoryService *history_service.HistoryService
}

func NewTaskService(repo TaskRepository, historyService *history_service.HistoryService) *TaskService {
	return &TaskService{
		Repository:     repo,
		HistoryService: historyService,
	}
}
