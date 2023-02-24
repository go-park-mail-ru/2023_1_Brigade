package user

import "example.com/m/model"

type Repository interface {
	GetUserInDB(userID int) (model.User, error)
	EdidUserInDB(userID int, data []byte) (model.User, error)
	DeleteUserInDB(userID int) error
}
