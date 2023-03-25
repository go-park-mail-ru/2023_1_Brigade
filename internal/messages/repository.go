package messages

import "project/internal/model"

type Repository interface {
	InsertMessageInDB(message model.Message) (model.Message, error)
	InsertMessageReceiveInDB(message model.ProducerMessage) error
	MarkMessageReading(messageID uint64) error
	GetChatById(chatID uint64) ([]model.ChatMembers, error)
}
