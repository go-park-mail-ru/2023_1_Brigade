package chat

import (
	"context"
	"project/internal/pkg/http_utils"
)

type Usecase interface {
	CreateChat(ctx context.Context, jsonChatData []byte) http_utils.Response
	GetChatById(ctx context.Context, chatID int) http_utils.Response
	GetAllChats(ctx context.Context) http_utils.Response
	DeleteChatById(ctx context.Context, chatID int) http_utils.Response
}
