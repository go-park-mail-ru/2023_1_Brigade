package chat

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	GetChatById(ctx context.Context, chatID uint64) (model.Chat, error)
	EditChat(ctx context.Context, editChat model.EditChat) (model.Chat, error)
	CreateChat(ctx context.Context, chat model.CreateChat, userID uint64) (model.Chat, error)
	DeleteChatById(ctx context.Context, chatID uint64) error
	CheckExistUserInChat(ctx context.Context, chat model.Chat, userID uint64) error
	GetListUserChats(ctx context.Context, userID uint64) ([]model.ChatInListUser, error)
}
