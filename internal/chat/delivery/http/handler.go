package http

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/auth"
	"project/internal/chat"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strconv"
)

type chatHandler struct {
	chatUsecase chat.Usecase
	authUsecase auth.Usecase
}

func (u *chatHandler) GetChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)

	if err != nil {
		return err
	}

	chat, err := u.chatUsecase.GetChatById(ctx, chatID)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	userInChat := u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	if !userInChat {
		return myErrors.ErrNotChatAccess
	}

	return ctx.JSON(http.StatusOK, chat)
}

func (u *chatHandler) CreateChatHandler(ctx echo.Context) error {
	var chat model.CreateChat
	body := ctx.Get("body").([]byte)

	err := json.Unmarshal(body, &chat)
	if err != nil {
		return err
	}

	dbChat, err := u.chatUsecase.CreateChat(ctx, chat)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, dbChat)
}

func (u *chatHandler) DeleteChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)
	if err != nil {
		return err
	}

	chat, err := u.chatUsecase.GetChatById(ctx, chatID)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	userInChat := u.chatUsecase.CheckExistUserInChat(ctx, chat, session.UserId)
	if !userInChat {
		return myErrors.ErrNotChatAccess
	}

	err = u.chatUsecase.DeleteChatById(ctx, chatID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewChatHandler(e *echo.Echo, chatUsecase chat.Usecase, authUsecase auth.Usecase) chatHandler {
	handler := chatHandler{chatUsecase: chatUsecase, authUsecase: authUsecase}
	getChatUrl := "/chats/:chatID"
	deleteChatUrl := "/chats/:chatID"
	createChatUrl := "/chats/"

	api := e.Group("api/v1")

	getChat := api.Group(getChatUrl)
	createChat := api.Group(createChatUrl)
	deleteChat := api.Group(deleteChatUrl)

	getChat.OPTIONS("", handler.GetChatHandler)
	createChat.OPTIONS("", handler.CreateChatHandler)
	deleteChat.OPTIONS("", handler.DeleteChatHandler)

	getChat.GET("", handler.GetChatHandler)
	createChat.POST("", handler.CreateChatHandler)
	deleteChat.DELETE("", handler.DeleteChatHandler)

	return handler
}
