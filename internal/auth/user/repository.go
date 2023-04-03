package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	CheckCorrectPassword(ctx context.Context, email string, password string) error
	CheckExistEmail(ctx context.Context, email string) error
	CheckExistUsername(ctx context.Context, username string) error
}
