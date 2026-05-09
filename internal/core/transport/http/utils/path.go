package core_transport_http_utils

import (
	"fmt"
	"net/http"

	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"github.com/google/uuid"
)


func GetPathParam(r *http.Request, param string) (string, error) {
	val := r.PathValue(param)
	if val == "" {
		return "", fmt.Errorf("Не удалось получить значение из пармаетра пути %s: %w", param, exceptions.BadRequestException)
	}
	if err := uuid.Validate(val); err != nil {
		return "", fmt.Errorf("Формат значения параметра пути %s не соответствует формату UUID: %v; %w", param, err, exceptions.BadRequestException)
	}
	return val, nil
}