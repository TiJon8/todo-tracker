package core_domain

import (
	"time"
)

type EnumTaskStatus int

const (
	InProgress EnumTaskStatus = iota
	Patched
	Done
	Deleted
)

func (s EnumTaskStatus) String() string {
	switch s {
	case InProgress:
		return "in_progress"
	case Patched:
		return "patched"
	case Done:
		return "done"
	case Deleted:
		return "deleted"
	default:
		return "in_progress"
	}
}

func ToEnum(s string) EnumTaskStatus {
	switch s {
	case "in_progress":
		return InProgress
	case "patched":
		return Patched
	case "done":
		return Done
	case "deleted":
		return Deleted
	default:
		return InProgress
	}
}

type History struct {
	ID          string
	Title       string
	Description *string
	TaskID      string
	Status      string
	UpdatedBy   string
	Previous    *string
	CreatedAt   time.Time
}

func NewHistory(
	id string,
	title string,
	description *string,
	taskId string,
	status EnumTaskStatus,
	updatedBy string,
	previous *string,
	createdAt time.Time,
) History {
	return History{
		ID:          id,
		Title:       title,
		Description: description,
		TaskID:      taskId,
		Status:      status.String(),
		UpdatedBy:   updatedBy,
		Previous:    previous,
		CreatedAt:   createdAt,
	}
}

func NewHistoryUninitialized(title string, description *string, taskId string, status EnumTaskStatus, updatedBy string, previous *string) History {
	return NewHistory(
		uninitializedID,
		title,
		description,
		taskId,
		status,
		updatedBy,
		previous,
		time.Now(),
	)
}

func TaskDomainToHistory(task Task) History {
	return History{
		Title:       task.Title,
		Description: task.Description,
		TaskID:      task.ID,
	}
}
