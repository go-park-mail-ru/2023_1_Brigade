package session

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Repository interface {
	GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error)
	CreateSession(ctx echo.Context, session model.Session) error
	DeleteSession(ctx echo.Context, cookie string) error
}
