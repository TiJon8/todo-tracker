package core_transport_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_transport_http_exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
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
