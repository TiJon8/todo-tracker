package users_http_transport

import (
	"fmt"
	"net/http"

	core_logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	core_transport_http_response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	core_transport_http_utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTO

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LoggerFromContext(ctx)
	responseWriter := core_transport_http_response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов GetUsers обработчика")

	limit, offset, err := getQueryParams(r)
	if err != nil {
		responseWriter.ResponseError(err, "Не получилось проанализировать query params")
		return
	}

	users, err := h.UserService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseWriter.ResponseError(err, "Не получилось получить пользователей")
		return
	}

	response := GetUsersResponse(getUsersDTOList(users))
	responseWriter.ResponseJSON(response, http.StatusOK)
}

func getQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_transport_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("Не удалось получить limit из query params: %w", err)
	}
	offset, err := core_transport_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("Не удалось получить offset из query params: %w", err)
	}
	return limit, offset, nil
}
