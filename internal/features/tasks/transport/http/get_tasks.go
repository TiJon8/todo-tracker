package tasks_transport_http

import (
	"fmt"
	"net/http"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type GetTasksResponse []TaskDTO

func (h *TaskHTTPHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов GetTasks обработчика")

	id, limit, offset, err := getQueryParams(r)
	if err != nil {
		responseWriter.ResponseError(err, "Не получилось проанализировать query params")
		return
	}
	// id, err := utils.GetPathParam(r, "userId")
	// if err != nil {
	// 	responseWriter.ResponseError(err, "Не удалось получить userId")
	// 	return
	// }

	tasks, err := h.TaskService.GetTasks(ctx, id, limit, offset)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := GetTasksResponse(getTasksDTOList(tasks))
	responseWriter.ResponseJSON(res, http.StatusOK)
}

func getQueryParams(r *http.Request) (string, *int, *int, error) {
	userId, err := utils.GetUUIDQueryParam(r, "userId")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить userId из query params: %w", err)
	}
	limit, err := utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить limit из query params: %w", err)
	}
	offset, err := utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить offset из query params: %w", err)
	}
	return *userId, limit, offset, nil
}
