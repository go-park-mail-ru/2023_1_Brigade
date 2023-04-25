package user

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	Signup(ctx context.Context, registrationUser model.RegistrationUser) (model.User, error)
	Login(ctx context.Context, loginUser model.LoginUser) (model.User, error)
}
