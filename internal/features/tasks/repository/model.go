package task_repository_postgres

import (
	"time"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

type TaskModel struct {
	ID          string
	Version     int
	Title       string
	Description *string
	Completed   bool
	CompletedAt *time.Time
	AuthorID    string
	GroupID     *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt *time.Time
}

func DomainFromModelSlice(tasks []TaskModel) []domain.Task {
	s := make([]domain.Task, len(tasks))

	for i, t := range tasks {
		s[i] = domain.NewTask(t.ID, t.Version, t.Title, t.Description, t.Completed, t.CompletedAt, t.AuthorID, t.GroupID, t.CreatedAt, t.UpdatedAt, t.DeletedAt)
	}
	return s
}
