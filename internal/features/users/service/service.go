package users_http_service

import (
	"context"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)


type UserRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id string) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	DeleteUser(ctx context.Context, id string) (domain.User, error)
	PatchUser(ctx context.Context, id string, user domain.User) (domain.User, error)
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}
