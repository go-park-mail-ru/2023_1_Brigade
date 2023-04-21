package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"net"
	"project/internal/generated"
	"project/internal/pkg/model_conversion"
	"project/internal/user"
)

type usersServiceGRPCServer struct {
	grpcServer  *grpc.Server
	userUsecase user.Usecase
}

func NewUsersServiceGRPCServer(grpcServer *grpc.Server, userUsecase user.Usecase) *usersServiceGRPCServer {
	return &usersServiceGRPCServer{
		grpcServer:  grpcServer,
		userUsecase: userUsecase,
	}
}

func (c *usersServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterUsersServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *usersServiceGRPCServer) DeleteUserById(ctx context.Context, id *generated.UserID) (*empty.Empty, error) {
	var echoCtx echo.Context
	err := c.userUsecase.DeleteUserById(echoCtx, model_conversion.FromProtoUserIDToUserID(id))
	return nil, err
}

func (c *usersServiceGRPCServer) CheckExistUserById(ctx context.Context, id *generated.UserID) (*empty.Empty, error) {
	var echoCtx echo.Context
	err := c.userUsecase.CheckExistUserById(echoCtx, model_conversion.FromProtoUserIDToUserID(id))
	return nil, err
}

func (c *usersServiceGRPCServer) GetUserById(ctx context.Context, id *generated.UserID) (*generated.User, error) {
	var echoCtx echo.Context
	user, err := c.userUsecase.GetUserById(echoCtx, model_conversion.FromProtoUserIDToUserID(id))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromUserToProtoUser(user), nil
}

func (c *usersServiceGRPCServer) AddUserContact(ctx context.Context, arguments *generated.AddUserContactArguments) (*generated.Contacts, error) {
	var echoCtx echo.Context
	contacts, err := c.userUsecase.AddUserContact(echoCtx, arguments.UserID, arguments.ContactID)
	if err != nil {
		return nil, err
	}

	return &generated.Contacts{Contacts: model_conversion.FromMembersToProtoMembers(contacts)}, nil
}

func (c *usersServiceGRPCServer) GetUserContacts(ctx context.Context, id *generated.UserID) (*generated.Contacts, error) {
	var echoCtx echo.Context
	contacts, err := c.userUsecase.GetUserContacts(echoCtx, model_conversion.FromProtoUserIDToUserID(id))
	if err != nil {
		return nil, err
	}

	return &generated.Contacts{Contacts: model_conversion.FromMembersToProtoMembers(contacts)}, nil
}

func (c *usersServiceGRPCServer) PutUserById(ctx context.Context, arguments *generated.PutUserArguments) (*generated.User, error) {
	var echoCtx echo.Context
	user, err := c.userUsecase.PutUserById(echoCtx, model_conversion.FromProtoUpdateUserToUpdateUser(arguments.User), arguments.UserID)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromUserToProtoUser(user), nil
}

func (c *usersServiceGRPCServer) GetAllUsersExceptCurrentUser(ctx context.Context, id *generated.UserID) (*generated.Contacts, error) {
	var echoCtx echo.Context
	contacts, err := c.userUsecase.GetAllUsersExceptCurrentUser(echoCtx, model_conversion.FromProtoUserIDToUserID(id))
	if err != nil {
		return nil, err
	}

	return &generated.Contacts{Contacts: model_conversion.FromMembersToProtoMembers(contacts)}, nil
}
