package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	pgerr "github.com/TiJon8/todo-tracker/internal/core/infra/postgres"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

func (repo *RepositoryPostgres) PatchUser(ctx context.Context, id string, user domain.User) (domain.User, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		UPDATE todo.users
		SET row_version=row_version+1, name=$3, phone=$4
		WHERE id=$1 AND row_version=$2
		RETURNING id, row_version, name, phone;
	`

	row := repo.Pool.QueryRow(context, query, id, user.Version, user.Name, user.Phone)
	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.Name, &userModel.Phone); err != nil {
		if errors.Is(err, pgerr.ErrNoRows) {
			return domain.User{}, fmt.Errorf("При изменении пользователя произошла ошибка. Версия поля для пользователя %s могла поменяться: %w", id, exceptions.ConflictException)
		}
		return domain.User{}, fmt.Errorf("Ошибка при валидации данных из бд в модель: %w", err)
	}

	domain := domain.NewUser(userModel.ID, userModel.Version, userModel.Name, userModel.Phone)
	return domain, nil
}
