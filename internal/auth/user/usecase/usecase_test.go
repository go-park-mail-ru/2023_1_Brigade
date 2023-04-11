package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	authUserMock "project/internal/auth/user/repository/mocks"
	"project/internal/configs"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

type testUserCase struct {
	expectedUser  model.AuthorizedUser
	expectedError error
	name          string
}

func Test_Signup_OK(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))

	user := model.RegistrationUser{
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Password: "password",
	}

	hashedUser := model.AuthorizedUser{
		Id:       0,
		Avatar:   configs.DefaultAvatarUrl,
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "Hello! I'm use technogramm",
		Password: hashedPassword,
	}

	test := testUserCase{
		expectedUser: model.AuthorizedUser{
			Id:       1,
			Avatar:   configs.DefaultAvatarUrl,
			Nickname: "marcussss",
			Email:    "marcussss@gmail.com",
			Status:   "Hello! I'm use technogramm",
			Password: hashedPassword,
		},
		expectedError: myErrors.ErrEmailNotFound,
		name:          "Successfull signup",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(context.Background(), user.Email).Return(test.expectedError).Times(1)
	authRepository.EXPECT().CreateUser(context.Background(), hashedUser).Return(test.expectedUser, nil).Times(1)

	myUser, err := usecase.Signup(ctx, user)

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserToUser(test.expectedUser), myUser, test.name)
}

func Test_Signup_UserIsAlreadyRegistred(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))

	user := model.RegistrationUser{
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Password: "password",
	}

	//hashedUser := model.User{
	//	Id:       0,
	//	Username: "id_marcussss",
	//	Nickname: "marcussss",
	//	Email:    "marcussss@gmail.com",
	//	Password: hashedPassword,
	//}

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
	usecase := NewAuthUserUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(context.Background(), user.Email).Return(nil).Times(1)
	_, err := usecase.Signup(ctx, user)
	require.Error(t, err, test.expectedError)
}

func Test_Login_OK(t *testing.T) {
	hashedPassword := security.Hash([]byte("password"))

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
	usecase := NewAuthUserUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(context.Background(), user.Email).Return(nil).Times(1)
	authRepository.EXPECT().CheckCorrectPassword(context.Background(), user.Email, hashedPassword).Return(nil).Times(1)
	userRepository.EXPECT().GetUserByEmail(context.Background(), user.Email).Return(test.expectedUser, test.expectedError).Times(1)

	myUser, err := usecase.Login(ctx, user)
	require.NoError(t, err)
	require.Equal(t, myUser, model_conversion.FromAuthorizedUserToUser(test.expectedUser), test.name)
}
