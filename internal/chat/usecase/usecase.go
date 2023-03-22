package usecase

import (
	"github.com/labstack/echo/v4"
	"project/internal/chat"
	"project/internal/model"
	"project/internal/user"
)

type usecase struct {
	chatRepo chat.Repository
	userRepo user.Repository
}

func NewChatUsecase(chatRepo chat.Repository, userRepo user.Repository) chat.Usecase {
	return &usecase{chatRepo: chatRepo, userRepo: userRepo}
}

func (u *usecase) CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) bool {
	members := chat.Members
	userInChat := false
	for _, member := range members {
		if member.Id == userID {
			userInChat = true
			break
		}
	}

	return userInChat
}

func (u *usecase) GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error) {
	chat, err := u.chatRepo.GetChatById(ctx, chatID)
	if err != nil {
		return chat, err
	}

	return chat, err
}

func (u *usecase) CreateChat(ctx echo.Context, chat model.CreateChat) (model.Chat, error) {
	var members []model.User
	for _, userID := range chat.Members {
		user, err := u.userRepo.GetUserById(ctx, userID)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, user)
	}

	dbChat := model.Chat{
		Title:   chat.Title,
		Members: members,
	}
	dbChat, err := u.chatRepo.CreateChat(ctx, dbChat)

	return dbChat, err
}

func (u *usecase) DeleteChatById(ctx echo.Context, chatID uint64) error {
	err := u.chatRepo.DeleteChatById(ctx, chatID)
	return err
}
