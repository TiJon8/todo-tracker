package users_http_transport

import (
	"net/http"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	core_logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	core_transport_http_request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	core_transport_http_response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Phone *string `json:"phone" validate:"omitempty,startswith=+,min=10,max=16"`
}

type CreateUserResponse UserDTO


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LoggerFromContext(ctx)
	responseWriter := core_transport_http_response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов CreateUser обработчика")
	var UserData CreateUserRequest
	if err := core_transport_http_request.Validate(r, &UserData); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	} 
	
	userDomain := domainFromDTO(UserData)
	user, err := h.UserService.CreateUser(ctx, userDomain)
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось создать пользователя")
		return
	}
	res := CreateUserResponse(DTOFromDomain(user))
	responseWriter.ResponseJSON(res, http.StatusCreated)
}

func domainFromDTO(request CreateUserRequest) domain.User {
	return domain.NewUserUnitialized(request.Name, request.Phone)
}
