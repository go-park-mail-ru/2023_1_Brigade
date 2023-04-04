package chat

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error)
	//GetUserChats(ctx echo.Context, userID uint64) ([]model.Chat, error)
	CreateChat(ctx echo.Context, chat model.CreateChat) (model.Chat, error)
	DeleteChatById(ctx echo.Context, chatID uint64) error
	AddUserInChat(ctx echo.Context, chatID uint64, memberID uint64) error
	CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) error
	GetListUserChats(ctx echo.Context, userID uint64) ([]model.ChatInListUser, error)
}
