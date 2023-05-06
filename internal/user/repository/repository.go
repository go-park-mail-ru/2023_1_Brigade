package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"project/internal/images"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/user"
)

func NewUserMemoryRepository(db *sqlx.DB, s3 images.Repository) user.Repository {
	return &repository{db: db, s3: s3}
}

type repository struct {
	db *sqlx.DB
	s3 images.Repository
}

func (r repository) DeleteUserById(ctx context.Context, userID uint64) error {
	rows, err := r.db.Query("DELETE FROM profile WHERE id=$1", userID)
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r repository) GetUserById(ctx context.Context, userID uint64) (model.AuthorizedUser, error) {
	var user model.AuthorizedUser
	err := r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.AuthorizedUser{}, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model.AuthorizedUser, error) {
	var user model.AuthorizedUser
	err := r.db.Get(&user, "SELECT * FROM profile WHERE email=$1", email)

	if errors.Is(err, sql.ErrNoRows) {
		return model.AuthorizedUser{}, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r repository) GetUserContacts(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error) {
	var contacts []model.UserContact
	err := r.db.Select(&contacts, "SELECT * FROM user_contacts WHERE id_user=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return []model.AuthorizedUser{}, myErrors.ErrUserNotFound
	}

	var contactsInfo []model.AuthorizedUser
	for _, contact := range contacts {
		contactInfo, err := r.GetUserById(ctx, contact.IdContact)
		if err != nil {
			log.Error(err)
		}

		contactsInfo = append(contactsInfo, contactInfo)
	}

	return contactsInfo, nil
}

func (r repository) UpdateUserById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	result, err := r.db.Exec("UPDATE profile SET username=$1, nickname=$2, status=$3, password=$4 WHERE id=$5", user.Username, user.Nickname, user.Status, user.Password, user.Id)
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	if rowsAffected == 0 {
		return model.AuthorizedUser{}, err
	}

	err = r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", user.Id)
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) CheckUserIsContact(ctx context.Context, contact model.UserContact) error {
	rows, err := r.db.NamedQuery("SELECT * FROM user_contacts WHERE id_user=:id_user AND id_contact=:id_contact", contact)
	defer rows.Close()
	if err == nil && rows.Next() {
		return myErrors.ErrUserIsAlreadyContact
	}

	return err
}

func (r repository) AddUserInContact(ctx context.Context, contact model.UserContact) error {
	rows, err := r.db.NamedQuery("INSERT INTO user_contacts (id_user, id_contact) VALUES (:id_user, :id_contact)", contact)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r repository) CheckExistUserById(ctx context.Context, userID uint64) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE id=$1)", userID)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrUserNotFound
	}

	return nil
}

func (r repository) GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error) {
	var users []model.AuthorizedUser
	rows, err := r.db.Query("SELECT * FROM profile WHERE id != $1", userID)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrUserNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var user model.AuthorizedUser
		err := rows.Scan(&user.Id, &user.Avatar, &user.Username, &user.Nickname, &user.Email, &user.Status, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func (r repository) GetSearchUsers(ctx context.Context, string string) ([]model.AuthorizedUser, error) {
	var searchUsers []model.AuthorizedUser
	err := r.db.Select(&searchUsers, `SELECT * FROM profile WHERE nickname ILIKE $1`, "%"+string+"%")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrUserNotFound
		}

		return nil, err
	}

	return searchUsers, nil
}
