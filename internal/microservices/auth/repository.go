package auth

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	UpdateUserAvatar(ctx context.Context, url string, userID uint64) (model.AuthorizedUser, error)
	CreateUser(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error)
	CheckCorrectPassword(ctx context.Context, email string, password string) error
	CheckExistEmail(ctx context.Context, email string) error
	CheckExistUsername(ctx context.Context, username string) error
}
