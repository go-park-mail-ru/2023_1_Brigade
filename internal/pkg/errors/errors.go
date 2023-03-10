package errors

import (
	"errors"
)

var (
	UserGetting           = errors.New("Successful return of the user")
	UserCreated           = errors.New("User created")
	SessionSuccessDeleted = errors.New("Session success deleted")

	ErrInvalidUsername = errors.New("Invalid username")
	ErrInvalidEmail    = errors.New("Invalid email")
	ErrInvalidName     = errors.New("Invalid name")
	ErrInvalidPassword = errors.New("Invalid password")

	ErrUserIsAlreadyCreated       = errors.New("The user is already created")
	ErrSessionIsAlreadyCreated    = errors.New("The session is already created")
	ErrEmailIsAlreadyRegistred    = errors.New("The email is already registered")
	ErrUsernameIsAlreadyRegistred = errors.New("The username is already registered")

	ErrCookieNotFound = errors.New("Cookie not found")

	ErrSessionNotFound   = errors.New("Session not found")
	ErrUserNotFound      = errors.New("User not found")
	ErrIncorrectPassword = errors.New("Incorrect password")

	ErrInternal = errors.New("Internal error")
)
