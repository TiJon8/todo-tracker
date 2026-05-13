package core_domain

import (
	"fmt"
	"time"

	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
)

type Task struct {
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
	DeletedAt   *time.Time
}

func NewTask(
	id string,
	version int,
	title string,
	description *string,
	completed bool,
	completedAt *time.Time,
	authorId string,
	groupId *string,
	createdAt time.Time,
	updatedAt *time.Time,
	deletedAt *time.Time,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CompletedAt: completedAt,
		AuthorID:    authorId,
		GroupID:     groupId,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}

func NewTaskUninitialized(title string, description *string, authorId string, groupId *string) Task {
	return NewTask(
		uninitializedID,
		uninitializedVersion,
		title,
		description,
		false,
		nil,
		authorId,
		groupId,
		time.Now(),
		nil,
		nil,
	)
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("Title не должен быть пустым: %w", exceptions.BadRequestException)
	}
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf("Title не должен быть больше 100 символов: получено: %d; %w", titleLen, exceptions.BadRequestException)
	}
	if t.AuthorID == "" {
		return fmt.Errorf("author_id не может быть пустым: %w", exceptions.BadRequestException)
	}
	// if !(t.Completed == false && t.CompletedAt == nil) || !(t.Completed == true && t.CompletedAt != nil && t.CompletedAt.Unix() >= t.CreatedAt.Unix()) {
	// 	return fmt.Errorf("Completed не может быть null и при этом CompletedAt не null и наоборот, а также CompletedAt не может быть больше CreatedAt: %w", exceptions.BadRequestException)
	// }
	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskUser(title Nullable[string], description Nullable[string], completed Nullable[bool]) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title' can't be null: %w", exceptions.BadRequestException)
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be null. completed must be type of bool [true|false]: %w", exceptions.BadRequestException)
	}
	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("Не удалось провалидировать структуру: %w", err)
	}
	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}
	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value
		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("Не валидный user patch: %w", err)
	}
	*t = tmp

	return nil
}
