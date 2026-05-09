package core_transport_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"github.com/go-playground/validator/v10"
)

var core = validator.New()

type validateble interface {
	Validate() error
}

func Validate(r *http.Request, obj any) error {
	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return fmt.Errorf("Произошла ошибка при обработке json: %v: %w", err, exceptions.BadRequestException)
	}
	var err error

	v, ok := obj.(validateble)
	if ok {
		err = v.Validate()
	} else {
		err = core.Struct(obj)
	}
	if err != nil {
		return fmt.Errorf("Ошибка при валидации данных: %v: %w", err, exceptions.BadRequestException)
	}

	return nil
}
