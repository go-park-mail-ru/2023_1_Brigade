package chat

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	DeleteChatMembers(ctx context.Context, chatID uint64) error
	UpdateChatById(ctx context.Context, title string, chatID uint64) (model.DBChat, error)
	DeleteChatById(ctx context.Context, chatID uint64) error
	GetChatById(ctx context.Context, chatID uint64) (model.Chat, error)
	GetChatMembersByChatId(ctx context.Context, chatID uint64) ([]model.ChatMembers, error)
	CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error)
	GetChatsByUserId(ctx context.Context, userID uint64) ([]model.ChatMembers, error)
	AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error
}
