package history_service

import (
	"context"
	"fmt"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
)

func (s *HistoryService) CreateSnapshot(ctx context.Context, event string, task domain.Task) error {
	var status domain.EnumTaskStatus
	switch event {
	case "create":
		status = domain.InProgress
	case "patch":
		if task.CompletedAt != nil {
			status = domain.Done
		} else if task.CompletedAt == nil {
			status = domain.InProgress
		}
	case "done":
		status = domain.Done
	case "delete":
		status = domain.Deleted
	}

	var (
		domainHistory domain.History
	)

	if event != "create" {
		snapshot, err := s.GetSnapshot(ctx, task.ID)
		if err != nil {
			return fmt.Errorf("Get snapshot error: %w", err)
		}
		domainHistory = domain.NewHistoryUninitialized(task.Title, task.Description, task.ID, status, task.AuthorID, &snapshot.ID)
	} else {
		domainHistory = domain.NewHistoryUninitialized(task.Title, task.Description, task.ID, status, task.AuthorID, nil)
	}
	_, err := s.Repository.CreateSnapshot(ctx, domainHistory)
	if err != nil {
		return fmt.Errorf("history snapshot create: %w", err)
	}

	return nil
}
