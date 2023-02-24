package repository

import (
	"fmt"
	"project/internal/chat"
	"project/internal/model"
)

func NewChatMemoryRepository() chat.Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct{}

func (r *repositoryImpl) InsertChatInDB(chat model.Chat) (model.Chat, error) {
	fmt.Println("POST CHAT")
	return model.Chat{}, nil
}

func (r *repositoryImpl) GetChatInDB(chatID int) (model.Chat, error) {
	fmt.Println("GET ID CHAT")
	return model.Chat{}, nil
}

func (r *repositoryImpl) GetAllChatsInDB() ([]model.Chat, error) {
	fmt.Println("GET ALL CHATS")
	return []model.Chat{}, nil
}

func (r *repositoryImpl) DeleteChatInDB(chatID int) error {
	fmt.Println("DELETE CHAT")
	return nil
}
