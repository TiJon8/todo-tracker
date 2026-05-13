package core_transport_http_utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_transport_http_exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"github.com/google/uuid"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"Error: param=%s; key=%s; not a valid integer; expected %v; %w",
			param,
			key,
			err,
			core_transport_http_exceptions.BadRequestException,
		)
	}
	return &val, nil
}

func GetUUIDQueryParam(r *http.Request, key string) (*string, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, fmt.Errorf("Парметр переданный в поле, где ожидается uuid не может быть пустым!")
	}
	if err := uuid.Validate(param); err != nil {
		return nil, fmt.Errorf(
			"Error: param=%s; key=%s; not a valid uuid; expected %v; %w",
			param,
			key,
			err,
			core_transport_http_exceptions.BadRequestException,
		)
	}
	return &param, nil
}

func GetTimeQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, param)
	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга даты: %w", err)
	}
	return &parsedTime, nil
}
