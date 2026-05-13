package statistic_trasnport_http

import (
	"fmt"
	"net/http"
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type GetStatisticResponse struct {
	TotalTasks           int
	CompletedTasks       int
	CompletedRate        *float64
	AverageCompletedTime *string
}

func (s *StatisticHTTPHandlers) GetStatistic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов GetStatistic обработчика")

	userId, from, to, err := getStatisticQueryParams(r)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при передаче query params")
		return
	}

	statistic, err := s.StatisticService.GetStatistic(ctx, userId, from, to)
	if err != nil {
		responseWriter.ResponseError(err, "Ошибка при вызове сервиса")
		return
	}

	res := DTOFromDomain(statistic)
	responseWriter.ResponseJSON(res, http.StatusOK)
}

func DTOFromDomain(d domain.Statistic) GetStatisticResponse {
	var avg *string
	if d.AverageCompletedTime != nil {
		d := d.AverageCompletedTime.String()
		avg = &d
	}
	return GetStatisticResponse{
		TotalTasks: d.TotalTasks,
		CompletedTasks: d.CompletedTasks,
		CompletedRate: d.CompletedRate,
		AverageCompletedTime: avg,
	}
}

func getStatisticQueryParams(r *http.Request) (string, *time.Time, *time.Time, error) {
	userId, err := utils.GetUUIDQueryParam(r, "userId")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить userId: %w", err)
	}
	from, err := utils.GetTimeQueryParam(r, "from")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить дату для поля from: %w", err)
	}
	to, err := utils.GetTimeQueryParam(r, "to")
	if err != nil {
		return "", nil, nil, fmt.Errorf("Не удалось получить дату для поля to: %w", err)
	}
	return *userId, from, to, nil
}
