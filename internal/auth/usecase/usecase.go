package usecase

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	my_errors "project/internal/pkg/errors"
	"project/internal/pkg/security"
)

type usecaseImpl struct {
	repo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &usecaseImpl{repo: authRepo}
}

func (u *usecaseImpl) Signup(ctx context.Context, r *http.Request) ([]byte, []error) {
	user := model.User{}
	emptyResponse := []byte("")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Error(err)
		return emptyResponse, []error{my_errors.InternalError}
	}

	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, my_errors.EmailIsAlreadyRegistred) {
			log.Error(err)
			return emptyResponse, []error{my_errors.EmailIsAlreadyRegistred}
		}
		if !errors.Is(err, my_errors.NoUserFound) {
			log.Error(err)
			return emptyResponse, []error{my_errors.InternalError}
		}
	}

	userDB, err = u.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if errors.Is(err, my_errors.UsernameIsAlreadyRegistred) {
			log.Error(err)
			return emptyResponse, []error{my_errors.UsernameIsAlreadyRegistred}
		}
		if !errors.Is(err, my_errors.NoUserFound) {
			log.Error(err)
			return emptyResponse, []error{my_errors.InternalError}
		}
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		log.Error(err)
		return emptyResponse, []error{my_errors.InternalError}
	}
	user.Password = hashedPassword

	validateErrors := security.ValidateSignup(user)
	if len(validateErrors) != 0 {
		log.Error(validateErrors)
		return emptyResponse, validateErrors
	}

	userDB, err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Error(err)
		return emptyResponse, []error{my_errors.InternalError}
	}

	jsonUserDB, err := json.Marshal(userDB)
	if err != nil {
		log.Error(err)
		return emptyResponse, []error{my_errors.InternalError}
	}

	return jsonUserDB, nil
}

func (u *usecaseImpl) Login(ctx context.Context, r *http.Request) ([]byte, error) {
	user := model.User{}
	emptyResponse := []byte("")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Error(err)
		return emptyResponse, my_errors.InternalError
	}

	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)
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

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		log.Error(err)
		return emptyResponse, my_errors.InternalError
	}

	if userDB.Password != hashedPassword {
		log.Error(my_errors.IncorrectPassword)
		return emptyResponse, my_errors.IncorrectPassword
	}

	jsonUserDB, err := json.Marshal(userDB)
	if err != nil {
		log.Error(err)
		return emptyResponse, my_errors.InternalError
	}

	return jsonUserDB, nil
}
