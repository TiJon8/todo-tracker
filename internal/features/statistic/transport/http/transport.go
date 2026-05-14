package statistic_trasnport_http

import (
	"context"
	"net/http"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	server "github.com/TiJon8/todo-tracker/internal/core/transport/http/server"
)

type StatisticHTTPHandlers struct {
	StatisticService StatisticService
}

type StatisticService interface {
	GetStatistic(ctx context.Context, userId string, from *time.Time, to *time.Time) (domain.Statistic, error)
}

func NewStatisticHTTPHandlers(service StatisticService) *StatisticHTTPHandlers {
	return &StatisticHTTPHandlers{
		StatisticService: service,
	}
}

func (h *StatisticHTTPHandlers) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistic",
			Handler: h.GetStatistic,
		},
	}
}
