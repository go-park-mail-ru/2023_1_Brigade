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
	"project/internal/pkg/security"
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

func Test_DeleteUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().DeleteUserById(context.TODO(), uint64(1)).Return(nil).Times(1)

	err := usecase.DeleteUserById(context.TODO(), uint64(1))

	require.NoError(t, err)
}

func Test_PutUserInfoById_OK(t *testing.T) {
	updateUser := model.UpdateUser{
		Nickname: "marcussss",
		Status:   "Hello world!",
	}

	updateAuthorizedUser := model.AuthorizedUser{
		Id:       1,
		Nickname: "marcussss",
		Status:   "Hello world!",
	}

	expectedUser := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Hello world!",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().UpdateUserEmailStatusById(context.TODO(), updateAuthorizedUser).Return(expectedUser, nil).Times(1)
	userRepository.EXPECT().UpdateUserAvatarNicknameById(context.TODO(), expectedUser).Return(expectedUser, nil).Times(1)

	user, err := usecase.PutUserById(context.TODO(), updateUser, uint64(1))

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserToUser(expectedUser), user)
}

func Test_PutUserPasswordById_OK(t *testing.T) {
	hashedOldPassword := security.Hash([]byte("12345678"))
	hashedNewPassword := security.Hash([]byte("87654321"))

	updateUser := model.UpdateUser{
		CurrentPassword: "12345678",
		NewPassword:     "87654321",
	}

	expectedUser := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Hello world!",
		Password: hashedOldPassword,
	}

	expectedNewUser := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Hello world!",
		Password: hashedNewPassword,
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(expectedUser, nil).Times(1)
	userRepository.EXPECT().UpdateUserPasswordById(context.TODO(), expectedNewUser).Return(expectedUser, nil).Times(1)

	user, err := usecase.PutUserById(context.TODO(), updateUser, uint64(1))

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserToUser(expectedUser), user)
}

func Test_GetUserContacts_OK(t *testing.T) {
	var expectedContacts []model.User

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().GetUserContacts(context.TODO(), uint64(1)).Return([]model.AuthorizedUser{}, nil).Times(1)

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

	userRepository.EXPECT().CheckExistUserById(context.TODO(), uint64(2)).Return(nil).Times(1)
	userRepository.EXPECT().CheckUserIsContact(context.TODO(), contact).Return(nil).Times(1)
	userRepository.EXPECT().AddUserInContact(context.TODO(), contact).Return(nil).Times(1)
	userRepository.EXPECT().GetUserContacts(context.TODO(), uint64(1)).Return(expectedContacts, nil).Times(1)

	contacts, err := usecase.AddUserContact(context.TODO(), uint64(1), uint64(2))

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserArrayToUserArray(expectedContacts), contacts)
}

func Test_CheckExistUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().CheckExistUserById(context.TODO(), uint64(1)).Return(nil).Times(1)

	err := usecase.CheckExistUserById(context.TODO(), uint64(1))

	require.NoError(t, err)
}

func Test_GetAllUsersExceptCurrentUser_OK(t *testing.T) {
	var expectedContacts []model.AuthorizedUser

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().GetAllUsersExceptCurrentUser(context.TODO(), uint64(1)).Return([]model.AuthorizedUser{}, nil).Times(1)

	contacts, err := usecase.GetAllUsersExceptCurrentUser(context.TODO(), uint64(1))

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserArrayToUserArray(expectedContacts), contacts)
}

func Test_GetSearchUsers_OK(t *testing.T) {
	searchString := "abc"
	var expectedContacts []model.AuthorizedUser

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	authRepository := authUserMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewUserUsecase(userRepository, authRepository)

	userRepository.EXPECT().GetSearchUsers(context.TODO(), searchString).Return([]model.AuthorizedUser{}, nil).Times(1)

	contacts, err := usecase.GetSearchUsers(context.TODO(), searchString)

	require.NoError(t, err)
	require.Equal(t, model_conversion.FromAuthorizedUserArrayToUserArray(expectedContacts), contacts)
}
