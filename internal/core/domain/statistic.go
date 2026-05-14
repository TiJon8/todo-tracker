package core_domain

import "time"

type Statistic struct {
	TotalTasks           int
	CompletedTasks       int
	CompletedRate        *float64
	AverageCompletedTime *time.Duration
}

func NewStatistic(total int, completed int, completedRate *float64, averageCompletedTime *time.Duration) Statistic {
	return Statistic{
		TotalTasks:           total,
		CompletedTasks:       completed,
		CompletedRate:        completedRate,
		AverageCompletedTime: averageCompletedTime,
	}
}
