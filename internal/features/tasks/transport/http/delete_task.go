package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type DeleteTaskRequest struct {
	AuthorID string `json:"author_id" validate:"uuid"`
}

type DeleteTaskResponse TaskDTO

func (h *TaskHTTPHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов DeleteTask обработчика")

	var DeleteTask DeleteTaskRequest
	if err := request.Validate(r, &DeleteTask); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	}
	// id, err := utils.GetPathParam(r, "userId")
	// if err != nil {
	// 	responseWriter.ResponseError(err, "Не удалось получить userId")
	// 	return
	// }
	taskId, err := utils.GetPathParam(r, "taskId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить taskId")
		return
	}

	task, err := h.TaskService.DeleteTask(ctx, DeleteTask.AuthorID, taskId)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := DeleteTaskResponse(DTOFromDomain(task))
	responseWriter.ResponseJSON(res, http.StatusOK)
}
