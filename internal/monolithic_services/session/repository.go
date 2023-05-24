package session

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateSession(ctx context.Context, session model.Session) error
	DeleteSession(ctx context.Context, cookie string) error
}
