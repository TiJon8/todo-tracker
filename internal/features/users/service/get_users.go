package users_http_service

import (
	"context"
	"fmt"

	core_domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *UserService) GetUsers(ctx context.Context, limit *int, offset *int) ([]core_domain.User, error) { 
	users, err := us.UserRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получения пользователей из бд: %w", err)
	}
	return users, nil
}