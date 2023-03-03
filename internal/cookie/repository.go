package cookie

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	GetSessionById(ctx context.Context, userId uint64) (model.Session, error)
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateCookie(ctx context.Context, session model.Session) (model.Session, error)
}
