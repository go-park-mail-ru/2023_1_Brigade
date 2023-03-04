package http_utils

import (
	"net/http"
	err "project/internal/pkg/errors"
)

var error2HttpCode = []struct {
	Error    error
	HttpCode int
}{
	{
		Error:    err.UserIdGiven,
		HttpCode: http.StatusOK,
	},
	{
		Error:    err.UserCreated,
		HttpCode: http.StatusCreated,
	},
	{
		Error:    err.SessionSuccessDeleted,
		HttpCode: http.StatusNoContent,
	},
	{
		Error:    err.ErrInvalidUsername,
		HttpCode: http.StatusBadRequest,
	},
	{
		Error:    err.ErrInvalidEmail,
		HttpCode: http.StatusBadRequest,
	},
	{
		Error:    err.ErrInvalidName,
		HttpCode: http.StatusBadRequest,
	},
	{
		Error:    err.ErrInvalidPassword,
		HttpCode: http.StatusBadRequest,
	},
	{
		Error:    err.ErrEmailIsAlreadyRegistred,
		HttpCode: http.StatusConflict,
	},
	{
		Error:    err.ErrUsernameIsAlreadyRegistred,
		HttpCode: http.StatusConflict,
	},
	{
		Error:    err.ErrCookieNotFound,
		HttpCode: http.StatusUnauthorized,
	},
	{
		Error:    err.ErrSessionNotFound,
		HttpCode: http.StatusNotFound,
	},
	{
		Error:    err.ErrUserNotFound,
		HttpCode: http.StatusNotFound,
	},
	{
		Error:    err.ErrIncorrectPassword,
		HttpCode: http.StatusNotFound,
	},
}
