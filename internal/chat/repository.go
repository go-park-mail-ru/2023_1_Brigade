package chat

import (
	"context"
	"project/internal/model"
	"project/internal/pkg/http_utils"
)

type Repository interface {
	InsertChatInDB(ctx context.Context, chat model.Chat) http_utils.Response
	GetChatInDB(ctx context.Context, chatID int) http_utils.Response
	GetAllChatsInDB(ctx context.Context) http_utils.Response
	DeleteChatInDB(ctx context.Context, chatID int) http_utils.Response
}
