package usecase

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	"project/internal/pkg/http_utils"
	"project/internal/pkg/security"
)

type usecaseImpl struct {
	repo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &usecaseImpl{repo: authRepo}
}

func (u *usecaseImpl) Signup(ctx context.Context, r *http.Request) http_utils.Response {
	user := model.User{
		Id:       0,
		Username: r.FormValue("username"),
		Name:     r.FormValue("nickname"),
		Email:    r.FormValue("email"),
		Status:   "",
		Password: r.FormValue("password"),
	}
	exist := u.repo.CheckExistUserByEmail(ctx, user.Email)
	response := http_utils.Response{Status: http_utils.STATUS_INTERNAL_ERR}

	if exist {
		return response
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return response
	}
	user.Password = hashedPassword

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		return response
	}

	userDB, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return response
	}

	jsonUserDB, err := json.Marshal(userDB)
	response.Data = jsonUserDB
	if err != nil {
		return response
	}

	response.Status = http_utils.STATUS_CREATED
	return response
}
