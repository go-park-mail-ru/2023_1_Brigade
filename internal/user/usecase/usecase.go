package usecase

import (
	"context"
	"errors"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/user"

	log "github.com/sirupsen/logrus"
)

type usecase struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &usecase{repo: userRepo}
}

func (u *usecase) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, userID)

	if err != nil {
		if errors.Is(err, myErrors.ErrUserNotFound) {
			log.Error(err)
			return user, myErrors.ErrUserNotFound
		}
		if !errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			log.Error(err)
			return user, myErrors.ErrInternal
		}
	}

	return user, nil
}
