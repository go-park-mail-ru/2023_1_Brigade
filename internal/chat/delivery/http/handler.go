package http

import (
	"github.com/labstack/echo/v4"
	"project/internal/chat"
)

type chatHandler struct {
	usecase chat.Usecase
}

func (u *chatHandler) GetChatHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (u *chatHandler) CreateChatHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (u *chatHandler) DeleteChatHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func NewChatHandler(e *echo.Echo, us chat.Usecase) chatHandler {
	handler := chatHandler{usecase: us}
	getChatUrl := "/chats/{chatID:[0-9]+}"
	createChatUrl := "/chats/"
	deleteChatUrl := "/chats/"

	e.OPTIONS(getChatUrl, handler.GetChatHandler())
	e.OPTIONS(createChatUrl, handler.CreateChatHandler())
	e.OPTIONS(deleteChatUrl, handler.DeleteChatHandler())

	e.GET(getChatUrl, handler.GetChatHandler())
	e.POST(createChatUrl, handler.CreateChatHandler())
	e.DELETE(deleteChatUrl, handler.DeleteChatHandler())

	return handler
}
