package usecase

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	auth "project/internal/auth/user"
	"project/internal/images"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
	"project/internal/pkg/validation"
	"project/internal/user"
)

type usecase struct {
	authRepo      auth.Repository
	userRepo      user.Repository
	imagesUsecase images.Usecase
}

func NewAuthUserUsecase(authRepo auth.Repository, userRepo user.Repository, imagesUsecase images.Usecase) auth.Usecase {
	return &usecase{authRepo: authRepo, userRepo: userRepo, imagesUsecase: imagesUsecase}
}

func (u usecase) Signup(ctx context.Context, registrationUser model.RegistrationUser) (model.User, error) {
	user := model.AuthorizedUser{
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
		Status:   "Привет, я использую технограм!",
	}

	err := u.authRepo.CheckExistEmail(context.Background(), user.Email)
	if err == nil {
		return model.User{}, myErrors.ErrEmailIsAlreadyRegistered
	}
	if !errors.Is(err, myErrors.ErrEmailNotFound) {
		return model.User{}, err
	}

	validationErrors := validation.ValidateUser(user)
	if len(validationErrors) != 0 {
		return model.User{}, validationErrors[0]
	}

	hashedPassword := security.Hash([]byte(registrationUser.Password))
	user.Password = hashedPassword

	sessionUser, err := u.authRepo.CreateUser(context.Background(), user)
	if err != nil {
		return model.User{}, err
	}

	//file, err := image_generation.GenerateAvatar(string(sessionUser.Nickname[0]))
	//if err != nil {
	//	log.Error(err)
	//}
	//
	err = u.imagesUsecase.UploadGeneratedImage(ctx, "brigade_user_avatars", string(sessionUser.Id), string(sessionUser.Nickname[0]))
	if err != nil {
		log.Error(err)
	}

	url, err := u.imagesUsecase.GetImage(ctx, "brigade_user_avatars", string(sessionUser.Id))
	if err != nil {
		log.Error(err)
	}
	sessionUser.Avatar = url

	return model_conversion.FromAuthorizedUserToUser(sessionUser), err
}

func (u usecase) Login(ctx context.Context, loginUser model.LoginUser) (model.User, error) {
	user := model.AuthorizedUser{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	err := u.authRepo.CheckExistEmail(context.Background(), user.Email)
	if err != nil {
		return model.User{}, err
	}

	hashedPassword := security.Hash([]byte(user.Password))
	user.Password = hashedPassword

	err = u.authRepo.CheckCorrectPassword(context.Background(), user.Email, user.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err = u.userRepo.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromAuthorizedUserToUser(user), err
}
