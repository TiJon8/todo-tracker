package users_repository_postgres

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (repo *RepositoryPostgres) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, row_version, name, phone FROM todo.users
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
	`
	rows, err := repo.Pool.Query(context, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при выполнении запроса в бд: %w", err)
	}
	defer rows.Close()

	var userModelSlice []UserModel
	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(&userModel.ID, &userModel.Version, &userModel.Name, &userModel.Phone); err != nil {
			return nil, fmt.Errorf("Ошибка при валидации данных из бд в модель: %w", err)
		}
		userModelSlice = append(userModelSlice, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows next: %w", err)
	}

	userDomain := DomainFromModel(userModelSlice)
	return userDomain, nil
}
