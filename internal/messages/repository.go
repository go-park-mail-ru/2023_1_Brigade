package messages

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetMessageById(ctx context.Context, messageID string) (model.Message, error)
	GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error)
	GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error)
	InsertMessageInDB(ctx context.Context, message model.Message) (model.Message, error)
}
