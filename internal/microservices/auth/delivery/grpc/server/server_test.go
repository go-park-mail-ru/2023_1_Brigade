package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	authUserMock "project/internal/microservices/auth/usecase/mocks"
	"project/internal/model"
	authSessionMock "project/internal/monolithic_services/session/usecase/mocks"
	"project/internal/pkg/model_conversion"
	"testing"
)

func TestServer_Signup_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)

	authUserService := NewAuthUserServiceGRPCServer(grpcServer, authUserUsecase, authSessionUsecase)

	registrationUser := model.RegistrationUser{
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Password: "password",
	}

	expectedUser := model.User{
		Id:       1,
		Username: "",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Привет, я использую технограм!",
		Avatar:   "",
	}

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(expectedUser, nil).Times(1)

	user, err := authUserService.Signup(context.TODO(), model_conversion.FromRegistrationUserToProtoRegistrationUser(registrationUser))

	require.NoError(t, err)
	require.Equal(t, expectedUser, model_conversion.FromProtoUserToUser(user))
}

func TestServer_Login_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)

	authUserService := NewAuthUserServiceGRPCServer(grpcServer, authUserUsecase, authSessionUsecase)

	loginUser := model.LoginUser{
		Email:    "marcussss@mail.ru",
		Password: "password",
	}

	expectedUser := model.User{
		Id:       1,
		Username: "",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Привет, я использую технограм!",
		Avatar:   "",
	}

	authUserUsecase.EXPECT().Login(context.TODO(), loginUser).Return(expectedUser, nil).Times(1)

	user, err := authUserService.Login(context.TODO(), model_conversion.FromLoginUserToProtoLoginUser(loginUser))

	require.NoError(t, err)
	require.Equal(t, expectedUser, model_conversion.FromProtoUserToUser(user))
}
