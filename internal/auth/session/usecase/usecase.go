package usecase

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"project/internal/auth/session"
	"project/internal/model"
)

type usecase struct {
	sessionRepo session.Repository
}

func NewAuthUserUsecase(sessionRepo session.Repository) session.Usecase {
	return &usecase{sessionRepo: sessionRepo}
}

func (u *usecase) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	session, err := u.sessionRepo.GetSessionByCookie(ctx, cookie)
	return session, err
}

func (u *usecase) CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error) {
	session := model.Session{UserId: userID, Cookie: uuid.New().String()}
	err := u.sessionRepo.CreateSession(ctx, session)
	return session, err
}

func (u *usecase) DeleteSessionByCookie(ctx echo.Context, cookie string) error {
	err := u.sessionRepo.DeleteSession(ctx, cookie)
	return err
}