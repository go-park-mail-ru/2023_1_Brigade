package chat

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error)
	CreateChat(ctx echo.Context, chat model.Chat) (model.Chat, error)
	DeleteChatById(ctx echo.Context, chatID uint64) error
	CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) error
}
