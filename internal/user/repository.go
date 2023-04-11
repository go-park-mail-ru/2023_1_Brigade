package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	DeleteUserById(ctx context.Context, userID uint64) error
	GetUserById(ctx context.Context, userID uint64) (model.AuthorizedUser, error)
	GetUserByEmail(ctx context.Context, email string) (model.AuthorizedUser, error)
	AddUserInContact(ctx context.Context, contact model.UserContact) error
	UpdateUserById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error)
	GetUserContacts(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error)
	CheckUserIsContact(ctx context.Context, contact model.UserContact) error
	CheckExistUserById(ctx context.Context, userID uint64) error

	GetUserAvatar(ctx context.Context, userID uint64) (string, error)
	GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error)
}
