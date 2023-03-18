package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	authMock "project/internal/auth/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

type testSessionCase struct {
	expectedSession model.Session
	expectedError   error
	name            string
}

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
			Id:       0,
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

	authRepository := authMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(false, nil).Times(1)
	authRepository.EXPECT().CheckExistUsername(ctx, user.Username).Return(false, nil).Times(1)
	authRepository.EXPECT().CreateUser(ctx, hashedUser).Return(test.expectedUser, myErrors.ErrUserNotFound).Times(1)

	myUser, err := usecase.Signup(ctx, user)

	require.Equal(t, err, myErrors.ErrUserNotFound)
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

	authRepository := authMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUsecase(authRepository, userRepository)
	var ctx echo.Context

	for i, test := range tests {
		var myUser model.User
		var err error

		if i == 0 {
			authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(true, test.expectedError).Times(1)
			myUser, err = usecase.Signup(ctx, user)
		} else {
			authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(false, nil).Times(1)
			authRepository.EXPECT().CheckExistUsername(ctx, user.Username).Return(true, test.expectedError).Times(1)
			myUser, err = usecase.Signup(ctx, user)
		}

		require.Error(t, err, test.expectedError)
		require.Equal(t, myUser, test.expectedUser, test.name)
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

	authRepository := authMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUsecase(authRepository, userRepository)
	var ctx echo.Context

	authRepository.EXPECT().CheckExistEmail(ctx, user.Email).Return(true, nil).Times(1)
	authRepository.EXPECT().CheckCorrectPassword(ctx, hashedPasswordUser).Return(true, nil).Times(1)
	userRepository.EXPECT().GetUserByEmail(ctx, user.Email).Return(test.expectedUser, test.expectedError).Times(1)
	//userRepository.EXPECT().GetUserByEmail(ctx, user.Email).Return(test.expectedUser, nil).Times(1)
	//userRepository.EXPECT().GetUserById(ctx, test.expectedUser.Id).Return(test.expectedUser, nil).Times(1)

	myUser, err := usecase.Login(ctx, user)
	require.NoError(t, err)
	require.Equal(t, myUser, test.expectedUser, test.name)
}

func Test_GetSessionByCookie(t *testing.T) {
	tests := []testSessionCase{
		{
			expectedSession: model.Session{
				UserId: 1,
				Cookie: uuid.New().String(),
			},
			expectedError: myErrors.ErrSessionIsAlreadyCreated,
			name:          "Successfull getting session",
		},
		{
			expectedSession: model.Session{
				UserId: 1,
				Cookie: uuid.New().String(),
			},
			expectedError: myErrors.ErrSessionNotFound,
			name:          "Session not found",
		},
		{
			expectedSession: model.Session{
				UserId: 1,
				Cookie: uuid.New().String(),
			},
			expectedError: myErrors.ErrInternal,
			name:          "Internal error",
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewAuthUsecase(authRepository, userRepository)
	var ctx echo.Context

	for _, test := range tests {
		authRepository.EXPECT().GetSessionByCookie(ctx, "").Return(test.expectedSession, test.expectedError).Times(1)
		session, err := usecase.GetSessionByCookie(ctx, "")

		require.Error(t, err, test.expectedError)
		require.Equal(t, session, test.expectedSession, test.name)
	}
}
