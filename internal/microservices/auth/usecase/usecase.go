package usecase

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"project/internal/config"
	"project/internal/microservices/auth"
	"project/internal/microservices/chat"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/images"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
	"project/internal/pkg/validation"
	"strconv"
)

type usecase struct {
	authRepo      auth.Repository
	userRepo      user.Repository
	chatRepo      chat.Repository
	imagesUsecase images.Usecase
}

func NewAuthUserUsecase(authRepo auth.Repository, userRepo user.Repository, chatRepo chat.Repository, imagesUsecase images.Usecase) auth.Usecase {
	return &usecase{authRepo: authRepo, userRepo: userRepo, chatRepo: chatRepo, imagesUsecase: imagesUsecase}
}

func (u usecase) Signup(ctx context.Context, registrationUser model.RegistrationUser) (model.User, error) {
	username := "id" + strconv.Itoa(rand.Int())
	user := model.AuthorizedUser{
		Username: username,
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
		Status:   "Привет, я использую технограм!",
	}

	err := u.authRepo.CheckExistEmail(ctx, user.Email)
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

	sessionUser, err := u.authRepo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	err = u.chatRepo.CreateTechnogrammChat(ctx, sessionUser)
	if err != nil {
		return model.User{}, err
	}

	filename := strconv.FormatUint(sessionUser.Id, 10)
	firstCharacterName := string(sessionUser.Nickname[0])

	err = u.imagesUsecase.UploadGeneratedImage(ctx, config.UserAvatarsBucket, filename, firstCharacterName)
	if err != nil {
		log.Error(err)
	}

	url, err := u.imagesUsecase.GetImage(ctx, config.UserAvatarsBucket, filename)
	if err != nil {
		log.Error(err)
	}

	sessionUser, err = u.authRepo.UpdateUserAvatar(ctx, url, sessionUser.Id)
	if err != nil {
		log.Error(err)
	}

	return model_conversion.FromAuthorizedUserToUser(sessionUser), err
}

func (u usecase) Login(ctx context.Context, loginUser model.LoginUser) (model.User, error) {
	user := model.AuthorizedUser{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	err := u.authRepo.CheckExistEmail(ctx, user.Email)
	if err != nil {
		return model.User{}, err
	}

	hashedPassword := security.Hash([]byte(user.Password))
	user.Password = hashedPassword

	err = u.authRepo.CheckCorrectPassword(ctx, user.Email, user.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err = u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromAuthorizedUserToUser(user), err
}
