package auth

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	CheckCorrectPassword(ctx context.Context, hashedPassword string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)

	GetSessionById(ctx context.Context, userId uint64) (model.Session, error)
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateSession(ctx context.Context, session model.Session) (model.Session, error)
	DeleteSession(ctx context.Context, session model.Session) error
}
