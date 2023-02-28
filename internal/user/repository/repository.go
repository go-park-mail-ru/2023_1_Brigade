package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"project/internal/model"
	"project/internal/user"
)

func NewUserMemoryRepository(db *sqlx.DB) user.Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sqlx.DB
}

func (r *repositoryImpl) GetUserInDB(ctx context.Context, userID int) (model.User, error) {
	return model.User{}, nil
}

func (r *repositoryImpl) ChangeUserInDB(ctx context.Context, userID int, data []byte) (model.User, error) {
	fmt.Println("PUT USER")
	return model.User{}, nil
}

func (r *repositoryImpl) DeleteUserInDB(ctx context.Context, userID int) error {
	fmt.Println("DELETE USER")
	return nil
}
