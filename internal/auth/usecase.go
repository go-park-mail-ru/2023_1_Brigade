package auth

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, user model.User) (model.User, error)

	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	GetUserById(ctx context.Context, userID uint64) (model.User, error)
	CreateSessionById(ctx context.Context, userID uint64) (model.Session, error)
	DeleteSessionByCookie(ctx context.Context, cookie string) error
}
