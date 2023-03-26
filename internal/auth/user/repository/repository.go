package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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

func (r *repository) CreateUser(ctx echo.Context, user model.User) (model.User, error) {
	rows, err := r.db.NamedQuery("INSERT INTO profile (username, email, status, password)"+
		"VALUES (:username, :email, :status, :password) RETURNING id", user)

	if err != nil {
		return user, err
	}
	if rows.Next() {
		rows.Scan(&user.Id)
	}

	return user, nil
}

func (r *repository) CheckCorrectPassword(ctx echo.Context, user model.User) error {
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE email=:email AND password=:password", user)

	if err != nil {
		return err
	}
	if !rows.Next() {
		return myErrors.ErrIncorrectPassword
	}

	return nil
}

func (r *repository) CheckExistEmail(ctx echo.Context, email string) error {
	user := model.User{Email: email}
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE email=:email", user)

	if err != nil {
		return err
	}
	if !rows.Next() {
		return myErrors.ErrUserNotFound
	}

	return nil
}

func (r *repository) CheckExistUsername(ctx echo.Context, username string) error {
	user := model.User{Username: username}
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE username=:username", user)

	if err != nil {
		return err
	}
	if !rows.Next() {
		return myErrors.ErrUserNotFound
	}

	return nil
}
