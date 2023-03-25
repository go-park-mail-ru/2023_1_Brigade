package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Repository interface {
	DeleteUserById(ctx echo.Context, userID uint64) error
	GetUserById(ctx echo.Context, userID uint64) (model.User, error)
	GetUserByEmail(ctx echo.Context, email string) (model.User, error)
	AddUserInContact(ctx echo.Context, contact model.UserContact) error
	UpdateUserById(ctx echo.Context, user model.User) (model.User, error)
	GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error)
	CheckUserIsContact(ctx echo.Context, contact model.UserContact) (bool, error)
}
