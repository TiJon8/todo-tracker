package users_repository_postgres

import (
	"context"
	"fmt"

	core_domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (repo *RepositoryPostgres) SaveUser(ctx context.Context, user core_domain.User) (core_domain.User, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())

	defer cancel()

	query := `
		INSERT INTO todo.users (name, phone)
		VALUES ($1, $2)
		RETURNING id, row_version, name, phone;
	`
	row := repo.Pool.QueryRow(context, query, user.Name, user.Phone)

	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.Name, &userModel.Phone); err != nil {
		return core_domain.User{}, fmt.Errorf("Ошибка при валидации данных из бд в модель: %w", err)
	}

	return core_domain.NewUser(userModel.ID, userModel.Version, userModel.Name, userModel.Phone), nil

}
