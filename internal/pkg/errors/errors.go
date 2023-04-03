package errors

import (
	"errors"
)

var (
	ErrInvalidUsername = errors.New("невалидный username")
	ErrInvalidNickname = errors.New("невалидный nickname")
	ErrInvalidEmail    = errors.New("невалидный email")
	ErrInvalidPassword = errors.New("невалидный пароль")

	ErrUserIsAlreadyContact        = errors.New("пользователь уже является контактом текущего пользователя")
	ErrUserIsAlreadyCreated        = errors.New("такой пользователь уже существует")
	ErrSessionIsAlreadyCreated     = errors.New("такая сессия уже существует")
	ErrEmailIsAlreadyRegistered    = errors.New("такой email уже зарегистрирован")
	ErrUsernameIsAlreadyRegistered = errors.New("такой username уже зарегистрирован")
	ErrUserIsAlreadyInChat         = errors.New("такой пользователь уже есть в чате")

	ErrCookieNotFound = errors.New("cookie не найдена")

	ErrNotChatAccess = errors.New("такой пользователь не имеет доступа в чат")

	ErrSessionNotFound   = errors.New("сессия не найдена")
	ErrChatNotFound      = errors.New("чат не найден")
	ErrUsernameNotFound  = errors.New("username не найден")
	ErrEmailNotFound     = errors.New("email не найден")
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrIncorrectPassword = errors.New("неправильный пароль")
	ErrAvatarNotFound    = errors.New("аватар не найден")
	ErrImageNotFound     = errors.New("изображение не найдено")

	ErrInternal = errors.New("внутренняя ошибка сервера")
)
