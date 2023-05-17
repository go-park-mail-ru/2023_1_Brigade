package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	authUserMock "project/internal/microservices/auth/repository/mocks"
	userMock "project/internal/microservices/user/repository/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"testing"
)

type testUserCase struct {
	expectedUser  model.AuthorizedUser
	expectedError error
	name          string
}

func Test_GetUserById(t *testing.T) {
	tests := []testUserCase{
		{
			expectedUser:  model.AuthorizedUser{},
			expectedError: myErrors.ErrUserIsAlreadyCreated,
			name:          "Successfull getting user",
		},
		{
			expectedUser:  model.AuthorizedUser{},
			expectedError: myErrors.ErrUserNotFound,
			name:          "User not found",
		},
		{
			expectedUser:  model.AuthorizedUser{},
			expectedError: myErrors.ErrInternal,
			name:          "Internal error",
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	for _, test := range tests {
		userRepository.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(test.expectedUser, test.expectedError).Times(1)
		user, err := usecase.GetUserById(context.TODO(), uint64(1))

		require.Error(t, err, test.expectedError)
		require.Equal(t, model_conversion.FromAuthorizedUserToUser(test.expectedUser), user, test.name)
	}
}

func Test_GetUserContacts_OK(t *testing.T) {
	var expectedContacts []model.User

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().GetUserContacts(context.Background(), uint64(1)).Return([]model.AuthorizedUser{}, nil).Times(1)

	contacts, err := usecase.GetUserContacts(context.TODO(), uint64(1))

	require.NoError(t, err)
	require.Equal(t, expectedContacts, contacts)
}

func Test_AddUserInContacts_OK(t *testing.T) {
	expectedContacts := []model.AuthorizedUser{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
	}
	contact := model.UserContact{
		IdUser:    uint64(1),
		IdContact: uint64(2),
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().CheckExistUserById(context.Background(), uint64(2)).Return(nil).Times(1)
	userRepository.EXPECT().CheckUserIsContact(context.Background(), contact).Return(nil).Times(1)
	userRepository.EXPECT().AddUserInContact(context.Background(), contact).Return(nil).Times(1)
	userRepository.EXPECT().GetUserContacts(context.Background(), uint64(1)).Return(expectedContacts, nil).Times(1)

	contacts, err := usecase.AddUserContact(context.TODO(), uint64(1), uint64(2))

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserArrayToUserArray(expectedContacts), contacts)
}
