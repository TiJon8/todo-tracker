package users_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *UserService) DeleteUser(ctx context.Context, id string) (domain.User, error) {
	user, err := us.UserRepository.DeleteUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("Ошибка при удалении пользователя из бд: %w", err)
	}
	return user, nil
}
