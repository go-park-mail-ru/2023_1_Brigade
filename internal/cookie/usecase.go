package cookie

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	GetSessionById(ctx context.Context, userID uint64) (model.Session, error)
	GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error)
	CreateCookieById(ctx context.Context, userID uint64) (model.Session, error)
}
