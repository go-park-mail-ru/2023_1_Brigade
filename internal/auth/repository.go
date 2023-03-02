package auth

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	CheckCorrectPassword(ctx context.Context, hashedPassword string) (bool, error)
}
