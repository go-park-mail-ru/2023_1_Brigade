package usecase

import "context"

type Usecase interface {
	StartConsumeMessages(ctx context.Context)
}
