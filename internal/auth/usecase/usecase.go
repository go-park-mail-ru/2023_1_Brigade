package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/pkg/security"
)

type usecase struct {
	repo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &usecase{repo: authRepo}
}

func (u *usecase) Signup(ctx context.Context, user model.User) (model.User, []error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		if !errors.Is(err, myErrors.ErrUserNotFound) {
			return userDB, []error{err}
		}
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return user, []error{err}
	}
	user.Password = hashedPassword

	validateErrors := security.ValidateSignup(user)
	if len(validateErrors) != 0 {
		return user, httpUtils.ErrorsConversion(validateErrors)
	}

	userDB, err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return user, []error{err}
	}

	userDB, err = u.repo.GetUserByEmail(ctx, user.Email) // для получения нормального айдишника
	if err != nil {
		if !errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			return user, []error{err}
		}
	}

	return userDB, nil
}

func (u *usecase) Login(ctx context.Context, user model.User) (model.User, error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		if errors.Is(err, myErrors.ErrUserNotFound) {
			return user, myErrors.ErrUserNotFound
		}
		if !errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			return user, err
		}
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return user, err
	}

	isCorrectPassword, err := u.repo.CheckCorrectPassword(ctx, hashedPassword)
	if err != nil {
		return user, err
	}

	if !isCorrectPassword {
		return user, myErrors.ErrIncorrectPassword
	}

	return userDB, nil
}

func (u *usecase) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)

	switch err {
	case myErrors.ErrSessionIsAlreadyCreated:
		return session, myErrors.ErrSessionIsAlreadyCreated
	case myErrors.ErrSessionNotFound:
		return session, myErrors.ErrSessionNotFound
	default:
		return session, err
	}
}

func (u *usecase) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	userDB, err := u.repo.GetUserById(ctx, userID)

	switch err {
	case myErrors.ErrUserIsAlreadyCreated:
		return userDB, myErrors.ErrUserIsAlreadyCreated
	case myErrors.ErrUserNotFound:
		return userDB, myErrors.ErrUserNotFound
	default:
		return userDB, err
	}
}

func (u *usecase) CreateSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	session := model.Session{userID, uuid.New().String()}
	session, err := u.repo.CreateSession(ctx, session)

	if err != nil {
		return session, err
	}

	return session, nil
}

func (u *usecase) DeleteSessionByCookie(ctx context.Context, cookie string) error {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)

	if err != nil {
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCreated) {
			return err
		}
	}

	err = u.repo.DeleteSession(ctx, session)
	if err != nil {
		return err
	}

	return nil
}
