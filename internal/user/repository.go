package user

import (
	"context"
	"project/internal/model"
)

type Repository interface {
	DeleteUserById(ctx context.Context, userID uint64) error
	GetUserById(ctx context.Context, userID uint64) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	AddUserInContact(ctx context.Context, contact model.UserContact) error
	UpdateUserById(ctx context.Context, user model.User) (model.User, error)
	GetUserContacts(ctx context.Context, userID uint64) ([]model.User, error)
	CheckUserIsContact(ctx context.Context, contact model.UserContact) error
	CheckExistUserById(ctx context.Context, userID uint64) error

	GetUserAvatar(ctx context.Context, userID uint64) (string, error)
}
