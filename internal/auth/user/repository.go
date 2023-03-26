package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Repository interface {
	CreateUser(ctx echo.Context, user model.User) (model.User, error)
	CheckCorrectPassword(ctx echo.Context, user model.User) error
	CheckExistEmail(ctx echo.Context, email string) error
	CheckExistUsername(ctx echo.Context, username string) error
}
