package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	"project/internal/user"
)

type userServiceGRPCClient struct {
	userClient generated.UsersClient
}

func NewUserServiceGRPSClient(con *grpc.ClientConn) user.Usecase {
	return &userServiceGRPCClient{
		userClient: generated.NewUsersClient(con),
	}
}

func (u userServiceGRPCClient) DeleteUserById(ctx echo.Context, userID uint64) error {
	_, err := u.userClient.DeleteUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))
	return err
}

func (u userServiceGRPCClient) CheckExistUserById(ctx echo.Context, userID uint64) error {
	_, err := u.userClient.CheckExistUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))
	return err
}

func (u userServiceGRPCClient) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	user, err := u.userClient.GetUserById(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromProtoUserToUser(user), nil
}

func (u userServiceGRPCClient) AddUserContact(ctx echo.Context, userID uint64, contactID uint64) ([]model.User, error) {
	contacts, err := u.userClient.AddUserContact(context.TODO(),
		&generated.AddUserContactArguments{
			UserID:    userID,
			ContactID: contactID,
		},
	)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(contacts.Contacts), nil
}

func (u userServiceGRPCClient) GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error) {
	contacts, err := u.userClient.GetUserContacts(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(contacts.Contacts), nil
}

func (u userServiceGRPCClient) PutUserById(ctx echo.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	user, err := u.userClient.PutUserById(context.TODO(),
		&generated.PutUserArguments{
			User:   model_conversion.FromUpdateUserToProtoUpdateUser(updateUser),
			UserID: userID,
		})
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromProtoUserToUser(user), nil
}

func (u userServiceGRPCClient) GetAllUsersExceptCurrentUser(ctx echo.Context, userID uint64) ([]model.User, error) {
	users, err := u.userClient.GetAllUsersExceptCurrentUser(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(users.Contacts), nil
}
