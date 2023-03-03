package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"
)

type usecase struct {
	repo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &usecase{repo: authRepo}
}

func (u *usecase) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	userDB, err := u.repo.GetUserByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			log.Error(err)
			return userDB, myErrors.ErrEmailIsAlreadyRegistred
		}
		if !errors.Is(err, myErrors.ErrNoUserFound) {
			log.Error(err)
			return userDB, myErrors.ErrInternal
		}
	}

	return userDB, myErrors.ErrNoUserFound
}

func (u *usecase) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	userDB, err := u.repo.GetUserByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, myErrors.ErrUsernameIsAlreadyRegistred) {
			log.Error(err)
			return userDB, myErrors.ErrUsernameIsAlreadyRegistred
		}
		if !errors.Is(err, myErrors.ErrNoUserFound) {
			log.Error(err)
			return userDB, myErrors.ErrInternal
		}
	}

	return userDB, myErrors.ErrNoUserFound
}

func (u *usecase) Signup(ctx context.Context, user model.User) (model.User, []error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		if !errors.Is(err, myErrors.ErrNoUserFound) {
			log.Error(err)
			return userDB, []error{err}
		}
	}

	userDB, err = u.repo.GetUserByUsername(ctx, user.Username)

	if err != nil {
		if !errors.Is(err, myErrors.ErrNoUserFound) {
			log.Error(err)
			return userDB, []error{err}
		}
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		log.Error(err)
		return user, []error{myErrors.ErrInternal}
	}
	user.Password = hashedPassword

	validateErrors := security.ValidateSignup(user)
	if len(validateErrors) != 0 {
		log.Error(validateErrors)
		return user, validateErrors
	}

	userDB, err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Error(err)
		return user, []error{myErrors.ErrInternal}
	}

	userDB, err = u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if !errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			log.Error(err)
			return user, []error{myErrors.ErrInternal}
		}
	}
	log.Println(userDB)
	return userDB, nil
}

func (u *usecase) Login(ctx context.Context, user model.User) (model.User, error) {
	userDB, err := u.repo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		if errors.Is(err, myErrors.ErrNoUserFound) {
			log.Error(err)
			return user, myErrors.ErrNoUserFound
		}
		if !errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred) {
			log.Error(err)
			return user, myErrors.ErrInternal
		}
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		log.Error(err)
		return user, myErrors.ErrInternal
	}

	isCorrectPassword, err := u.repo.CheckCorrectPassword(ctx, hashedPassword)
	if err != nil {
		log.Error(err)
		return user, myErrors.ErrInternal
	}

	if !isCorrectPassword {
		log.Error(myErrors.ErrIncorrectPassword)
		return user, myErrors.ErrIncorrectPassword
	}

	return userDB, nil
}

func (u *usecase) GetSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	session, err := u.repo.GetSessionById(ctx, userID)

	if err != nil {
		if errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			log.Error(err)
			return session, myErrors.ErrSessionIsAlreadyCrated
		}
		if !errors.Is(err, myErrors.ErrNoSessionFound) {
			log.Error(err)
			return session, myErrors.ErrInternal
		}
	}

	return session, myErrors.ErrNoSessionFound
}

func (u *usecase) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)

	if err != nil {
		if errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			log.Error(err)
			return session, myErrors.ErrSessionIsAlreadyCrated
		}
		if !errors.Is(err, myErrors.ErrNoSessionFound) {
			log.Error(err)
			return session, myErrors.ErrInternal
		}
	}

	return session, myErrors.ErrNoSessionFound
}

func (u *usecase) CreateSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	session, err := u.repo.GetSessionById(ctx, userID)

	if err != nil {
		if !errors.Is(err, myErrors.ErrNoSessionFound) {
			log.Error(err)
			return session, err
		}
	}

	session.UserId = userID
	session.Cookie = uuid.New().String()
	session, err = u.repo.CreateSession(ctx, session)

	if err != nil {
		log.Error(err)
		return session, myErrors.ErrInternal
	}

	return session, nil
}

func (u *usecase) DeleteSessionByCookie(ctx context.Context, cookie string) error {
	session, err := u.repo.GetSessionByCookie(ctx, cookie)

	if err != nil {
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			log.Error(err)
			return err
		}
	}

	err = u.repo.DeleteSession(ctx, session)
	if err != nil {
		log.Error(err)
		return myErrors.ErrInternal
	}

	return nil
}
