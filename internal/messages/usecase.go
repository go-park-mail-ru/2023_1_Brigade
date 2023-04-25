package messages

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	SwitchMesssageType(ctx context.Context, jsonWebSocketMessage []byte) error
	SendMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error
	EditMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error
	DeleteMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error
	ReceiveMessage(ctx context.Context) ([]byte, error)
}
