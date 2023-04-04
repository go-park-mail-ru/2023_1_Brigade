package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/user"
)

type usecase struct {
	chatRepo     chat.Repository
	userRepo     user.Repository
	messagesRepo messages.Repository
}

func NewChatUsecase(chatRepo chat.Repository, userRepo user.Repository, messagesRepo messages.Repository) chat.Usecase {
	return &usecase{chatRepo: chatRepo, userRepo: userRepo, messagesRepo: messagesRepo}
}

func (u usecase) CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) error {
	members := chat.Members
	for _, member := range members {
		if member.Id == userID {
			return myErrors.ErrUserIsAlreadyInChat
		}
	}

	return nil
}

func (u usecase) GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error) {
	chat, err := u.chatRepo.GetChatById(context.Background(), chatID)
	return chat, err
}

//func (u usecase) GetUserChats(ctx echo.Context, userID uint64) ([]model.Chat, error) {
//	chats, err := u.chatRepo.GetChatsByUserId(context.Background(), userID)
//	return chats, err
//}

func (u usecase) CreateChat(ctx echo.Context, chat model.CreateChat) (model.Chat, error) {
	var members []model.User
	for _, userID := range chat.Members {
		user, err := u.userRepo.GetUserById(context.Background(), userID)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	createdChat := model.Chat{
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   configs.DefaultAvatarUrl,
		Members:  members,
		Messages: []model.Message{},
	}
	chatFromDB, err := u.chatRepo.CreateChat(context.Background(), createdChat)

	return chatFromDB, err
}

func (u usecase) DeleteChatById(ctx echo.Context, chatID uint64) error {
	err := u.chatRepo.DeleteChatById(context.Background(), chatID)
	return err
}

func (u usecase) AddUserInChat(ctx echo.Context, chatID uint64, userID uint64) error {
	chat, err := u.GetChatById(ctx, chatID)
	if err != nil {
		return err
	}

	err = u.CheckExistUserInChat(ctx, chat, userID)
	if err != nil {
		return err
	}

	err = u.userRepo.CheckExistUserById(context.Background(), userID)
	if err != nil {
		return err
	}

	err = u.chatRepo.AddUserInChatDB(context.Background(), chatID, userID)
	return err
}

func (u usecase) GetListUserChats(ctx echo.Context, userID uint64) ([]model.ChatInListUser, error) {
	var chatsInListUser []model.ChatInListUser
	userChats, err := u.chatRepo.GetChatsByUserId(context.Background(), userID)
	log.Warn(userChats)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, userChat := range userChats {
		chat, err := u.chatRepo.GetChatById(context.Background(), userChat.ChatId)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(context.Background(), chat.Id)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		lastMessageAuthor, err := u.userRepo.GetUserById(context.Background(), lastMessage.AuthorId)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		chatsInListUser = append(chatsInListUser, model.ChatInListUser{
			Id:                chat.Id,
			Type:              chat.Type,
			Title:             chat.Title,
			Avatar:            chat.Avatar,
			LastMessage:       lastMessage,
			LastMessageAuthor: model_conversion.FromAuthorizedUserToUser(lastMessageAuthor),
		})
	}

	return chatsInListUser, nil
}
