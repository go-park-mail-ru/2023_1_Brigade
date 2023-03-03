package auth

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx context.Context, user model.User) (model.User, []error)
	Login(ctx context.Context, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)

	GetSessionById(ctx context.Context, userID uint64) (model.Session, error)
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateSessionById(ctx context.Context, userID uint64) (model.Session, error)
	DeleteSessionByCookie(ctx context.Context, cookie string) error
}
