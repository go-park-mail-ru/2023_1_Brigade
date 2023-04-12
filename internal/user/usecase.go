package user

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
)

type Usecase interface {
	DeleteUserById(ctx echo.Context, userID uint64) error
	CheckExistUserById(ctx echo.Context, userID uint64) error
	GetUserById(ctx echo.Context, userID uint64) (model.User, error)
	AddUserContact(ctx echo.Context, userID uint64, contactID uint64) ([]model.User, error)
	GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error)
	PutUserById(ctx echo.Context, user model.UpdateUser, userID uint64) (model.User, error)
	GetAllUsersExceptCurrentUser(ctx echo.Context, userID uint64) ([]model.User, error)
}
