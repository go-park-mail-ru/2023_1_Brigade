package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/auth"
	"project/internal/model"
	my_errors "project/internal/pkg/errors"
)

func NewAuthMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sqlx.DB
}

func (u *repositoryImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	_, err := u.db.Exec(
		"INSERT INTO profile (username, name, email, status, password) VALUES ($1, $2, $3, $4, $5)",
		user.Username, user.Name, user.Email, user.Status, user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *repositoryImpl) CheckCorrectPassword(ctx context.Context, hashedPassword string) (bool, error) {
	err := u.db.QueryRow("SELECT * FROM profile WHERE password=$1", hashedPassword).Scan()

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return true, nil
}

func (u *repositoryImpl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}
	err := u.db.QueryRow("SELECT * FROM profile WHERE email=$1", email).
		Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Status, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, my_errors.NoUserFound
		} else {
			return user, err
		}
	}

	return user, my_errors.EmailIsAlreadyRegistred
}

func (u *repositoryImpl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	user := model.User{}
	err := u.db.QueryRow("SELECT * FROM profile WHERE username=$1", username).
		Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Status, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, my_errors.NoUserFound
		} else {
			return user, err
		}
	}

	return user, my_errors.UsernameIsAlreadyRegistred
}
