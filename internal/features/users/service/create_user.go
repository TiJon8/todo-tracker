package users_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("При валидации пользователя произошла ошибка: %w", err)
	}

	user, err := us.UserRepository.SaveUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("Ошибка при сохранении пользователя в бд: %w", err)
	}
	return user, nil
}
