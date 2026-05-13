package tasks_transport_http

import (
	"net/http"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type GetTaskRequest struct {
	AuthorID string `json:"author_id" validate:"uuid"`
}

type GetTaskResponse TaskDTO

func (h *TaskHTTPHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов GetTask обработчика")

	var GetTask GetTaskRequest
	if err := request.Validate(r, &GetTask); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	}

	taskId, err := utils.GetPathParam(r, "taskId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить userId")
		return
	}

	task, err := h.TaskService.GetTask(ctx, GetTask.AuthorID, taskId)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := GetTaskResponse(DTOFromDomain(task))
	responseWriter.ResponseJSON(res, http.StatusOK)
}
