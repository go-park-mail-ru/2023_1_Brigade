package user

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	GetUserById(ctx context.Context, userID int) (model.User, error)
	ChangeUserById(ctx context.Context, userID int, data []byte) (model.User, error)
	DeleteUserById(ctx context.Context, userID int) error
}
