package usecase

import (
	"errors"
	"github.com/labstack/echo/v4"
	auth "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/pkg/security"
	"project/internal/user"
)

type usecase struct {
	authRepo auth.Repository
	userRepo user.Repository
}

func NewAuthUserUsecase(authRepo auth.Repository, userRepo user.Repository) auth.Usecase {
	return &usecase{authRepo: authRepo, userRepo: userRepo}
}

func (u *usecase) Signup(ctx echo.Context, registrationUser model.RegistrationUser) (model.User, error) {
	user := model.User{
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
	}

	err := u.authRepo.CheckExistEmail(ctx, user.Email)
	if !errors.Is(err, myErrors.ErrUserNotFound) {
		return user, err
	}

	hashedPassword, err := security.Hash(registrationUser.Password)
	if err != nil {
		return user, err
	}
	registrationUser.Password = hashedPassword

	validateError := security.ValidateUser(user)
	if validateError != nil {
		return user, httpUtils.ErrorConversion(validateError[0])
	}

	userDB, err := u.authRepo.CreateUser(ctx, user)
	userDB.Username = "id" + string(userDB.Id)
	return userDB, err
}

func (u *usecase) Login(ctx echo.Context, loginUser model.LoginUser) (model.User, error) {
	user := model.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	err := u.authRepo.CheckExistEmail(ctx, user.Email)
	if err != nil {
		return user, err
	}

	user.Password, err = security.Hash(user.Password)
	if err != nil {
		return user, err
	}

	err = u.authRepo.CheckCorrectPassword(ctx, user)
	if err != nil {
		return user, err
	}

	user, err = u.userRepo.GetUserByEmail(ctx, user.Email)
	return user, err
}
