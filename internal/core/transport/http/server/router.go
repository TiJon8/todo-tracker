package core_transport_http_server

import (
	"fmt"
	"net/http"

	middleware "github.com/TiJon8/todo-tracker/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	Middleware []middleware.Middleware
}

func NewApiVersionRouter(version ApiVersion, middleware ...middleware.Middleware) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: version,
		Middleware: middleware,
	}
}

func (r *ApiVersionRouter) Register(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *ApiVersionRouter) WithMiddleware() http.Handler {
	return middleware.Chain(r, r.Middleware...)
}
