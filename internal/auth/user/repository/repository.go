package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	auth "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthUserMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) CreateUser(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	row, err := r.db.NamedQuery(`INSERT INTO profile (avatar, username, nickname, email, status, password) `+
		`VALUES (:avatar, :username, :nickname, :email, :status, :password) RETURNING id`, user)

	if err != nil {
		return model.AuthorizedUser{}, err
	}
	if row.Next() {
		err = row.Scan(&user.Id)
		if err != nil {
			return model.AuthorizedUser{}, err
		}
	}

	return user, nil
}

func (r repository) CheckCorrectPassword(ctx context.Context, email string, password string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)", email, password)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrIncorrectPassword
	}

	return nil
}

func (r repository) CheckExistEmail(ctx context.Context, email string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)", email)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrEmailNotFound
	}

	return nil
}

func (r repository) CheckExistUsername(ctx context.Context, username string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)", username)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrUsernameNotFound
	}

	return nil
}
