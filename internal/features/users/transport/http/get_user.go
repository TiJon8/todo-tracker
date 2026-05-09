package users_http_transport

import (
	"net/http"

	core_logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	core_transport_http_response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LoggerFromContext(ctx)
	responseWriter := core_transport_http_response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов GetUser обработчика")
	id, err := utils.GetPathParam(r, "userId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить userId")
		return
	}

	user, err := h.UserService.GetUser(ctx, id)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := DTOFromDomain(user)
	responseWriter.ResponseJSON(res, http.StatusOK)
}
