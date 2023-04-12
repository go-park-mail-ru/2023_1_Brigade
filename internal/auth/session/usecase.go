package session

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error)
	CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error)
	DeleteSessionByCookie(ctx echo.Context, cookie string) error
}
