package repository

import (
	"context"
	"fmt"
	"project/internal/chat"
	"project/internal/model"
)

func NewChatMemoryRepository() chat.Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct{}

func (r *repositoryImpl) InsertChatInDB(ctx context.Context, chat model.Chat) (model.Chat, error) {
	fmt.Println("POST CHAT")
	return model.Chat{}, nil
}

func (r *repositoryImpl) GetChatInDB(ctx context.Context, chatID int) (model.Chat, error) {
	fmt.Println("GET ID CHAT")
	return model.Chat{}, nil
}

func (r *repositoryImpl) GetAllChatsInDB(ctx context.Context) ([]model.Chat, error) {
	fmt.Println("GET ALL CHATS")
	return []model.Chat{}, nil
}

func (r *repositoryImpl) DeleteChatInDB(ctx context.Context, chatID int) error {
	fmt.Println("DELETE CHAT")
	return nil
}
