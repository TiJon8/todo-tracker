package tasks_transport_http

import (
	"fmt"
	"net/http"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	types "github.com/TiJon8/todo-tracker/internal/core/transport/http/types"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
	"github.com/google/uuid"
)

type PatchTaskRequest struct {
	Title       types.Nullable[string] `json:"title"`
	Description types.Nullable[string] `json:"description"`
	Completed   types.Nullable[bool]   `json:"completed"`
	AuthorID    types.Nullable[string] `json:"author_id" validate:"uuid4"`
}

func (p *PatchTaskRequest) Validate() error {
	if p.Title.Value == nil {
		return fmt.Errorf("Title не должен быть пустым")
	}
	titleLen := len([]rune(*p.Title.Value))
	if titleLen < 1 && titleLen > 100 {
		return fmt.Errorf("Title не должен быть меньше 1 символа и больше 100 ")
	}

	if p.Description.Set {
		if p.Description.Value != nil {
			descLen := len([]rune(*p.Description.Value))
			if descLen > 1000 {
				return fmt.Errorf("Поле description не должно быть больше 1000 символов")
			}
		}
	}
	if p.Completed.Set {
		if p.Completed.Value == nil {
			return fmt.Errorf("Поле completed должно быть типа bool [true|false]")
		}
	}
	if !p.AuthorID.Set {
		return fmt.Errorf("author_id должен быть!")
	}
	if err := uuid.Validate(*p.AuthorID.Value); err != nil {
		return fmt.Errorf("Формат поля author_id не соответствует uuid: %v: %w", err, exceptions.BadRequestException)
	}
	return nil
}

type PatchTaskResponse TaskDTO

func (h *TaskHTTPHandlers) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов CreateTask обработчика")

	// var PatchTaskAuthorId PatchTaskRequest
	// if err := request.Validate(r, &PatchTaskAuthorId.AuthorID); err != nil {
	// 	responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
	// 	return
	// }

	var PatchTask PatchTaskRequest
	if err := request.Validate(r, &PatchTask); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	}

	taskId, err := utils.GetPathParam(r, "taskId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить userId")
		return
	}

	taskPatch := patchFromRequest(PatchTask)

	patchedTask, err := h.TaskService.PatchTask(ctx, *PatchTask.AuthorID.Value, taskId, taskPatch)
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось изменить пользователя")
		return
	}

	res := PatchTaskResponse(DTOFromDomain(patchedTask))
	responseWriter.ResponseJSON(res, http.StatusOK)
}

func patchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.TaskPatch{
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		Completed:   request.Completed.ToDomain(),
	}
}
