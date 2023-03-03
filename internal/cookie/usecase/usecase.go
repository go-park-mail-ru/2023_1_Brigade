package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"project/internal/cookie"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

type usecase struct {
	repo cookie.Repository
}

func NewUserUsecase(sessionRepo cookie.Repository) cookie.Usecase {
	return &usecase{repo: sessionRepo}
}

func (u *usecase) GetSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	session, err := u.repo.GetSessionById(ctx, userID)

	if err != nil {
		if errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			log.Error(err)
			return session, myErrors.ErrSessionIsAlreadyCrated
		}
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
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
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			log.Error(err)
			return session, myErrors.ErrInternal
		}
	}

	return session, myErrors.ErrNoSessionFound
}

func (u *usecase) CreateCookieById(ctx context.Context, userID uint64) (model.Session, error) {
	session, err := u.repo.GetSessionById(ctx, userID)

	if err != nil {
		if !errors.Is(err, myErrors.ErrSessionIsAlreadyCrated) {
			return session, err
		}
	}

	session.UserId = userID
	session.Cookie = uuid.New().String()
	session, err = u.repo.CreateCookie(ctx, session)

	if err != nil {
		log.Error(err)
		return session, myErrors.ErrInternal
	}

	return session, nil
}
