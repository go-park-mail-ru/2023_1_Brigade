package usecase

import "context"

type Usecase interface {
	ProduceMessage(ctx context.Context, message []byte) error
}
