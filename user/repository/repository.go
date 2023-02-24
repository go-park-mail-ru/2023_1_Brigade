package repository

import (
	"example.com/m/model"
	"example.com/m/user"
	"fmt"
)

func NewUserMemoryRepository() user.Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct{}

func (r *repositoryImpl) GetUserInDB(userID int) (model.User, error) {
	fmt.Println("REPOSITORY GET SUCCESS")
	return model.User{}, nil
}

func (r *repositoryImpl) EdidUserInDB(userID int, data []byte) (model.User, error) {
	fmt.Println("REPOSITORY EDIT SUCCESS")
	return model.User{}, nil
}

func (r *repositoryImpl) DeleteUserInDB(userID int) error {
	fmt.Println("REPOSITORY DELETE SUCCESS")
	return nil
}
