package core_transport_http_server

import (
	"fmt"
	"net/http"
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
}


func NewApiVersionRouter(version ApiVersion) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux: http.NewServeMux(),
		apiVersion: version,
	}
}


func (r *ApiVersionRouter) Register(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.Handler)
	}
}