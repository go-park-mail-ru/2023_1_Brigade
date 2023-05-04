package messages

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	DeleteMessageById(ctx context.Context, messageID string) error
	GetMessageById(ctx context.Context, messageID string) (model.Message, error)
	EditMessageById(ctx context.Context, producerMessage model.ProducerMessage) (model.Message, error)
	GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error)
	GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error)
	InsertMessageInDB(ctx context.Context, message model.Message) (model.Message, error)
	GetSearchMessages(ctx context.Context, userID uint64, string string) ([]model.Message, error)
}
