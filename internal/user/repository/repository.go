package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/user"
)

func NewUserMemoryRepository(db *sqlx.DB) user.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) DeleteUserById(ctx context.Context, userID uint64) error {
	_, err := r.db.Query("DELETE FROM profile WHERE id=$1", userID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r repository) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE email=$1", email)

	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r repository) GetUserContacts(ctx context.Context, userID uint64) ([]model.User, error) {
	var contacts []model.UserContact
	err := r.db.Select(&contacts, "SELECT * FROM user_contacts WHERE id_user=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return []model.User{}, myErrors.ErrUserNotFound
	}

	var contactsInfo []model.User
	for _, contact := range contacts {
		contactInfo, err := r.GetUserById(ctx, contact.IdContact)
		if err != nil {
			log.Error(err)
		}

		contactsInfo = append(contactsInfo, contactInfo)
	}

	return contactsInfo, nil
}

func (r repository) UpdateUserById(ctx context.Context, user model.User) (model.User, error) {
	rows, err := r.db.NamedQuery(`UPDATE profile SET username=:username, email=:email, status=:status, password=:password  WHERE :id = $1`, user)

	if err != nil {
		return model.User{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user)
		if err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (r repository) CheckUserIsContact(ctx context.Context, contact model.UserContact) error {
	rows, err := r.db.NamedQuery("SELECT * FROM user_contacts WHERE id_user=:id_user, id_contact:=id_contact", contact)

	if err != nil {
		return err
	}
	if !rows.Next() {
		return myErrors.ErrUserNotFound
	}

	return nil
}

func (r repository) AddUserInContact(ctx context.Context, contact model.UserContact) error {
	_, err := r.db.NamedQuery("INSERT INTO user_contacts (id_user, id_contact) VALUES (:id_user, :id_contact)", contact)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) CheckExistUserById(ctx context.Context, userID uint64) error {
	rows, err := r.db.Query("SELECT EXISTS(SELECT 1 FROM profile WHERE id=$1)", userID)

	if err != nil {
		return err
	}
	if !rows.Next() {
		return myErrors.ErrUserNotFound
	}

	return nil
}

func (r repository) GetUserAvatar(ctx context.Context, userID uint64) (string, error) {
	var userAvatar model.UserAvatar
	err := r.db.Get(&userAvatar, "SELECT * FROM users_avatar WHERE id_user=$1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", myErrors.ErrAvatarNotFound
		}
		return "", err
	}

	var avatarUrl model.ImageUrl
	err = r.db.Get(&avatarUrl, "SELECT * FROM images_urls WHERE id_image=$1", userAvatar.IdImage)

	if errors.Is(err, sql.ErrNoRows) {
		return "", myErrors.ErrImageNotFound
	}

	return avatarUrl.ImageUrl, err
}
