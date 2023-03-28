package errors

import (
	"errors"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")

	ErrUserIsAlreadyContact       = errors.New("the contact is already friend")
	ErrUserIsAlreadyCreated       = errors.New("the user is already created")
	ErrSessionIsAlreadyCreated    = errors.New("the session is already created")
	ErrEmailIsAlreadyRegistred    = errors.New("the email is already registered")
	ErrUsernameIsAlreadyRegistred = errors.New("the username is already registered")
	ErrUserIsAlreadyInChat        = errors.New("user is already in chat")

	ErrCookieNotFound = errors.New("cookie not found")

	ErrNotChatAccess = errors.New("this user dont have access to this chat")

	ErrSessionNotFound   = errors.New("session not found")
	ErrChatNotFound      = errors.New("chat not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrAvatarNotFound    = errors.New("avatar not found")

	ErrInternal = errors.New("internal error")
)
