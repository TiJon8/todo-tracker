package users_http_transport

import (
	"fmt"
	"net/http"
	"regexp"

	domain "github.com/TiJon8/todo-tracker/internal/core/domain"
	logger "github.com/TiJon8/todo-tracker/internal/core/logger"
	exceptions "github.com/TiJon8/todo-tracker/internal/core/transport/http/exceptions"
	request "github.com/TiJon8/todo-tracker/internal/core/transport/http/request"
	response "github.com/TiJon8/todo-tracker/internal/core/transport/http/response"
	types "github.com/TiJon8/todo-tracker/internal/core/transport/http/types"
	utils "github.com/TiJon8/todo-tracker/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	Name  types.Nullable[string] `json:"name"`
	Phone types.Nullable[string] `json:"phone"`
}

func (p *PatchUserRequest) Validate() error {
	if p.Name.Set {
		if p.Name.Value == nil {
			return fmt.Errorf("Ошибка валидации name")
		}
		nameLen := len([]rune(*p.Name.Value))
		if nameLen <= 2 || nameLen > 100 {
			return fmt.Errorf("Имя не должно быть меньше 2 символов и больше 100: Получено (%d) %w", nameLen, exceptions.BadRequestException)
		}
	}
	if p.Phone.Set {
		if p.Phone.Value != nil {
			phoneLen := len([]rune(*p.Phone.Value))
			if phoneLen < 10 || phoneLen > 16 {
				return fmt.Errorf("Номер телефона не должен быть меньше 10 символов и больше 16: Получено (%d) %w", phoneLen, exceptions.BadRequestException)
			}
			re := regexp.MustCompile(`^\+[0-9]{10,16}$`)
			if !re.MatchString(*p.Phone.Value) {
				return fmt.Errorf(
					"Номер телефона должен соответсвовать формату и начинаться с +: %w",
					exceptions.BadRequestException)
			}
		}
	}
	return nil
}

type PatchUserResponse UserDTO

func (h *UserHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.LoggerFromContext(ctx)
	responseWriter := response.NewHTTPResponseHandler(logger, w)
	logger.Debug("Вызов PatchUser обработчика")

	var RequestData PatchUserRequest
	if err := request.Validate(r, &RequestData); err != nil {
		responseWriter.ResponseError(err, "Не удалось провалидировать структуру")
		return
	}
	id, err := utils.GetPathParam(r, "userId")
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось получить userId")
		return
	}

	patch := patchFromRequest(RequestData)

	user, err := h.UserService.PatchUser(ctx, id, patch)
	if err != nil {
		responseWriter.ResponseError(err, "Не удалось изменить пользователя")
		return
	}

	res := PatchUserResponse(DTOFromDomain(user))
	responseWriter.ResponseJSON(res, http.StatusOK)
}

func patchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewPatchUser(request.Name.ToDomain(), request.Phone.ToDomain())
}
