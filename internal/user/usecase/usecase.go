package usecase

import (
	"context"
	"project/internal/model"
	"project/internal/user"
)

type usecaseImpl struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &usecaseImpl{repo: userRepo}
}

func (u *usecaseImpl) GetUserById(ctx context.Context, userID int) (model.User, error) {
	user, err := u.repo.GetUserInDB(ctx, userID)
	return user, err
}

func (u *usecaseImpl) ChangeUserById(ctx context.Context, userID int, newDataUser []byte) (model.User, error) {
	user, err := u.repo.ChangeUserInDB(ctx, userID, newDataUser)
	return user, err
}

func (u *usecaseImpl) DeleteUserById(ctx context.Context, userID int) error {
	err := u.repo.DeleteUserInDB(ctx, userID)
	return err
}
