package chat

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	DeleteChatById(ctx context.Context, chatID uint64) error
	GetChatById(ctx context.Context, chatID uint64) (model.Chat, error)
	GetMembersByChatId(ctx context.Context, chatID uint64) ([]model.User, error)
	CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error)
	GetChatsByUserId(ctx context.Context, userID uint64) ([]model.ChatMembers, error)
	AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error
}
