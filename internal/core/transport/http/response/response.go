package core_transport_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	core_transport_http_exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	logger *logger.Logger
	rw     http.ResponseWriter
}

func NewHTTPResponseHandler(logger *logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		logger: logger,
		rw:     rw,
	}
}

func (rp *HTTPResponseHandler) ResponseJSON(data any, statusCode int) {
	rp.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rp.rw).Encode(data); err != nil {
		rp.logger.Error("Не удалось закодировать json", zap.Error(err))
	}
}

func (er *HTTPResponseHandler) ResponseError(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_transport_http_exceptions.NotFoundException):
		statusCode = http.StatusNotFound
		logFunc = er.logger.Debug
	case errors.Is(err, core_transport_http_exceptions.ConflictException):
		statusCode = http.StatusConflict
		logFunc = er.logger.Warn
	case errors.Is(err, core_transport_http_exceptions.ForbiddenException):
		statusCode = http.StatusForbidden
		logFunc = er.logger.Warn
	case errors.Is(err, core_transport_http_exceptions.BadRequestException):
		statusCode = http.StatusBadRequest
		logFunc = er.logger.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = er.logger.Error
	}
	logFunc(msg, zap.Error(err))
	er.errorResponse(msg, statusCode, err)
}

func (rp *HTTPResponseHandler) PanicResponse(p any, msg string) {
	status := http.StatusInternalServerError
	err := fmt.Errorf("Unexpected panic: %v", p)

	rp.logger.Error(msg, zap.Error(err))
	rp.errorResponse(msg, status, err)
}

func (rp *HTTPResponseHandler) errorResponse(msg string, statusCode int, err error) {
	m := map[string]string{
		"message": msg,
		"details": err.Error(),
	}
	rp.ResponseJSON(m, statusCode)
}
