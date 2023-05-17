package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	authUserMock "project/internal/microservices/auth/repository/mocks"
	chatMock "project/internal/microservices/chat/repository/mocks"
	userMock "project/internal/microservices/user/repository/mocks"
	"project/internal/model"
	imagesMock "project/internal/monolithic_services/images/usecase/mocks"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
	"strconv"
	"testing"
)

type testUserCase struct {
	expectedUser  model.AuthorizedUser
	expectedError error
	name          string
}

func Test_Signup_OK(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))
	filename := strconv.FormatUint(1, 10)
	url := "https://vk.com/avatars/1"
	ctx := context.TODO()

	user := model.RegistrationUser{
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Password: "password",
	}

	hashedUser := model.AuthorizedUser{
		Id:       0,
		Avatar:   "",
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "Привет, я использую технограм!",
		Password: hashedPassword,
	}

	test := testUserCase{
		expectedUser: model.AuthorizedUser{
			Id:       1,
			Avatar:   url,
			Nickname: "marcussss",
			Email:    "marcussss@gmail.com",
			Status:   "Привет, я использую технограм!",
			Password: hashedPassword,
		},
		expectedError: myErrors.ErrEmailNotFound,
		name:          "Successfull signup",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	chatRepository := chatMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository, chatRepository, imagesUsecase)

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(test.expectedError).Times(1)
	authRepository.EXPECT().CreateUser(ctx, hashedUser).Return(test.expectedUser, nil).Times(1)
	imagesUsecase.EXPECT().UploadGeneratedImage(ctx, config.UserAvatarsBucket, filename, string(test.expectedUser.Nickname[0])).
		Return(nil).Times(1)
	imagesUsecase.EXPECT().GetImage(ctx, config.UserAvatarsBucket, filename).Return(url, nil).Times(1)
	authRepository.EXPECT().UpdateUserAvatar(ctx, url, test.expectedUser.Id).Return(test.expectedUser, nil).Times(1)

	myUser, err := usecase.Signup(ctx, user)

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserToUser(test.expectedUser), myUser, test.name)
}

func Test_Signup_UserIsAlreadyRegistred(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))
	ctx := context.TODO()

	user := model.RegistrationUser{
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Password: "password",
	}

	test := testUserCase{
		expectedUser: model.AuthorizedUser{
			Id:       1,
			Username: "id_marcussss",
			Nickname: "marcussss",
			Email:    "marcussss@gmail.com",
			Password: hashedPassword,
		},
		expectedError: myErrors.ErrEmailIsAlreadyRegistered,
		name:          "Email is already created",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	chatRepository := chatMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository, chatRepository, imagesUsecase)

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(nil).Times(1)
	_, err := usecase.Signup(ctx, user)
	require.Error(t, err, test.expectedError)
}

func Test_Login_OK(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))
	ctx := context.TODO()

	user := model.LoginUser{
		Email:    "marcussss@gmail.com",
		Password: "password",
	}

	test := testUserCase{
		expectedUser: model.AuthorizedUser{
			Id:       1,
			Username: "marcussss",
			Email:    "marcussss@gmail.com",
			Status:   "cool",
			Password: hashedPassword,
		},
		expectedError: nil,
		name:          "User successfull login",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	chatRepository := chatMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository, chatRepository, imagesUsecase)

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(nil).Times(1)
	authRepository.EXPECT().CheckCorrectPassword(ctx, user.Email, hashedPassword).Return(nil).Times(1)
	userRepository.EXPECT().GetUserByEmail(ctx, user.Email).Return(test.expectedUser, test.expectedError).Times(1)

	myUser, err := usecase.Login(ctx, user)
	require.NoError(t, err)
	require.Equal(t, myUser, model_conversion.FromAuthorizedUserToUser(test.expectedUser), test.name)
}
