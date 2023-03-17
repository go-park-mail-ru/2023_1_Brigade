package auth

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx echo.Context, user model.User) (model.User, error)
	CheckCorrectPassword(ctx echo.Context, user model.User) (bool, error)
	CheckExistEmail(ctx echo.Context, email string) (bool, error)
	CheckExistUsername(ctx echo.Context, username string) (bool, error)

	GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error)
	CreateSession(ctx echo.Context, session model.Session) error
	DeleteSession(ctx echo.Context, cookie string) error
}
