package usecase

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	my_errors "project/internal/pkg/errors"
	"project/internal/user"
)

type usecaseImpl struct {
	repo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &usecaseImpl{repo: userRepo}
}

func (u *usecaseImpl) GetUserById(ctx context.Context, userID int) ([]byte, error) {
	user, err := u.repo.GetUserById(ctx, userID)
	emptyResponse := []byte("")

	if err != nil {
		if errors.Is(err, my_errors.NoUserFound) {
			log.Error(err)
			return emptyResponse, my_errors.NoUserFound
		}
		if !errors.Is(err, my_errors.EmailIsAlreadyRegistred) {
			log.Error(err)
			return emptyResponse, my_errors.InternalError
		}
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Error(err)
		return emptyResponse, my_errors.InternalError
	}

	return jsonUser, nil
}
