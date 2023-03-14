package usecase

import (
	"errors"
	"github.com/labstack/echo/v4"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"

	"github.com/google/uuid"
)

type usecase struct {
	repo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &usecase{repo: authRepo}
}

func (u *usecase) Signup(ctx echo.Context, user model.User) (model.User, error) {
	exist, err := u.repo.CheckExistEmail(ctx, user.Email)
	if err != nil {
		if !errors.Is(err, myErrors.ErrUserNotFound) {
			return user, err
		}
	}
	if exist {
		return user, myErrors.ErrEmailIsAlreadyRegistred
	}

	exist, err = u.repo.CheckExistUsername(ctx, user.Username)
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

	validateErrors := security.ValidateSignup(user)
	if len(validateErrors) != 0 {
		return user, validateErrors[0]
	}

	userDB, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return userDB, err
	}

	return userDB, nil
}

func (u *usecase) Login(ctx echo.Context, user model.User) (model.User, error) {
	exist, err := u.repo.CheckExistEmail(ctx, user.Email)
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

	isCorrectPassword, err := u.repo.CheckCorrectPassword(ctx, user)
	if !isCorrectPassword {
		return user, myErrors.ErrIncorrectPassword
	}
	if err != nil {
		return user, err
	}

	user, err = u.repo.GetUserByEmail(ctx, user.Email)
	return user, err
}

func (u *usecase) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)
	return session, err
}

func (u *usecase) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, userID)
	return user, err
}

func (u *usecase) CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error) {
	session := model.Session{userID, uuid.New().String()}
	err := u.repo.CreateSession(ctx, session)
	return session, err
}

func (u *usecase) DeleteSessionByCookie(ctx echo.Context, cookie string) error {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)
	if err != nil {
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCreated) {
			return err
		}
	}

	return u.repo.DeleteSession(ctx, session)
}
