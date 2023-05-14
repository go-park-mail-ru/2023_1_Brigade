package messages

import (
	"context"
)

type Usecase interface {
	PutInProducer(ctx context.Context, jsonWebSocketMessage []byte) error
}
