package users_repository_postgres

import domain "github.com/TiJon8/todo-tracker/internal/core/domain"

type UserModel struct {
	ID      string
	Version int
	Name    string
	Phone   *string
}


func DomainFromModel(model []UserModel) ([]domain.User) {
	slice := make([]domain.User, len(model))

	for i, v := range model {
		slice[i] = domain.NewUser(v.ID, v.Version, v.Name, v.Phone)
	}
	return slice
}