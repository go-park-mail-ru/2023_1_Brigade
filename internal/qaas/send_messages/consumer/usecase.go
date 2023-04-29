package consumer

import "context"

type Usecase interface {
	ConsumeMessage(ctx context.Context) []byte
	StartConsumeMessages(ctx context.Context)
}
