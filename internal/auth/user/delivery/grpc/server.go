package grpc

import (
	"context"
	"net"
	authSession "project/internal/auth/session"
	authUser "project/internal/auth/user"
	"project/internal/generated"
	"project/internal/pkg/model_conversion"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

type authServiceGRPCServer struct {
	grpcServer         *grpc.Server
	authUserUsecase    authUser.Usecase
	authSessionUsecase authSession.Usecase
}

func NewAuthUserServiceGRPCServer(grpcServer *grpc.Server, userUsecase authUser.Usecase, sessionUsecase authSession.Usecase) *authServiceGRPCServer {
	return &authServiceGRPCServer{
		grpcServer:         grpcServer,
		authUserUsecase:    userUsecase,
		authSessionUsecase: sessionUsecase,
	}
}

func (a *authServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterAuthServer(a.grpcServer, a)

	return a.grpcServer.Serve(lis)
}

func (a *authServiceGRPCServer) Signup(ctx context.Context, registrationUser *generated.RegistrationUser) (*generated.User, error) {
	result, err := a.authUserUsecase.Signup(ctx, model_conversion.FromProtoRegistrationUserToRegistrationUser(registrationUser))
	return model_conversion.FromUserToProtoUser(result), err
}

func (a *authServiceGRPCServer) Login(ctx context.Context, loginUser *generated.LoginUser) (*generated.User, error) {
	result, err := a.authUserUsecase.Login(ctx, model_conversion.FromProtoLoginUserToLoginUser(loginUser))
	return model_conversion.FromUserToProtoUser(result), err
}

func (a *authServiceGRPCServer) GetSessionByCookie(ctx context.Context, cookie *generated.Cookie) (*generated.Session, error) {
	result, err := a.authSessionUsecase.GetSessionByCookie(ctx, model_conversion.FromProtoCookieToCookie(cookie))
	return model_conversion.FromSessionToProtoSession(result), err
}

func (a *authServiceGRPCServer) DeleteSessionByCookie(ctx context.Context, cookie *generated.Cookie) (*empty.Empty, error) {
	err := a.authSessionUsecase.DeleteSessionByCookie(ctx, model_conversion.FromProtoCookieToCookie(cookie))
	return nil, err
}

func (a *authServiceGRPCServer) CreateSessionById(ctx context.Context, userID *generated.UserID) (*generated.Session, error) {
	result, err := a.authSessionUsecase.CreateSessionById(ctx, model_conversion.FromProtoUserIDToUserID(userID))
	return model_conversion.FromSessionToProtoSession(result), err
}
