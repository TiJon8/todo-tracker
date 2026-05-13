package history_repository

import (
	"time"
)

type HistoryModel struct {
	ID          string
	Title       string
	Description *string
	TaskID      string
	Status      string //core_domain.EnumTaskStatus
	UpdatedBy   string
	Previous    *string
	CreatedAt   time.Time
}
