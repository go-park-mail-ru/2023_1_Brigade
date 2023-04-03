package usecase

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	auth "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"
	"project/internal/pkg/validation"
	"project/internal/user"
)

type usecase struct {
	authRepo auth.Repository
	userRepo user.Repository
}

func NewAuthUserUsecase(authRepo auth.Repository, userRepo user.Repository) auth.Usecase {
	return &usecase{authRepo: authRepo, userRepo: userRepo}
}

func (u usecase) Signup(ctx echo.Context, registrationUser model.RegistrationUser) (model.User, error) {
	user := model.User{
		Username: "id" + "_" + registrationUser.Nickname,
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
		Status:   "",
	}

	err := u.authRepo.CheckExistEmail(context.Background(), user.Email)
	if err == nil {
		return model.User{}, myErrors.ErrEmailIsAlreadyRegistered
	}
	if !errors.Is(err, myErrors.ErrEmailNotFound) {
		return model.User{}, err
	}

	validationErrors := validation.ValidateUser(user)
	if len(validationErrors) != 0 {
		return model.User{}, validationErrors[0]
	}

	hashedPassword := security.Hash([]byte(registrationUser.Password))
	user.Password = hashedPassword

	userFromDB, err := u.authRepo.CreateUser(context.Background(), user)
	return userFromDB, err
}

func (u usecase) Login(ctx echo.Context, loginUser model.LoginUser) (model.User, error) {
	user := model.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	err := u.authRepo.CheckExistEmail(context.Background(), user.Email)
	if err != nil {
		return model.User{}, err
	}

	hashedPassword := security.Hash([]byte(user.Password))
	user.Password = hashedPassword

	err = u.authRepo.CheckCorrectPassword(context.Background(), user.Email, user.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err = u.userRepo.GetUserByEmail(context.Background(), user.Email)
	return user, err
}
