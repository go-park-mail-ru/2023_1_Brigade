package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/microservices/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewUserMemoryRepository(db *sqlx.DB) user.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) DeleteUserById(ctx context.Context, userID uint64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM profile WHERE id=$1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return myErrors.ErrUserNotFound
		}

		return err
	}

	return nil
}

func (r repository) GetUserById(ctx context.Context, userID uint64) (model.AuthorizedUser, error) {
	var user model.AuthorizedUser
	err := r.db.GetContext(ctx, &user, "SELECT * FROM profile WHERE id=$1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
		}

		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model.AuthorizedUser, error) {
	var user model.AuthorizedUser
	err := r.db.GetContext(ctx, &user, "SELECT * FROM profile WHERE email=$1", email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
		}

		return model.AuthorizedUser{}, err
	}

	return user, err
}

func (r repository) GetUserContacts(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var contacts []model.UserContact
	err = r.db.SelectContext(ctx, &contacts, "SELECT * FROM user_contacts WHERE id_user=$1", userID)

	if err != nil {
		tx.Rollback() // nolint: errcheck
		return nil, err
	}

	var contactsInfo []model.AuthorizedUser
	for _, contact := range contacts {
		contactInfo, err := r.GetUserById(ctx, contact.IdContact)
		if err != nil {
			tx.Rollback() // nolint: errcheck
			return nil, err
		}

		contactsInfo = append(contactsInfo, contactInfo)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return contactsInfo, nil
}

func (r repository) UpdateUserPasswordById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	err := r.db.GetContext(ctx, &user, `UPDATE profile SET password=$1 WHERE id=$2 RETURNING *`,
		user.Password, user.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
		}

		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) UpdateUserInfoById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	err := r.db.GetContext(ctx, &user, `UPDATE profile SET avatar=$1, nickname=$2, email=$3, status=$4 WHERE id=$5 RETURNING *`,
		user.Avatar, user.Nickname, user.Email, user.Status, user.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
		}

		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) CheckUserIsContact(ctx context.Context, contact model.UserContact) error {
	_, err := r.db.ExecContext(ctx, `SELECT * FROM user_contacts WHERE id_user=$1 AND id_contact=$2`, contact.IdUser, contact.IdContact)
	if err != nil {
		return err
	}

	return myErrors.ErrUserIsAlreadyContact
}

func (r repository) AddUserInContact(ctx context.Context, contact model.UserContact) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO user_contacts (id_user, id_contact) VALUES ($1, $2)`, contact.IdUser, contact.IdContact)
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
	err := r.db.SelectContext(ctx, &users, "SELECT * FROM profile WHERE id != $1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrUserNotFound
		}
		return nil, err
	}

	return users, nil
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
