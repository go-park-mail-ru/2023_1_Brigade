package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"project/internal/auth/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/security"
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

	repository := mocks.NewMockRepository(ctl)
	usecase := NewAuthUsecase(repository)
	ctx := context.Background()

	repository.EXPECT().GetUserByEmail(ctx, user.Email).Return(user, myErrors.ErrUserNotFound).Times(1)
	repository.EXPECT().GetUserByUsername(ctx, user.Username).Return(user, myErrors.ErrUserNotFound).Times(1)
	repository.EXPECT().CreateUser(ctx, hashedUser).Return(test.expectedUser, myErrors.ErrUserNotFound).Times(1)

	myUser, errors := usecase.Signup(ctx, user)

	require.Equal(t, len(errors), 1)
	require.Equal(t, errors[0], myErrors.ErrUserNotFound)
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

	repository := mocks.NewMockRepository(ctl)
	usecase := NewAuthUsecase(repository)
	ctx := context.Background()

	for i, test := range tests {
		var myUser model.User
		var errors []error

		if i == 0 {
			repository.EXPECT().GetUserByEmail(ctx, user.Email).Return(user, test.expectedError).Times(1)
			myUser, errors = usecase.Signup(ctx, user)
		} else {
			repository.EXPECT().GetUserByEmail(ctx, user.Email).Return(user, myErrors.ErrUserNotFound).Times(1)
			repository.EXPECT().GetUserByUsername(ctx, user.Username).Return(user, test.expectedError).Times(1)
			myUser, errors = usecase.Signup(ctx, user)
		}

		require.Equal(t, len(errors), 1)
		require.Error(t, errors[0], test.expectedError)
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

	repository := mocks.NewMockRepository(ctl)
	usecase := NewAuthUsecase(repository)
	ctx := context.Background()

	repository.EXPECT().GetUserByEmail(ctx, user.Email).Return(test.expectedUser, myErrors.ErrUserNotFound).Times(1)
	repository.EXPECT().CheckCorrectPassword(ctx, test.expectedUser.Password).Return(true, nil).Times(1)

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

	repository := mocks.NewMockRepository(ctl)
	usecase := NewAuthUsecase(repository)
	ctx := context.Background()

	for _, test := range tests {
		repository.EXPECT().GetSessionByCookie(ctx, "").Return(test.expectedSession, test.expectedError).Times(1)
		session, err := usecase.GetSessionByCookie(ctx, "")

		require.Error(t, err, test.expectedError)
		require.Equal(t, session, test.expectedSession, test.name)
	}
}

func Test_GetUserById(t *testing.T) {
	tests := []testUserCase{
		{
			expectedUser: model.User{
				Id:       1,
				Username: "marcussss",
				Email:    "marcussss@gmail.com",
				Status:   "cool",
				Password: "password",
			},
			expectedError: myErrors.ErrUserIsAlreadyCreated,
			name:          "Successfull getting user",
		},
		{
			expectedUser: model.User{
				Id:       1,
				Username: "marcussss",
				Email:    "marcussss@gmail.com",
				Status:   "cool",
				Password: "password",
			},
			expectedError: myErrors.ErrUserNotFound,
			name:          "User not found",
		},
		{
			expectedUser: model.User{
				Id:       1,
				Username: "marcussss",
				Email:    "marcussss@gmail.com",
				Status:   "cool",
				Password: "password",
			},
			expectedError: myErrors.ErrInternal,
			name:          "Internal error",
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repository := mocks.NewMockRepository(ctl)
	usecase := NewAuthUsecase(repository)
	ctx := context.Background()

	for _, test := range tests {
		repository.EXPECT().GetUserById(ctx, 1).Return(test.expectedUser, test.expectedError).Times(1)
		user, err := usecase.GetUserById(ctx, 1)

		require.Error(t, err, test.expectedError)
		require.Equal(t, user, test.expectedUser, test.name)
	}
}
