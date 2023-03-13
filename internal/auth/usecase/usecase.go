package usecase

import (
	"context"
	"errors"
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

func (u *usecase) Signup(ctx context.Context, user model.User) (model.User, error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return userDB, myErrors.ErrEmailIsAlreadyRegistred
	}
	if !errors.Is(err, myErrors.ErrUserNotFound) {
		return userDB, err
	}

	userDB, err = u.repo.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return userDB, myErrors.ErrUsernameIsAlreadyRegistred
	}
	if !errors.Is(err, myErrors.ErrUserNotFound) {
		return userDB, err
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

	userDB, err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return user, err
	}

	return userDB, nil
}

func (u *usecase) Login(ctx context.Context, user model.User) (model.User, error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return userDB, err
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
	return session, err
}

func (u *usecase) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	userDB, err := u.repo.GetUserById(ctx, userID)
	return userDB, err
}

func (u *usecase) CreateSessionById(ctx context.Context, userID uint64) (session model.Session, err error) {
	session = model.Session{userID, uuid.New().String()}
	err = u.repo.CreateSession(ctx, session)
	return
}

func (u *usecase) DeleteSessionByCookie(ctx context.Context, cookie string) error {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)

	if err != nil {
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCreated) {
			return err
		}
	}

	return u.repo.DeleteSession(ctx, session)
}
