package auth

import (
	"context"
	authUser "project/internal/auth/user"
	"project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"

	"google.golang.org/grpc"
)

type authUserServiceGRPCClient struct {
	authClient generated.AuthClient
}

func NewAuthUserServiceGRPSClient(con *grpc.ClientConn) authUser.Usecase {
	return &authUserServiceGRPCClient{
		authClient: generated.NewAuthClient(con),
	}
}

func (au authUserServiceGRPCClient) Login(ctx context.Context, user model.LoginUser) (model.User, error) {
	result, err := au.authClient.Login(ctx, &generated.LoginUser{
		Email:    user.Email,
		Password: user.Password,
	})
	return model_conversion.FromProtoUserToUser(result), err
}

func (au authUserServiceGRPCClient) Signup(ctx context.Context, user model.RegistrationUser) (model.User, error) {
	result, err := au.authClient.Signup(ctx, &generated.RegistrationUser{
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
	})
	return model_conversion.FromProtoUserToUser(result), err
}
