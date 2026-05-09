package users_http_transport

import (
	"net/http"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов DeleteUser обработчика")

	id, err := utils.GetPathParam(r, "userId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить userId")
		return
	}

	user, err := h.UserService.DeleteUser(ctx, id)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := DTOFromDomain(user)
	responseWriter.ResponseJSON(res, http.StatusOK)
}
