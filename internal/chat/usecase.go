package chat

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	CreateChat(ctx context.Context, jsonChatData []byte) (model.Chat, error)
	GetChatById(ctx context.Context, chatID int) (model.Chat, error)
	GetAllChats(ctx context.Context) ([]model.Chat, error)
	DeleteChatById(ctx context.Context, chatID int) error
}
