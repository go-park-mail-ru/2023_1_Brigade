package chat

import "project/internal/model"

type ChatsMicroservice interface {
	GetChatById(chatID uint64) (model.Chat, error)
}
