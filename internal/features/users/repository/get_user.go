package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryPostgres) GetUser(ctx context.Context, id string) (core_domain.User, error) {
	context, cancel := context.WithTimeout(ctx, repo.Pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, row_version, name, phone FROM todo.users
		WHERE id=$1
	`
	row := repo.Pool.QueryRow(context, query, id)
	// return []core_domain.User{}, nil
	var userModel UserModel
	if err := row.Scan(&userModel.ID, &userModel.Version, &userModel.Name, &userModel.Phone); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.User{}, fmt.Errorf("Targter with id=%s not found: %w", id, exceptions.NotFoundException)
		}
		return core_domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return core_domain.NewUser(userModel.ID, userModel.Version, userModel.Name, userModel.Phone), nil
}
