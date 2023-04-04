package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/chat"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
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

	//session := ctx.Get("session").(model.Session)
	session := model.Session{UserId: 1}
	err = u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	if err != nil {
		return myErrors.ErrNotChatAccess
	}

	return ctx.JSON(http.StatusOK, chat)
}

func (u chatHandler) GetCurrentUserChatsHandler(ctx echo.Context) error {
	//session := ctx.Get("session").(model.Session)
	session := model.Session{UserId: 1}
	listUserChats, err := u.chatUsecase.GetListUserChats(ctx, session.UserId)
	if err != nil {
		return err
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

	//session := ctx.Get("session").(model.Session)
	session := model.Session{UserId: 1}
	err = u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	if err != nil {
		return err
	}

	err = u.chatUsecase.DeleteChatById(ctx, chatID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (u chatHandler) AddUserInChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)
	if err != nil {
		return err
	}

	userID, err := strconv.ParseUint(ctx.Param("userID"), 10, 64)
	if err != nil {
		return err
	}

	err = u.chatUsecase.AddUserInChat(ctx, chatID, userID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func NewChatHandler(e *echo.Echo, chatUsecase chat.Usecase, userUsecase user.Usecase) chatHandler {
	handler := chatHandler{chatUsecase: chatUsecase, userUsecase: userUsecase}
	currentUserChatsUrl := "/chats/"
	getChatUrl := "/chats/:chatID"
	deleteChatUrl := "/chats/:chatID"
	addUserInChatUrl := "/api/v1/chats/:chatID/add/:userID"

	api := e.Group("api/v1")

	getChat := api.Group(getChatUrl)
	currentUserChats := api.Group(currentUserChatsUrl)
	deleteChat := api.Group(deleteChatUrl)
	addUserInChat := api.Group(addUserInChatUrl)

	getChat.GET("", handler.GetChatHandler)
	deleteChat.DELETE("", handler.DeleteChatHandler)
	addUserInChat.POST("", handler.AddUserInChatHandler)
	currentUserChats.GET("", handler.GetCurrentUserChatsHandler)
	currentUserChats.POST("", handler.CreateCurrentUserChatHandler)

	return handler
}
