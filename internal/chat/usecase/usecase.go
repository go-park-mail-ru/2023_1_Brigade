package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"project/internal/chat"
	"project/internal/model"
	"project/internal/pkg/http_utils"
)

type repositoryImpl struct {
	repo chat.Repository
}

func NewChatUsecase(chatRepo chat.Repository) chat.Usecase {
	return &repositoryImpl{repo: chatRepo}
}

func (u *repositoryImpl) CreateChat(ctx context.Context, jsonChatData []byte) http_utils.Response {
	chat := model.Chat{}
	err := json.Unmarshal(jsonChatData, &chat)
	fmt.Println("POST USECASE")
	err = u.repo.InsertChatInDB(ctx, chat)

	if err != nil {
		return err
	}

	return nil
}

func (u *repositoryImpl) GetChatById(ctx context.Context, chatID int) http_utils.Response {
	return u.repo.GetChatInDB(ctx, chatID)
}

func (u *repositoryImpl) GetAllChats(ctx context.Context) http_utils.Response {
	return u.repo.GetAllChatsInDB(ctx)
}

func (u *repositoryImpl) DeleteChatById(ctx context.Context, chatID int) http_utils.Response {
	return u.repo.DeleteChatInDB(ctx, chatID)
}
