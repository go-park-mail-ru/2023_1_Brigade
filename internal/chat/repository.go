package chat

import "project/internal/model"

type Repository interface {
	InsertChatInDB(chat model.Chat) (model.Chat, error)
	GetChatInDB(chatID int) (model.Chat, error)
	GetAllChatsInDB() ([]model.Chat, error)
	DeleteChatInDB(chatID int) error
}
