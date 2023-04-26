package messages

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	SwitchMesssageType(ctx context.Context, jsonWebSocketMessage []byte) error
	PutInProducer(ctx context.Context, producerMessage model.ProducerMessage) error
	PullFromConsumer(ctx context.Context) ([]byte, error)
}
