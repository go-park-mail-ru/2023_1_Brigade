package user

import "example.com/m/model"

type Usecase interface {
	GetUserById(userID int) (model.User, error)
	EdidUserById(userID int, data []byte) (model.User, error)
	DeleteUserById(userID int) error
}
