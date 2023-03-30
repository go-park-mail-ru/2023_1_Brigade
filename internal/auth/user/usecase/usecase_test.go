package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	authUserMock "project/internal/auth/user/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

type testUserCase struct {
	expectedUser  model.User
	expectedError error
	name          string
}

func Test_Signup_OK(t *testing.T) {
	hashedPassword, err := security.Hash("password")
	require.NoError(t, err)

	user := model.User{
		Id:       0,
		Username: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "cool",
		Password: "password",
	}

	hashedUser := model.User{
		Id:       0,
		Username: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "cool",
		Password: hashedPassword,
	}

	test := testUserCase{
		expectedUser: model.User{
			Id:       1,
			Username: "marcussss",
			Email:    "marcussss@gmail.com",
			Status:   "cool",
			Password: hashedPassword,
		},
		expectedError: nil,
		name:          "Successfull signup",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(myErrors.ErrUserNotFound).Times(1)
	authRepository.EXPECT().CheckExistUsername(ctx, user.Username).Return(myErrors.ErrUserNotFound).Times(1)
	authRepository.EXPECT().CreateUser(ctx, hashedUser).Return(test.expectedUser, nil).Times(1)

	myUser, err := usecase.Signup(ctx, user)

	require.NoError(t, err)
	require.Equal(t, test.expectedUser, myUser, test.name)
}

func Test_Signup_UserIsAlreadyRegistred(t *testing.T) {
	user := model.User{
		Id:       0,
		Username: "",
		Email:    "marcussss@gmail.com",
		Status:   "",
		Password: "password",
	}

	tests := []testUserCase{
		{
			expectedUser: model.User{
				Id:       0,
				Username: "",
				Email:    "marcussss@gmail.com",
				Status:   "",
				Password: "password",
			},
			expectedError: myErrors.ErrEmailIsAlreadyRegistred,
			name:          "Email is already created",
		},
		{
			expectedUser: model.User{
				Id:       0,
				Username: "",
				Email:    "marcussss@gmail.com",
				Status:   "",
				Password: "password",
			},
			expectedError: myErrors.ErrUsernameIsAlreadyRegistred,
			name:          "Username is already created",
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository, userRepository)
	var ctx echo.Context

	for i, test := range tests {
		var err error

		if i == 0 {
			authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(test.expectedError).Times(1)
			_, err = usecase.Signup(ctx, user)
		} else {
			authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(myErrors.ErrUserNotFound).Times(1)
			authRepository.EXPECT().CheckExistUsername(ctx, user.Username).Return(test.expectedError).Times(1)
			_, err = usecase.Signup(ctx, user)
		}

		require.Error(t, err, test.expectedError)
	}
}

func Test_Login_OK(t *testing.T) {
	hashedPassword, err := security.Hash("password")
	require.NoError(t, err)

	user := model.User{
		Id:       0,
		Username: "",
		Email:    "marcussss@gmail.com",
		Status:   "",
		Password: "password",
	}

	hashedPasswordUser := model.User{
		Id:       0,
		Username: "",
		Email:    "marcussss@gmail.com",
		Status:   "",
		Password: hashedPassword,
	}

	test := testUserCase{
		expectedUser: model.User{
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

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(nil).Times(1)
	authRepository.EXPECT().CheckCorrectPassword(ctx, hashedPasswordUser).Return(nil).Times(1)
	userRepository.EXPECT().GetUserByEmail(ctx, user.Email).Return(test.expectedUser, test.expectedError).Times(1)

	myUser, err := usecase.Login(ctx, user)
	require.NoError(t, err)
	require.Equal(t, myUser, test.expectedUser, test.name)
}
