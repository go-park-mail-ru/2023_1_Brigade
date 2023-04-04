package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	authUser "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/pkg/security"
	"project/internal/pkg/validation"
	"project/internal/user"
)

type usecase struct {
	userRepo user.Repository
	authRepo authUser.Repository
}

func NewUserUsecase(userRepo user.Repository, authRepo authUser.Repository) user.Usecase {
	return &usecase{userRepo: userRepo, authRepo: authRepo}
}

func (u usecase) DeleteUserById(ctx echo.Context, userID uint64) error {
	err := u.userRepo.DeleteUserById(context.Background(), userID)
	return err
}

func (u usecase) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	user, err := u.userRepo.GetUserById(context.Background(), userID)
	return model_conversion.FromAuthorizedUserToUser(user), err
}

func (u usecase) PutUserById(ctx echo.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	oldUser := model.AuthorizedUser{
		Id:       userID,
		Username: updateUser.Username,
		Email:    updateUser.Email,
		Status:   updateUser.Status,
		Password: updateUser.CurrentPassword,
	}

	validateError := validation.ValidateUser(oldUser)
	if validateError != nil {
		return model.User{}, validation.ErrorConversion(validateError[0])
	}

	password := security.Hash([]byte(oldUser.Password))
	oldUser.Password = password

	err := u.authRepo.CheckCorrectPassword(context.Background(), oldUser.Email, oldUser.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err := u.userRepo.UpdateUserById(context.Background(), oldUser)
	return model_conversion.FromAuthorizedUserToUser(user), err
}

func (u usecase) GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error) {
	contactsFromDB, err := u.userRepo.GetUserContacts(context.Background(), userID)
	if err != nil {
		return []model.User{}, err
	}

	contacts := model_conversion.FromAuthorizedUserArrayToUserArray(contactsFromDB)
	return contacts, err
}

func (u usecase) AddUserContact(ctx echo.Context, userID uint64, contactID uint64) ([]model.User, error) {
	if userID == contactID {
		return nil, myErrors.ErrUserIsAlreadyContact
	}

	userContact := model.UserContact{
		IdUser:    userID,
		IdContact: contactID,
	}

	err := u.userRepo.CheckExistUserById(context.Background(), contactID)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.CheckUserIsContact(context.Background(), userContact)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.AddUserInContact(context.Background(), userContact)
	if err != nil {
		return nil, err
	}

	contacts, err := u.userRepo.GetUserContacts(context.Background(), userID)
	return model_conversion.FromAuthorizedUserArrayToUserArray(contacts), err
}

func (u usecase) CheckExistUserById(ctx echo.Context, userID uint64) error {
	err := u.userRepo.CheckExistUserById(context.Background(), userID)
	return err
}
