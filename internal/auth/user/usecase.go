package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx echo.Context, user model.User) (model.User, error)
	Login(ctx echo.Context, user model.User) (model.User, error)
}
