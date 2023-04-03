package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	authUser "project/internal/auth/user"
	"project/internal/model"
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
	return user, err
}

func (u usecase) PutUserById(ctx echo.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	oldUser := model.User{
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
	return user, err
}

func (u usecase) GetUserContacts(ctx echo.Context, userID uint64) ([]model.Contact, error) {
	var contacts []model.Contact
	contactsFromDB, err := u.userRepo.GetUserContacts(context.Background(), userID)
	if err != nil {
		return []model.Contact{}, err
	}

	for _, contact := range contactsFromDB {
		avatarUrl, err := u.userRepo.GetUserAvatar(context.Background(), userID)
		if err != nil {
			log.Error(err)
		}

		contacts = append(contacts, model.Contact{
			Username: contact.Username,
			Nickname: contact.Nickname,
			Status:   contact.Status,
			Avatar:   avatarUrl,
		})
	}

	return contacts, err
}

func (u usecase) AddUserContact(ctx echo.Context, userID uint64, contactID uint64) (model.User, error) {
	userContact := model.UserContact{
		IdUser:    userID,
		IdContact: contactID,
	}
	err := u.userRepo.CheckUserIsContact(context.Background(), userContact)
	if err != nil {
		return model.User{}, err
	}

	err = u.userRepo.AddUserInContact(context.Background(), userContact)
	if err != nil {
		return model.User{}, err
	}

	contact, err := u.userRepo.GetUserById(context.Background(), contactID)
	return contact, err
}

func (u usecase) CheckExistUserById(ctx echo.Context, userID uint64) error {
	err := u.userRepo.CheckExistUserById(context.Background(), userID)
	return err
}
