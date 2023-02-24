package user

import "project/internal/model"

type Usecase interface {
	GetUserById(userID int) (model.User, error)
	ChangeUserById(userID int, data []byte) (model.User, error)
	DeleteUserById(userID int) error
}
