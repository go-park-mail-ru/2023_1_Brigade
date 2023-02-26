package repository

import (
	"context"
	"fmt"
	"project/internal/model"
	"project/internal/user"
)

var users = [1]model.User{model.User{1, "marcussss", "danila", "123456", "88005553535"}}

func NewUserMemoryRepository() user.Repository {
	return &repositoryImpl{}
}

type repositoryImpl struct{}

func (r *repositoryImpl) GetUserInDB(ctx context.Context, userID int) (model.User, error) {
	fmt.Println("GET USER")
	return users[0], nil
}

func (r *repositoryImpl) ChangeUserInDB(ctx context.Context, userID int, data []byte) (model.User, error) {
	fmt.Println("PUT USER")
	return users[0], nil
}

func (r *repositoryImpl) DeleteUserInDB(ctx context.Context, userID int) error {
	fmt.Println("DELETE USER")
	return nil
}
