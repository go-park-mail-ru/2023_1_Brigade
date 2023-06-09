package client

import (
	"context"
	"google.golang.org/grpc"
	"project/internal/generated"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
)

type userServiceGRPCClient struct {
	userClient generated.UsersClient
}

func NewUserServiceGRPSClient(con *grpc.ClientConn) user.Usecase {
	return &userServiceGRPCClient{
		userClient: generated.NewUsersClient(con),
	}
}

func (u userServiceGRPCClient) DeleteUserById(ctx context.Context, userID uint64) error {
	_, err := u.userClient.DeleteUserById(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	return err
}

func (u userServiceGRPCClient) CheckExistUserById(ctx context.Context, userID uint64) error {
	_, err := u.userClient.CheckExistUserById(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	return err
}

func (u userServiceGRPCClient) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	user, err := u.userClient.GetUserById(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromProtoUserToUser(user), nil
}

func (u userServiceGRPCClient) AddUserContact(ctx context.Context, userID uint64, contactID uint64) ([]model.User, error) {
	contacts, err := u.userClient.AddUserContact(ctx,
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

func (u userServiceGRPCClient) GetUserContacts(ctx context.Context, userID uint64) ([]model.User, error) {
	contacts, err := u.userClient.GetUserContacts(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(contacts.Contacts), nil
}

func (u userServiceGRPCClient) PutUserById(ctx context.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	user, err := u.userClient.PutUserById(ctx,
		&generated.PutUserArguments{
			User:   model_conversion.FromUpdateUserToProtoUpdateUser(updateUser),
			UserID: userID,
		})
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromProtoUserToUser(user), nil
}

func (u userServiceGRPCClient) GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.User, error) {
	users, err := u.userClient.GetAllUsersExceptCurrentUser(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(users.Contacts), nil
}

func (u userServiceGRPCClient) GetSearchUsers(ctx context.Context, string string, userID uint64) ([]model.User, error) {
	searchUsers, err := u.userClient.GetSearchUsers(ctx, &generated.SearchUsersArguments{
		String_: string,
		UserID:  userID,
	})
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoMembersToMembers(searchUsers.Contacts), nil
}
