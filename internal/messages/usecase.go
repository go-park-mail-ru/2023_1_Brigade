package messages

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	SwitchMessageType(ctx context.Context, jsonWebSocketMessage []byte) error
	PutInProducer(ctx context.Context, webSocketMessage model.WebSocketMessage) error
	PullFromConsumer(ctx context.Context) ([]byte, error)
}
