package user

import "project/internal/model"

type Repository interface {
	GetUserInDB(userID int) (model.User, error)
	ChangeUserInDB(userID int, newDataUser []byte) (model.User, error)
	DeleteUserInDB(userID int) error
}
