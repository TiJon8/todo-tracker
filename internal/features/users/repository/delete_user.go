package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"github.com/jackc/pgx/v5"
)


func (repo *RepositoryPostgres) DeleteUser(ctx context.Context, id string) (domain.User, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		DELETE FROM todo.users
		WHERE id=$1
		RETURNING id, row_version, name, phone
	`
	row := repo.Pool.QueryRow(context, query, id)
	// return []core_domain.User{}, nil
	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.Name, &userModel.Phone); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("Targter with id=%s not found: %w", id, exceptions.NotFoundException)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return domain.NewUser(userModel.ID, userModel.Version, userModel.Name, userModel.Phone), nil
}
