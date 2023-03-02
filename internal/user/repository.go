package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetUserById(ctx context.Context, userID int) (model.User, error)
}
