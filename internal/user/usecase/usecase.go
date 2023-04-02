package usecase

import (
	"github.com/labstack/echo/v4"
	authUser "project/internal/auth/user"
	"project/internal/model"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/pkg/security"
	"project/internal/user"
)

type usecase struct {
	userRepo user.Repository
	authRepo authUser.Repository
}

func NewUserUsecase(userRepo user.Repository, authRepo authUser.Repository) user.Usecase {
	return &usecase{userRepo: userRepo, authRepo: authRepo}
}

func (u *usecase) DeleteUserById(ctx echo.Context, userID uint64) error {
	err := u.userRepo.DeleteUserById(ctx, userID)
	return err
}

func (u *usecase) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	user, err := u.userRepo.GetUserById(ctx, userID)
	return user, err
}

func (u *usecase) PutUserById(ctx echo.Context, updateUser model.UpdateUser, userID uint64) (model.User, error) {
	oldUser := model.User{
		Id:       userID,
		Username: updateUser.Username,
		Email:    updateUser.Email,
		Status:   updateUser.Status,
		Password: updateUser.CurrentPassword,
	}

	validateError := security.ValidateUser(oldUser)
	if validateError != nil {
		return oldUser, httpUtils.ErrorConversion(validateError[0])
	}

	password, err := security.Hash(oldUser.Password)
	oldUser.Password = password
	if err != nil {
		return oldUser, err
	}

	err = u.authRepo.CheckCorrectPassword(ctx, oldUser)
	if err != nil {
		return oldUser, err
	}

	user, err := u.userRepo.UpdateUserById(ctx, oldUser)
	return user, err
}

func (u *usecase) GetUserContacts(ctx echo.Context, userID uint64) ([]model.Contact, error) {
	var contacts []model.Contact
	contactsDB, err := u.userRepo.GetUserContacts(ctx, userID)
	if err != nil {
		return contacts, err
	}

	for _, contact := range contactsDB {
		contacts = append(contacts, model.Contact{
			Username: contact.Username,
			Nickname: contact.Nickname,
			Status:   contact.Status,
		})
	}

	return contacts, err
}

func (u *usecase) AddUserContact(ctx echo.Context, userID uint64, contactID uint64) (model.User, error) {
	userContact := model.UserContact{
		IdUser:    userID,
		IdContact: contactID,
	}
	err := u.userRepo.CheckUserIsContact(ctx, userContact)
	if err != nil {
		return model.User{}, err
	}

	err = u.userRepo.AddUserInContact(ctx, userContact)
	if err != nil {
		return model.User{}, err
	}

	contact, err := u.userRepo.GetUserById(ctx, contactID)
	return contact, err
}

func (u *usecase) CheckExistUserById(ctx echo.Context, userID uint64) error {
	err := u.userRepo.CheckExistUserById(ctx, userID)
	return err
}
