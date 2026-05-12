package core_domain

import (
	"fmt"
	"regexp"

	core_transport_http_exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

var (
	uninitializedID      = ""
	uninitializedVersion = -1
	uninitializedTime = ""
)

type User struct {
	ID      string
	Version int

	Name  string
	Phone *string
}

func NewUser(id string, version int, name string, phone *string) User {
	return User{
		ID:      id,
		Version: version,
		Name:    name,
		Phone:   phone,
	}
}

func NewUserUnitialized(name string, phone *string) User {
	return NewUser(uninitializedID, uninitializedVersion, name, phone)
}

func (u *User) Validate() error {
	nameLen := len([]rune(u.Name))
	if nameLen <= 2 || nameLen > 100 {
		return fmt.Errorf("Имя не должно быть меньше 2 символов и больше 100: Получено (%d) %w", nameLen, core_transport_http_exceptions.BadRequestException)
	}
	if u.Phone != nil {
		phoneLen := len([]rune(*u.Phone))
		if phoneLen < 10 || phoneLen > 16 {
			return fmt.Errorf("Номер телефона не должен быть меньше 10 символов и больше 16: Получено (%d) %w", phoneLen, core_transport_http_exceptions.BadRequestException)
		}
		re := regexp.MustCompile(`^\+[0-9]{10,16}$`)
		if !re.MatchString(*u.Phone) {
			return fmt.Errorf(
				"Номер телефона должен соответсвовать формату и начинаться с +: %w",
				exceptions.BadRequestException)
		}
	}

	return nil
}

type UserPatch struct {
	Name  Nullable[string]
	Phone Nullable[string]
}


func NewPatchUser(name Nullable[string], phone Nullable[string]) UserPatch {
	return UserPatch{
		Name: name,
		Phone: phone,
	}
}

func (p *UserPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf("Не валидная структура Name: %w", exceptions.BadRequestException)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("Не удалось провалидировать структуру: %w", err)
	}
	tmp := *u

	if patch.Name.Set {
		tmp.Name = *patch.Name.Value
	}

	if patch.Phone.Set {
		tmp.Phone = patch.Phone.Value
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("Не валидный user patch: %w", err)
	}
	*u = tmp
	return nil
}
