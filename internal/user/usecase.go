package user

import (
	"context"
)

type Usecase interface {
	GetUserById(ctx context.Context, userID int) ([]byte, error)
}
