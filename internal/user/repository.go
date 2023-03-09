package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetUserById(ctx context.Context, userID uint64) (model.User, error)
}
