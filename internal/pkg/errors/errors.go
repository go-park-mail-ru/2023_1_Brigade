package errors

import (
	"errors"
)

var (
	InvalidUsername = errors.New("Invalid username")
	InvalidEmail    = errors.New("Invalid email")
	InvalidName     = errors.New("Invalid name")
	InvalidPassword = errors.New("Invalid password")

	EmailIsAlreadyRegistred    = errors.New("The email is already registered")
	UsernameIsAlreadyRegistred = errors.New("The username is already registered")

	NoUserFound       = errors.New("No user found")
	IncorrectPassword = errors.New("Incorrect password")

	InternalError = errors.New("Internal error")
)
