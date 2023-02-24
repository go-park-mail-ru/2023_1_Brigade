package usecase

import (
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

func (u *repositoryImpl) CreateChat(jsonChatData []byte) (model.Chat, error) {
	chat := model.Chat{}
	err := json.Unmarshal(jsonChatData, &chat)

	if err != nil {
		return chat, err
	}

	return u.repo.InsertChatInDB(chat)
}

func (u *repositoryImpl) GetChatById(chatID int) (model.Chat, error) {
	return u.repo.GetChatInDB(chatID)
}

func (u *repositoryImpl) GetAllChats() ([]model.Chat, error) {
	return u.repo.GetAllChatsInDB()
}

func (u *repositoryImpl) DeleteChatById(chatID int) error {
	return u.repo.DeleteChatInDB(chatID)
}
