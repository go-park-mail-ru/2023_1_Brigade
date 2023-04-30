package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	protobuf "project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	mockUser "project/internal/user/usecase/mocks"
	"testing"
)

func TestServer_DeleteUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().DeleteUserById(context.TODO(), userID).Return(nil).Times(1)

	_, err := userService.DeleteUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
}

func TestServer_CheckExistUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().CheckExistUserById(context.TODO(), userID).Return(nil).Times(1)

	_, err := userService.CheckExistUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
}

func TestServer_GetUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)
	expectedUser := model.User{
		Id:       1,
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Привет, я использую технограм!",
	}

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().GetUserById(context.TODO(), userID).Return(expectedUser, nil).Times(1)

	user, err := userService.GetUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
	require.Equal(t, expectedUser, model_conversion.FromProtoUserToUser(user))
}

func TestServer_AddUserContact_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)
	contactID := uint64(2)
	var expectedContacts []model.User
	expectedContacts = append(expectedContacts, model.User{
		Id:       2,
		Username: "marcussss2",
		Nickname: "marcussss2",
		Email:    "marcussss2@mail.ru",
		Status:   "Привет, я использую технограм!",
	})

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().AddUserContact(context.TODO(), userID, contactID).Return(expectedContacts, nil).Times(1)

	contacts, err := userService.AddUserContact(context.TODO(), &protobuf.AddUserContactArguments{
		ContactID: contactID,
		UserID:    userID,
	})

	require.NoError(t, err)
	require.Equal(t, expectedContacts, []model.User{
		{
			Id:       contacts.Contacts[0].Id,
			Username: contacts.Contacts[0].Username,
			Nickname: contacts.Contacts[0].Nickname,
			Email:    contacts.Contacts[0].Email,
			Status:   contacts.Contacts[0].Status,
		},
	})
}

func TestServer_GetUserContacts_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)
	var expectedContacts []model.User
	expectedContacts = append(expectedContacts, model.User{
		Id:       2,
		Username: "marcussss2",
		Nickname: "marcussss2",
		Email:    "marcussss2@mail.ru",
		Status:   "Привет, я использую технограм!",
	})

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().GetUserContacts(context.TODO(), userID).Return(expectedContacts, nil).Times(1)

	contacts, err := userService.GetUserContacts(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
	require.Equal(t, expectedContacts, []model.User{
		{
			Id:       contacts.Contacts[0].Id,
			Username: contacts.Contacts[0].Username,
			Nickname: contacts.Contacts[0].Nickname,
			Email:    contacts.Contacts[0].Email,
			Status:   contacts.Contacts[0].Status,
		},
	})
}

func TestServer_PutUserById_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)
	updateUser := model.UpdateUser{
		Username:        "marcussss2",
		Nickname:        "marcussss2",
		Status:          "Привет, я использую технограм! Это круто",
		CurrentPassword: "password",
		NewPassword:     "password",
	}
	expectedUser := model.User{
		Id:       1,
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Привет, я использую технограм!",
	}

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().PutUserById(context.TODO(), updateUser, userID).Return(expectedUser, nil).Times(1)

	user, err := userService.PutUserById(context.TODO(), &protobuf.PutUserArguments{
		UserID: userID,
		User:   model_conversion.FromUpdateUserToProtoUpdateUser(updateUser),
	})

	require.NoError(t, err)
	require.Equal(t, expectedUser, model_conversion.FromProtoUserToUser(user))
}

func TestServer_GetAllUsersExceptCurrentUser_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := uint64(1)
	var expectedContacts []model.User
	expectedContacts = append(expectedContacts, model.User{
		Id:       2,
		Username: "marcussss2",
		Nickname: "marcussss2",
		Email:    "marcussss2@mail.ru",
		Status:   "Привет, я использую технограм!",
	})

	grpcServer := grpc.NewServer()

	userUsecase := mockUser.NewMockUsecase(ctl)

	userService := NewUsersServiceGRPCServer(grpcServer, userUsecase)

	userUsecase.EXPECT().GetAllUsersExceptCurrentUser(context.TODO(), userID).Return(expectedContacts, nil).Times(1)

	contacts, err := userService.GetAllUsersExceptCurrentUser(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
	require.Equal(t, expectedContacts, []model.User{
		{
			Id:       contacts.Contacts[0].Id,
			Username: contacts.Contacts[0].Username,
			Nickname: contacts.Contacts[0].Nickname,
			Email:    contacts.Contacts[0].Email,
			Status:   contacts.Contacts[0].Status,
		},
	})
}
