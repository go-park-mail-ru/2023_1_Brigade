package usecase

import (
	"context"
	"project/internal/cookie"
	"project/internal/pkg/http_utils"
)

type usecaseImpl struct {
	repo cookie.Repository
}

func NewCookieUsecase(cookieRepo cookie.Repository) cookie.Usecase {

}

func (u *usecaseImpl) CreateNewToken(ctx context.Context) string {

}

func (u *usecaseImpl) GetToken(ctx context.Context, token string) {

}