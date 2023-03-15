package auth

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx echo.Context, user model.User) (model.User, error)
	Login(ctx echo.Context, user model.User) (model.User, error)

	GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error)
	GetUserById(ctx echo.Context, userID uint64) (model.User, error)
	CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error)
	DeleteSessionByCookie(ctx echo.Context, cookie string) error
}
