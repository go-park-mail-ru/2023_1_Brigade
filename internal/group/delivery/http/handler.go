package http

import (
	"github.com/labstack/echo/v4"
	"project/internal/group"
)

type groupHadnler struct {
	usecase group.Usecase
}

func (u *groupHadnler) GetGroupHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (u *groupHadnler) CreateGroupHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (u *groupHadnler) DeleteGroupHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func NewChatHandler(e *echo.Echo, us group.Usecase) groupHadnler {
	handler := groupHadnler{usecase: us}
	getChatUrl := "/groups/{groupID:[0-9]+}"
	createChatUrl := "/groups/"
	deleteChatUrl := "/groups/"

	e.OPTIONS(getChatUrl, handler.GetGroupHandler())
	e.OPTIONS(createChatUrl, handler.CreateGroupHandler())
	e.OPTIONS(deleteChatUrl, handler.DeleteGroupHandler())

	e.GET(getChatUrl, handler.CreateGroupHandler())
	e.POST(createChatUrl, handler.CreateGroupHandler())
	e.DELETE(deleteChatUrl, handler.DeleteGroupHandler())

	return handler
}
