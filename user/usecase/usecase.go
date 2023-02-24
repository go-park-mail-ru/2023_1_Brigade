package usecase

import (
	"example.com/m/model"
	"example.com/m/user"
	"fmt"
)

type repositoryImpl struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &repositoryImpl{repo: userRepo}
}

func (u *repositoryImpl) GetUserById(userID int) (model.User, error) {
	fmt.Println("USECASE GET SUCCESS")
	u.repo.GetUserInDB(userID)
	return model.User{}, nil
}

func (u *repositoryImpl) EdidUserById(userID int, data []byte) (model.User, error) {
	fmt.Println("USECASE EDIT SUCCESS")
	u.repo.EdidUserInDB(userID, data)
	return model.User{}, nil
}

func (u *repositoryImpl) DeleteUserById(userID int) error {
	fmt.Println("USECASE DELETE SUCCESS")
	u.repo.DeleteUserInDB(userID)
	return nil
}
