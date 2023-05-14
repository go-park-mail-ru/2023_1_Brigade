package usecase

import (
	"context"
	"github.com/google/uuid"
	"project/internal/auth/session"
	"project/internal/model"
	"project/internal/pkg/security"
)

type usecase struct {
	sessionRepo session.Repository
}

func NewAuthUserUsecase(sessionRepo session.Repository) session.Usecase {
	return &usecase{sessionRepo: sessionRepo}
}

func (u usecase) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	session, err := u.sessionRepo.GetSessionByCookie(ctx, cookie)
	return session, err
}

func (u usecase) CreateSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	session := model.Session{UserId: userID, Cookie: security.Hash([]byte(uuid.New().String()))}
	err := u.sessionRepo.CreateSession(ctx, session)
	return session, err
}

func (u usecase) DeleteSessionByCookie(ctx context.Context, cookie string) error {
	err := u.sessionRepo.DeleteSession(ctx, cookie)
	return err
}
