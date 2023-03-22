package usecase

import (
	"github.com/labstack/echo/v4"
	"project/internal/chat"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

type usecase struct {
	repo chat.Repository
}

func NewChatUsecase(chatRepo chat.Repository) chat.Usecase {
	return &usecase{repo: chatRepo}
}

func (u *usecase) CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) error {
	members := chat.Members
	userInChat := false
	for _, member := range members {
		if member.Id == userID {
			userInChat = true
			break
		}
	}

	if !userInChat {
		return myErrors.ErrNotChatAccess
	}

	return nil
}

func (u *usecase) GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error) {
	chat, err := u.repo.GetChatById(ctx, chatID)
	if err != nil {
		return chat, err
	}

	return chat, err
}

func (u *usecase) CreateChat(ctx echo.Context, chat model.Chat) (model.Chat, error) {
	chat, err := u.repo.CreateChat(ctx, chat)
	return chat, err
}

func (u *usecase) DeleteChatById(ctx echo.Context, chatID uint64) error {
	err := u.repo.DeleteChatById(ctx, chatID)
	return err
}
