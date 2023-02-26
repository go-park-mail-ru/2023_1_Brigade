package chat

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	InsertChatInDB(ctx context.Context, chat model.Chat) (model.Chat, error)
	GetChatInDB(ctx context.Context, chatID int) (model.Chat, error)
	GetAllChatsInDB(ctx context.Context) ([]model.Chat, error)
	DeleteChatInDB(ctx context.Context, chatID int) error
}
