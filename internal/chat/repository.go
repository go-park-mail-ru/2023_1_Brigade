package chat

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetChatById(ctx context.Context, chatID uint64) (model.Chat, error)
	CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error)
	DeleteChatById(ctx context.Context, chatID uint64) error
	AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error
}
