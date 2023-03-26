package chat

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Repository interface {
	GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error)
	CreateChat(ctx echo.Context, chat model.Chat) (model.Chat, error)
	DeleteChatById(ctx echo.Context, chatID uint64) error
	AddUserInChatDB(ctx echo.Context, chatID uint64, memberID uint64) error
}
