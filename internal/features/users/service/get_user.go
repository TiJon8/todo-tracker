package users_http_service

import (
	"context"
	"fmt"

	core_domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *UserService) GetUser(ctx context.Context, id string) (core_domain.User, error) { 
	user, err := us.UserRepository.GetUser(ctx, id)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("Ошибка при получения пользователя из бд: %w", err)
	}
	return user, nil
}