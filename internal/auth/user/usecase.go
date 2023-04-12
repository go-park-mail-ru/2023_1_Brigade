package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx echo.Context, registrationUser model.RegistrationUser) (model.User, error)
	Login(ctx echo.Context, loginUser model.LoginUser) (model.User, error)
}
