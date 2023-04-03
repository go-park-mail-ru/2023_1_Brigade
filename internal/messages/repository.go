package messages

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	InsertMessageInDB(ctx context.Context, message model.Message) (model.Message, error)
	InsertMessageReceiveInDB(ctx context.Context, message model.ProducerMessage) error
	MarkMessageReading(ctx context.Context, messageID uint64) error
	GetChatById(ctx context.Context, chatID uint64) ([]model.ChatMembers, error)
}
