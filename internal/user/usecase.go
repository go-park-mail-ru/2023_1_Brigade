package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	GetUserById(ctx echo.Context, userID uint64) (model.User, error)
}
