package http_utils

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"reflect"
	"time"
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, myErrors.ErrInvalidUsername):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrInvalidEmail):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrInvalidPassword):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrCookieNotFound):
		return http.StatusUnauthorized
	case errors.Is(err, myErrors.ErrNotChatAccess):
		return http.StatusForbidden
	case errors.Is(err, myErrors.ErrSessionNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrEmailNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrUsernameNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrChatNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrIncorrectPassword):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrUserIsAlreadyContact):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrUsernameIsAlreadyRegistered):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrEmailIsAlreadyRegistered):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrSessionIsAlreadyCreated):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrUserIsAlreadyInChat):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func SetCookie(ctx echo.Context, session model.Session) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    session.Cookie,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Hour),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	ctx.SetCookie(cookie)
}

func DeleteCookie(ctx echo.Context) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	ctx.SetCookie(cookie)
}

func SanitizeStruct(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
	inputType := inputValue.Type()

	outputValue := reflect.New(inputType).Elem()

	p := bluemonday.UGCPolicy()

	for i := 0; i < inputValue.NumField(); i++ {
		inputFieldValue := inputValue.Field(i)

		if inputFieldValue.Kind() == reflect.String {
			outputValue.Field(i).SetString(p.Sanitize(inputFieldValue.String()))
		} else {
			outputValue.Field(i).Set(inputFieldValue)
		}
	}

	return outputValue.Interface()
}
