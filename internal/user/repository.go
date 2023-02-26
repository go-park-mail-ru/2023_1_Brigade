package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetUserInDB(ctx context.Context, userID int) (model.User, error)
	ChangeUserInDB(ctx context.Context, userID int, newDataUser []byte) (model.User, error)
	DeleteUserInDB(ctx context.Context, userID int) error
}
