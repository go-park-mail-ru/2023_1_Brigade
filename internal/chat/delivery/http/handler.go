package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/model"
	//	myErrors "project/internal/pkg/errors"

	//	myErrors "project/internal/pkg/errors"
	"project/internal/user"
	"strconv"
)

type chatHandler struct {
	chatUsecase chat.Usecase
	userUsecase user.Usecase
}

func (u chatHandler) GetChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)

	if err != nil {
		return err
	}

	chat, err := u.chatUsecase.GetChatById(ctx, chatID)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	//	err = u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	//	if err == nil {
	//		return myErrors.ErrNotChatAccess
	//	}

	if chat.Type == configs.Chat {
		if chat.Members[0].Id == session.UserId {
			chat.Title = chat.Members[1].Nickname
		} else {
			chat.Title = chat.Members[0].Nickname
		}
	}

	return ctx.JSON(http.StatusOK, chat)
}

func (u chatHandler) GetCurrentUserChatsHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	listUserChats, err := u.chatUsecase.GetListUserChats(ctx, session.UserId)
	if err != nil {
		return err
	}

	for ind := range listUserChats {
		if listUserChats[ind].Type == configs.Chat {
			if listUserChats[ind].Members[0].Id == session.UserId {
				listUserChats[ind].Title = listUserChats[ind].Members[1].Nickname
			} else {
				listUserChats[ind].Title = listUserChats[ind].Members[0].Nickname
			}
		}
	}

	return ctx.JSON(http.StatusOK, listUserChats)
}

func (u chatHandler) CreateCurrentUserChatHandler(ctx echo.Context) error {
	var chat model.CreateChat
	err := ctx.Bind(&chat)
	if err != nil {
		return err
	}

	dbChat, err := u.chatUsecase.CreateChat(ctx, chat)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)

	if chat.Type == configs.Chat {
		if dbChat.Members[0].Id == session.UserId {
			dbChat.Title = dbChat.Members[1].Nickname
		} else {
			dbChat.Title = dbChat.Members[0].Nickname
		}
	}

	return ctx.JSON(http.StatusCreated, dbChat)
}

func (u chatHandler) DeleteChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)
	if err != nil {
		return err
	}

	chat, err := u.chatUsecase.GetChatById(ctx, chatID)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	err = u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	if err == nil {
		return err
	}

	err = u.chatUsecase.DeleteChatById(ctx, chatID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (u chatHandler) EditChatHandler(ctx echo.Context) error {
	var chat model.EditChat
	err := ctx.Bind(&chat)
	if err != nil {
		return err
	}

	newChat, err := u.chatUsecase.EditChat(ctx, chat)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, newChat)
}

func NewChatHandler(e *echo.Echo, chatUsecase chat.Usecase, userUsecase user.Usecase) chatHandler {
	handler := chatHandler{chatUsecase: chatUsecase, userUsecase: userUsecase}
	currentUserChatsUrl := "/chats/"
	getChatUrl := "/chats/:chatID/"
	deleteChatUrl := "/chats/:chatID/"

	api := e.Group("api/v1")

	getChat := api.Group(getChatUrl)
	currentUserChats := api.Group(currentUserChatsUrl)
	deleteChat := api.Group(deleteChatUrl)

	getChat.GET("", handler.GetChatHandler)
	deleteChat.DELETE("", handler.DeleteChatHandler)
	currentUserChats.PUT("", handler.EditChatHandler)
	currentUserChats.GET("", handler.GetCurrentUserChatsHandler)
	currentUserChats.POST("", handler.CreateCurrentUserChatHandler)

	return handler
}
