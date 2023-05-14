package usecase

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	ProduceMessage(ctx context.Context, message model.ProducerMessage) error
}
