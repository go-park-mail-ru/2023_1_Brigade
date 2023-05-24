package session

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateSessionById(ctx context.Context, userID uint64) (model.Session, error)
	DeleteSessionByCookie(ctx context.Context, cookie string) error
}
