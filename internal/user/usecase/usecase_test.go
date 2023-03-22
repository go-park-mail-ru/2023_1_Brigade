package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

type testUserCase struct {
	expectedUser  model.User
	expectedError error
	name          string
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

	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository)
	var ctx echo.Context

	for _, test := range tests {
		userRepository.EXPECT().GetUserById(ctx, 1).Return(test.expectedUser, test.expectedError).Times(1)
		user, err := usecase.GetUserById(ctx, 1)

		require.Error(t, err, test.expectedError)
		require.Equal(t, user, test.expectedUser, test.name)
	}
}
