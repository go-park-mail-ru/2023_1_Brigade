package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUsername = status.New(codes.Unknown, errors.New("невалидный username").Error()).Err()
	ErrInvalidNickname = status.New(codes.Unknown, errors.New("невалидный nickname").Error()).Err()
	ErrInvalidEmail    = status.New(codes.Unknown, errors.New("невалидный email").Error()).Err()
	ErrInvalidPassword = status.New(codes.Unknown, errors.New("невалидный пароль").Error()).Err()

	ErrUserIsAlreadyContact     = status.New(codes.Unknown, errors.New("пользователь уже является контактом текущего пользователя").Error()).Err()
	ErrUserIsAlreadyCreated     = status.New(codes.Unknown, errors.New("такой пользователь уже существует").Error()).Err()
	ErrSessionIsAlreadyCreated  = status.New(codes.Unknown, errors.New("такая сессия уже существует").Error()).Err()
	ErrEmailIsAlreadyRegistered = status.New(codes.Unknown, errors.New("такой email уже зарегистрирован").Error()).Err()

	ErrUsernameIsAlreadyRegistered = status.New(codes.Unknown, errors.New("такой username уже зарегистрирован").Error()).Err()
	ErrUserIsAlreadyInChat         = status.New(codes.Unknown, errors.New("такой пользователь уже есть в чате").Error()).Err()

	ErrCookieNotFound = status.New(codes.Unknown, errors.New("cookie не найдена").Error()).Err()

	ErrNotChatAccess = status.New(codes.Unknown, errors.New("такой пользователь не имеет доступа в чат").Error()).Err()

	ErrSessionNotFound   = status.New(codes.Unknown, errors.New("сессия не найдена").Error()).Err()
	ErrChatNotFound      = status.New(codes.Unknown, errors.New("чат не найден").Error()).Err()
	ErrUsernameNotFound  = status.New(codes.Unknown, errors.New("username не найден").Error()).Err()
	ErrEmailNotFound     = status.New(codes.Unknown, errors.New("email не найден").Error()).Err()
	ErrUserNotFound      = status.New(codes.Unknown, errors.New("пользователь не найден").Error()).Err()
	ErrIncorrectPassword = status.New(codes.Unknown, errors.New("неправильный пароль").Error()).Err()
	ErrAvatarNotFound    = status.New(codes.Unknown, errors.New("аватар не найден").Error()).Err()
	ErrImageNotFound     = status.New(codes.Unknown, errors.New("изображение не найдено").Error()).Err()
	ErrMessageNotFound   = status.New(codes.Unknown, errors.New("сообщение не найдено").Error()).Err()
	ErrMembersNotFound   = status.New(codes.Unknown, errors.New("участники не найдены").Error()).Err()

	ErrInternal = errors.New("внутренняя ошибка сервера")
)
