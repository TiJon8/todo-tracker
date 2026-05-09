package core_transport_http_exceptions

import (
	"errors"
)


var (
	NotFoundException = errors.New("Not Found") // http.StatusNotFound
	BadRequestException = errors.New("Bad Request") // http.StatusBadRequest
	ForbiddenException = errors.New("Forbidden") // http.StatusForbidden
	ConflictException = errors.New("Conflict") // http.StatusConflict
)