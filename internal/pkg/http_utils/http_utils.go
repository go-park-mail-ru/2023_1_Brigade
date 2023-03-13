package http_utils

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"time"
)

type jsonError struct {
	Err error
}

func (j jsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Err.Error())
}

func statusCode(err error) int {
	switch {
	case errors.Is(err, myErrors.UserGetting):
		return http.StatusOK
	case errors.Is(err, myErrors.UserCreated):
		return http.StatusCreated
	case errors.Is(err, myErrors.SessionSuccessDeleted):
		return http.StatusNoContent
	case errors.Is(err, myErrors.ErrInvalidUsername):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrInvalidEmail):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrInvalidPassword):
		return http.StatusBadRequest
	case errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrUsernameIsAlreadyRegistred):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrSessionIsAlreadyCreated):
		return http.StatusConflict
	case errors.Is(err, myErrors.ErrCookieNotFound):
		return http.StatusUnauthorized
	case errors.Is(err, myErrors.ErrSessionNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, myErrors.ErrIncorrectPassword):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func SendJsonError(ctx echo.Context, err error) error {
	response := jsonError{Err: err}
	jsonResponse, marshalError := json.Marshal(&response)

	if marshalError != nil {
		log.Error(marshalError.Error())
		return ctx.NoContent(statusCode(myErrors.ErrInternal))
	}

	log.Error(string(jsonResponse))
	return ctx.JSONBlob(statusCode(err), jsonResponse)
}

func SendJsonUser(ctx echo.Context, user model.User, status error) error {
	jsonResponse, marshalError := json.Marshal(&user)
	if marshalError != nil {
		log.Error(marshalError.Error())
		return ctx.NoContent(statusCode(myErrors.ErrInternal))
	}

	return ctx.JSONBlob(statusCode(status), jsonResponse)
}

func ParsingIdUrl(ctx echo.Context, param string) (uint64, error) {
	log.Error(ctx.Get("userID"))

	return 0, nil
	//vars := mux.Vars(r)
	//log.Fatal(vars[param])
	//return strconv.ParseUint(vars[param], 10, 64)
}

func SetCookie(ctx echo.Context, session model.Session) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    session.Cookie,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Hour),
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
	}
	ctx.SetCookie(cookie)
}
