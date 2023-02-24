package chat

import "project/internal/model"

type Usecase interface {
	CreateChat(jsonChatData []byte) (model.Chat, error)
	GetChatById(chatID int) (model.Chat, error)
	GetAllChats() ([]model.Chat, error)
	DeleteChatById(chatID int) error
}
