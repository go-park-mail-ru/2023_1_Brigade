package usecase

import (
	"project/internal/model"
	"project/internal/user"
)

type repositoryImpl struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &repositoryImpl{repo: userRepo}
}

func (u *repositoryImpl) GetUserById(userID int) (model.User, error) {
	user, err := u.repo.GetUserInDB(userID)
	return user, err
}

func (u *repositoryImpl) ChangeUserById(userID int, newDataUser []byte) (model.User, error) {
	user, err := u.repo.ChangeUserInDB(userID, newDataUser)
	return user, err
}

func (u *repositoryImpl) DeleteUserById(userID int) error {
	err := u.repo.DeleteUserInDB(userID)
	return err
}
