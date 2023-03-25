package errors

import (
	"errors"
)

var (
	ErrInvalidUsername = errors.New("Invalid username")
	ErrInvalidEmail    = errors.New("Invalid email")
	//ErrInvalidName     = errors.New("Invalid name")
	ErrInvalidPassword = errors.New("Invalid password")

	ErrUserIsAlreadyContact       = errors.New("The contact is already friend")
	ErrUserIsAlreadyCreated       = errors.New("The user is already created")
	ErrSessionIsAlreadyCreated    = errors.New("The session is already created")
	ErrEmailIsAlreadyRegistred    = errors.New("The email is already registered")
	ErrUsernameIsAlreadyRegistred = errors.New("The username is already registered")

	ErrCookieNotFound = errors.New("Cookie not found")

	ErrNotChatAccess = errors.New("This user dont have access to this chat")

	ErrSessionNotFound   = errors.New("Session not found")
	ErrChatNotFound      = errors.New("Chat not found")
	ErrUserNotFound      = errors.New("User not found")
	ErrIncorrectPassword = errors.New("Incorrect password")

	ErrInternal = errors.New("Internal error")
)
