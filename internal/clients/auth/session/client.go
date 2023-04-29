package auth

import (
	"context"
	authSession "project/internal/auth/session"
	"project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"

	"google.golang.org/grpc"
)

type authSessionServiceGRPCClient struct {
	authClient generated.AuthClient
}

func NewAuthUserServiceGRPSClient(con *grpc.ClientConn) authSession.Usecase {
	return &authSessionServiceGRPCClient{
		authClient: generated.NewAuthClient(con),
	}
}

func (as authSessionServiceGRPCClient) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	result, err := as.authClient.GetSessionByCookie(ctx, &generated.Cookie{
		Cookie: cookie,
	})
	return model_conversion.FromProtoSessionToSession(result), err
}

func (as authSessionServiceGRPCClient) CreateSessionById(ctx context.Context, userID uint64) (model.Session, error) {
	result, err := as.authClient.CreateSessionById(ctx, &generated.UserID{
		UserID: userID,
	})
	return model_conversion.FromProtoSessionToSession(result), err
}

func (as authSessionServiceGRPCClient) DeleteSessionByCookie(ctx context.Context, cookie string) error {
	_, err := as.authClient.DeleteSessionByCookie(ctx, &generated.Cookie{
		Cookie: cookie,
	})
	return err
}
