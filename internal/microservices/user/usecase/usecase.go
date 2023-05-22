package usecase

import (
	"context"
	authUser "project/internal/microservices/auth"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/images"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
)

type usecase struct {
	userRepo      user.Repository
	authRepo      authUser.Repository
	imagesUsecase images.Usecase
}

func NewUserUsecase(userRepo user.Repository, authRepo authUser.Repository) user.Usecase {
	return &usecase{userRepo: userRepo, authRepo: authRepo}
}

func (u usecase) DeleteUserById(ctx context.Context, userID uint64) error {
	err := u.userRepo.DeleteUserById(ctx, userID)
	return err
}

func (u usecase) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	user, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	return model_conversion.FromAuthorizedUserToUser(user), err
}

func (u usecase) PutUserById(ctx context.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	switch updateUser.CurrentPassword {
	case "":
		user := model.AuthorizedUser{
			Id:       userID,
			Avatar:   updateUser.NewAvatarUrl,
			Email:    updateUser.Email,
			Nickname: updateUser.Nickname,
			Status:   updateUser.Status,
		}

		user, err := u.userRepo.UpdateUserInfoById(ctx, user)
		if err != nil {
			return model.User{}, err
		}

		return model_conversion.FromAuthorizedUserToUser(user), nil
	default:
		userFromDB, err := u.userRepo.GetUserById(context.TODO(), userID)
		if err != nil {
			return model.User{}, err
		}

		currentPassword := security.Hash([]byte(updateUser.CurrentPassword))
		if currentPassword != userFromDB.Password {
			return model.User{}, myErrors.ErrIncorrectPassword
		}

		newPassword := security.Hash([]byte(updateUser.NewPassword))

		userFromDB.Password = newPassword
		user, err := u.userRepo.UpdateUserPasswordById(ctx, userFromDB)
		if err != nil {
			return model.User{}, err
		}

		return model_conversion.FromAuthorizedUserToUser(user), nil
	}
}

func (u usecase) GetUserContacts(ctx context.Context, userID uint64) ([]model.User, error) {
	contacts, err := u.userRepo.GetUserContacts(ctx, userID)
	if err != nil {
		return []model.User{}, err
	}

	return model_conversion.FromAuthorizedUserArrayToUserArray(contacts), err
}

func (u usecase) AddUserContact(ctx context.Context, userID uint64, contactID uint64) ([]model.User, error) {
	if userID == contactID {
		return nil, myErrors.ErrUserIsAlreadyContact
	}

	userContact := model.UserContact{
		IdUser:    userID,
		IdContact: contactID,
	}

	err := u.userRepo.CheckExistUserById(ctx, contactID)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.CheckUserIsContact(ctx, userContact)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.AddUserInContact(ctx, userContact)
	if err != nil {
		return nil, err
	}

	contacts, err := u.userRepo.GetUserContacts(ctx, userID)
	return model_conversion.FromAuthorizedUserArrayToUserArray(contacts), err
}

func (u usecase) CheckExistUserById(ctx context.Context, userID uint64) error {
	err := u.userRepo.CheckExistUserById(ctx, userID)
	return err
}

func (u usecase) GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.User, error) {
	users, err := u.userRepo.GetAllUsersExceptCurrentUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromAuthorizedUserArrayToUserArray(users), err
}

func (u usecase) GetSearchUsers(ctx context.Context, string string) ([]model.User, error) {
	searchContacts, err := u.userRepo.GetSearchUsers(ctx, string)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromAuthorizedUserArrayToUserArray(searchContacts), err
}
