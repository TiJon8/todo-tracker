package tasks_transport_http

import (
	"context"
	"net/http"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	server "github.com/TiJon8/todo-tracker/internal/core/transport/http/server"
)

type TaskDTO struct {
	ID          string     `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at"`
	AuthorID    string     `json:"author_id"`
	GroupID     *string    `json:"group_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func DTOFromDomain(task domain.Task) TaskDTO {
	return TaskDTO{
		ID:          task.ID,
		Version:     task.Version,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CompletedAt: task.CompletedAt,
		AuthorID:    task.AuthorID,
		GroupID:     task.GroupID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		DeletedAt:   task.DeletedAt,
	}
}

func getTasksDTOList(tasks []domain.Task) []TaskDTO {
	res := make([]TaskDTO, len(tasks))

	for i, v := range tasks {
		res[i] = DTOFromDomain(v)
	}
	return res
}

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userId string, limit *int, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, authorId string, taskId string) (domain.Task, error)
	DeleteTask(ctx context.Context, id string, taskId string) (domain.Task, error)
	PatchTask(ctx context.Context, authorId string, taskId string, patch domain.TaskPatch) (domain.Task, error)
}

type TaskHTTPHandlers struct {
	TaskService TaskService
}

func NewTaskHTTPHandlers(service TaskService) *TaskHTTPHandlers {
	return &TaskHTTPHandlers{
		TaskService: service,
	}
}

func (h *TaskHTTPHandlers) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{taskId}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{taskId}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{taskId}",
			Handler: h.PatchTask,
		},
	}
}
