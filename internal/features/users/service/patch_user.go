package users_http_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (us *UserService) PatchUser(ctx context.Context, id string, patch domain.UserPatch) (domain.User, error) {
	user, err := us.UserRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("Aplly patch error: %w", err)
	}
	patched, err := us.UserRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch error: %w", err)
	}
	return patched, nil
}
