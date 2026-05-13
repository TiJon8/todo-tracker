package tasks_transport_http

import (
	"fmt"
	"net/http"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorID    string  `json:"author_id" validate:"required,uuid4"`
	GroupID     *string `json:"group_id" validate:"omitempty,uuid4"`
}

type CreateTaskResponse TaskDTO

func (h *TaskHTTPHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов CreateTask обработчика")

	var TaskRequestDTO CreateTaskRequest
	if err := request.Validate(r, &TaskRequestDTO); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	}

	domainTask := domain.NewTaskUninitialized(
		TaskRequestDTO.Title,
		TaskRequestDTO.Description,
		TaskRequestDTO.AuthorID,
		TaskRequestDTO.GroupID,
	)
	fmt.Println(domainTask)
	task, err := h.TaskService.CreateTask(ctx, domainTask)
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось создать задачу")
		return
	}

	res := CreateTaskResponse(DTOFromDomain(task))
	responseWriter.ResponseJSON(res, http.StatusCreated)
}

func createDTOFromDomain(dm domain.Task) TaskDTO {
	return TaskDTO{
		ID:          dm.ID,
		Version:     dm.Version,
		Title:       dm.Title,
		Description: dm.Description,
		Completed:   dm.Completed,
		CompletedAt: dm.CompletedAt,
		AuthorID:    dm.AuthorID,
		GroupID:     dm.GroupID,
		CreatedAt:   dm.CreatedAt,
		UpdatedAt:   dm.UpdatedAt,
		DeletedAt:   dm.DeletedAt,
	}
}
