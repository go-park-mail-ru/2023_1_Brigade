package usecase

import (
	"context"
	"encoding/json"
	"project/internal/chat"
	"project/internal/model"
)

type repositoryImpl struct {
	repo chat.Repository
}

func NewChatUsecase(chatRepo chat.Repository) chat.Usecase {
	return &repositoryImpl{repo: chatRepo}
}

func (u *repositoryImpl) CreateChat(ctx context.Context, jsonChatData []byte) (model.Chat, error) {
	chat := model.Chat{}
	err := json.Unmarshal(jsonChatData, &chat)

	if err != nil {
		return chat, err
	}

	return u.repo.InsertChatInDB(ctx, chat)
}

func (u *repositoryImpl) GetChatById(ctx context.Context, chatID int) (model.Chat, error) {
	return u.repo.GetChatInDB(ctx, chatID)
}

func (u *repositoryImpl) GetAllChats(ctx context.Context) ([]model.Chat, error) {
	return u.repo.GetAllChatsInDB(ctx)
}

func (u *repositoryImpl) DeleteChatById(ctx context.Context, chatID int) error {
	return u.repo.DeleteChatInDB(ctx, chatID)
}
