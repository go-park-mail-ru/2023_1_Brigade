package user

import (
	"context"
	"project/internal/model"
)

type Usecase interface {
	DeleteUserById(ctx context.Context, userID uint64) error
	CheckExistUserById(ctx context.Context, userID uint64) error
	GetUserById(ctx context.Context, userID uint64) (model.User, error)
	AddUserContact(ctx context.Context, userID uint64, contactID uint64) ([]model.User, error)
	GetUserContacts(ctx context.Context, userID uint64) ([]model.User, error)
	PutUserById(ctx context.Context, user model.UpdateUser, userID uint64) (model.User, error)
	GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.User, error)
}
