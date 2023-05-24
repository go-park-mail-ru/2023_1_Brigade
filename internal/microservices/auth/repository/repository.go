package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	auth "project/internal/microservices/auth"
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
	err := r.db.QueryRowContext(ctx, `INSERT INTO profile (avatar, username, nickname, email, status, password) `+
		`VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		"", user.Username, user.Nickname, user.Email, user.Status, user.Password).Scan(&user.Id)
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) UpdateUserAvatar(ctx context.Context, url string, userID uint64) (model.AuthorizedUser, error) {
	var user model.AuthorizedUser
	err := r.db.GetContext(ctx, &user, `UPDATE profile SET avatar=$1 WHERE id=$2 RETURNING *`, url, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
		}
		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) CheckCorrectPassword(ctx context.Context, email string, password string) error {
	var exists bool
	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)", email, password)

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
	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)", email)

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
	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)", username)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrUsernameNotFound
	}

	return nil
}
