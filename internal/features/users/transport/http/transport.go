package users_http_transport

import (
	"context"
	"net/http"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	server "github.com/TiJon8/todo-tracker/internal/core/transport/http/server"
)

type UserDTO struct {
	ID      string  `json:"id"`
	Version int     `json:"version"`
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
}

func DTOFromDomain(user domain.User) UserDTO {
	return UserDTO{
		ID:      user.ID,
		Version: user.Version,
		Name:    user.Name,
		Phone:   user.Phone,
	}
}

func getUsersDTOList(users []domain.User) []UserDTO {
	res := make([]UserDTO, len(users))

	for i, v := range users {
		res[i] = DTOFromDomain(v)
	}
	return res
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id string) (users domain.User, err error)
	GetUsers(ctx context.Context, limit *int, offset *int) (users []domain.User, err error)
	DeleteUser(ctx context.Context, id string) (users domain.User, err error)
	PatchUser(ctx context.Context, id string, patch domain.UserPatch) (users domain.User, err error)
}

type UserHandler struct {
	UserService UserService
}

func NewUserHandlers(service UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:     http.MethodGet,
			Path:       "/users",
			Handler:    h.GetUsers,
			/*
			Возможность вешать свои middleware для определенных обработчиков
			*/
			// Middleware: []middleware.Middleware{middleware.Dumb("/get users/")},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{userId}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{userId}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{userId}",
			Handler: h.PatchUser,
		},
	}
}
