package server

import (
	"context"
	"net"
	"project/internal/generated"
	authUser "project/internal/microservices/auth"
	authSession "project/internal/monolithic_services/session"
	"project/internal/pkg/model_conversion"

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
