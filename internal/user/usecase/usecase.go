package usecase

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
	"project/internal/user"
)

type usecase struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &usecase{repo: userRepo}
}

func (u *usecase) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, userID)
	return user, err
}
