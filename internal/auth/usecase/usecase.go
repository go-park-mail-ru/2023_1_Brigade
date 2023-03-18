package usecase

import (
	"errors"
	"github.com/labstack/echo/v4"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/pkg/security"
	"project/internal/user"

	"github.com/google/uuid"
)

type usecase struct {
	authRepo auth.Repository
	userRepo user.Repository
}

func NewAuthUsecase(authRepo auth.Repository, userRepo user.Repository) auth.Usecase {
	return &usecase{authRepo: authRepo, userRepo: userRepo}
}

func (u *usecase) Signup(ctx echo.Context, user model.User) (model.User, error) {
	exist, err := u.authRepo.CheckExistEmail(ctx, user.Email)
	if err != nil {
		if !errors.Is(err, myErrors.ErrUserNotFound) {
			return user, err
		}
	}
	if exist {
		return user, myErrors.ErrEmailIsAlreadyRegistred
	}

	exist, err = u.authRepo.CheckExistUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, myErrors.ErrUserNotFound) {
			return user, err
		}
	}
	if exist {
		return user, myErrors.ErrUsernameIsAlreadyRegistred
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

	validateError := security.ValidateSignup(user)
	if validateError != nil {
		return user, httpUtils.ErrorConversion(validateError[0])
	}

	userDB, err := u.authRepo.CreateUser(ctx, user)
	if err != nil {
		return userDB, err
	}

	return userDB, nil
}

func (u *usecase) Login(ctx echo.Context, user model.User) (model.User, error) {
	exist, err := u.authRepo.CheckExistEmail(ctx, user.Email)
	if !exist {
		return user, myErrors.ErrUserNotFound
	}
	if err != nil {
		return user, err
	}

	user.Password, err = security.Hash(user.Password)
	if err != nil {
		return user, err
	}

	isCorrectPassword, err := u.authRepo.CheckCorrectPassword(ctx, user)
	if !isCorrectPassword {
		return user, myErrors.ErrIncorrectPassword
	}
	if err != nil {
		return user, err
	}

	user, err = u.userRepo.GetUserByEmail(ctx, user.Email)
	return user, err
}

func (u *usecase) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	session, err := u.authRepo.GetSessionByCookie(ctx, cookie)
	return session, err
}

func (u *usecase) CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error) {
	session := model.Session{userID, uuid.New().String()}
	err := u.authRepo.CreateSession(ctx, session)
	return session, err
}

func (u *usecase) DeleteSessionByCookie(ctx echo.Context, cookie string) error {
	err := u.authRepo.DeleteSession(ctx, cookie)
	return err
}
