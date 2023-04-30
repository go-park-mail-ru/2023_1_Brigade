package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	authSessionMock "project/internal/auth/session/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"testing"
)

type testSessionCase struct {
	expectedSession model.Session
	expectedError   error
	name            string
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

	authRepository := authSessionMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository)

	for _, test := range tests {
		authRepository.EXPECT().GetSessionByCookie(context.Background(), "").Return(test.expectedSession, test.expectedError).Times(1)
		session, err := usecase.GetSessionByCookie(context.TODO(), "")

		require.Error(t, err, test.expectedError)
		require.Equal(t, session, test.expectedSession, test.name)
	}
}

func Test_DeletSessionByCookie_OK(t *testing.T) {
	test := testSessionCase{
		expectedSession: model.Session{
			UserId: 1,
			Cookie: uuid.New().String(),
		},
		expectedError: nil,
		name:          "Successfull getting session",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authSessionMock.NewMockRepository(ctl)
	usecase := NewAuthUserUsecase(authRepository)

	authRepository.EXPECT().DeleteSession(context.TODO(), test.expectedSession.Cookie).Return(nil).Times(1)

	err := usecase.DeleteSessionByCookie(context.TODO(), test.expectedSession.Cookie)

	require.NoError(t, test.expectedError, err)
}
