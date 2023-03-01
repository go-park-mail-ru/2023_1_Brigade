package auth

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	CheckExistUserByEmail(ctx context.Context, email string) bool
}
