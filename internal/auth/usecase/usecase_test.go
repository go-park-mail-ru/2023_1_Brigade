package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"project/internal/auth/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
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

//func Test_GetUserById(t *testing.T) {
//	tests := []testUserCase{
//		{
//			expectedUser: model.User{
//				Id:       1,
//				Username: "marcussss",
//				Email:    "marcussss@gmail.com",
//				Status:   "cool",
//				Password: "password",
//			},
//			expectedError: myErrors.ErrUserIsAlreadyCreated,
//			name:          "Successfull getting user",
//		},
//		{
//			expectedUser: model.User{
//				Id:       1,
//				Username: "marcussss",
//				Email:    "marcussss@gmail.com",
//				Status:   "cool",
//				Password: "password",
//			},
//			expectedError: myErrors.ErrUserNotFound,
//			name:          "User not found",
//		},
//		{
//			expectedUser: model.User{
//				Id:       1,
//				Username: "marcussss",
//				Email:    "marcussss@gmail.com",
//				Status:   "cool",
//				Password: "password",
//			},
//			expectedError: myErrors.ErrInternal,
//			name:          "Internal error",
//		},
//	}
//
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	repository := mocks.NewMockRepository(ctl)
//	usecase := NewAuthUsecase(repository)
//	ctx := context.Background()
//
//	for _, test := range tests {
//		repository.EXPECT().GetUserById(ctx, 1).Return(test.expectedUser, test.expectedError).Times(1)
//		user, err := usecase.GetUserById(ctx, 1)
//
//		require.Error(t, err, test.expectedError)
//		require.Equal(t, user, test.expectedUser, test.name)
//	}
//}
