package http

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/model"
	httpUtils "project/internal/pkg/http_utils"

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

	session := ctx.Get("session").(model.Session)
	chat, err := u.chatUsecase.GetChatById(context.TODO(), chatID, session.UserId)
	if err != nil {
		return err
	}

	if chat.Type == configs.Chat {
		if len(chat.Members) > 0 {
			if chat.Members[0].Id == session.UserId {
				chat.Title = chat.Members[1].Nickname
				chat.Avatar = chat.Members[1].Avatar
			} else {
				chat.Title = chat.Members[0].Nickname
				chat.Avatar = chat.Members[0].Avatar
			}
		}
	}

	chat = httpUtils.SanitizeStruct(chat).(model.Chat)

	return ctx.JSON(http.StatusOK, chat)
}

func (u chatHandler) GetCurrentUserChatsHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	listUserChats, err := u.chatUsecase.GetListUserChats(context.TODO(), session.UserId)
	if err != nil {
		return err
	}

	for ind := range listUserChats {
		if listUserChats[ind].Type == configs.Chat {
			if listUserChats[ind].Members[0].Id == session.UserId {
				listUserChats[ind].Title = listUserChats[ind].Members[1].Nickname
				listUserChats[ind].Avatar = listUserChats[ind].Members[1].Avatar
			} else {
				listUserChats[ind].Title = listUserChats[ind].Members[0].Nickname
				listUserChats[ind].Avatar = listUserChats[ind].Members[0].Avatar
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

	session := ctx.Get("session").(model.Session)

	dbChat, err := u.chatUsecase.CreateChat(context.TODO(), chat, session.UserId)
	if err != nil {
		return err
	}

	if chat.Type == configs.Chat {
		if len(dbChat.Members) > 0 {
			if dbChat.Members[0].Id == session.UserId {
				dbChat.Title = dbChat.Members[1].Nickname
				dbChat.Avatar = dbChat.Members[1].Avatar
			} else {
				dbChat.Title = dbChat.Members[0].Nickname
				dbChat.Avatar = dbChat.Members[0].Avatar
			}
		}
	}

	dbChat.MasterID = session.UserId
	return ctx.JSON(http.StatusCreated, dbChat)
}

func (u chatHandler) DeleteChatHandler(ctx echo.Context) error {
	chatID, err := strconv.ParseUint(ctx.Param("chatID"), 10, 64)
	if err != nil {
		return err
	}

	err = u.chatUsecase.DeleteChatById(context.TODO(), chatID)
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

	newChat, err := u.chatUsecase.EditChat(context.TODO(), chat)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, newChat)
}

func (u chatHandler) GetChatsMessagesHandler(ctx echo.Context) error {
	string := ctx.Param("string")
	session := ctx.Get("session").(model.Session)
	string, err := url.QueryUnescape(string)
	if err != nil {
		log.Error(err)
	}

	searchChats, err := u.chatUsecase.GetSearchChatsMessagesChannels(context.TODO(), session.UserId, string)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, searchChats)
}

func NewChatHandler(e *echo.Echo, chatUsecase chat.Usecase, userUsecase user.Usecase) chatHandler {
	handler := chatHandler{chatUsecase: chatUsecase, userUsecase: userUsecase}
	currentUserChatsUrl := "/chats/"
	getChatUrl := "/chats/:chatID/"
	deleteChatUrl := "/chats/:chatID/"
	searchChatsMessagesUrl := "/chats/search/:string/"

	api := e.Group("api/v1")

	getChat := api.Group(getChatUrl)
	currentUserChats := api.Group(currentUserChatsUrl)
	deleteChat := api.Group(deleteChatUrl)
	searchChatsMessages := api.Group(searchChatsMessagesUrl)

	getChat.GET("", handler.GetChatHandler)
	deleteChat.DELETE("", handler.DeleteChatHandler)
	currentUserChats.PUT("", handler.EditChatHandler)
	currentUserChats.GET("", handler.GetCurrentUserChatsHandler)
	currentUserChats.POST("", handler.CreateCurrentUserChatHandler)
	searchChatsMessages.GET("", handler.GetChatsMessagesHandler)

	return handler
}
