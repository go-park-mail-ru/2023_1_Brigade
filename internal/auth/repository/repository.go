package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) CreateUser(ctx echo.Context, user model.User) (model.User, error) {
	rows, err := r.db.NamedQuery("INSERT INTO profile (username, email, status, password) "+
		"VALUES (:username, :email, :status, :password) RETURNING id", user)

	if err != nil {
		return user, err
	}
	if rows.Next() {
		rows.Scan(&user.Id)
	}

	return user, nil
}

func (r *repository) CheckCorrectPassword(ctx echo.Context, user model.User) (bool, error) {
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE email=:email AND password=:password", user)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, myErrors.ErrIncorrectPassword
	}

	return true, nil
}

func (r *repository) CheckExistEmail(ctx echo.Context, email string) (bool, error) {
	user := model.User{Email: email}
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE email=:email", user)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, myErrors.ErrUserNotFound
	}

	return true, nil
}

func (r *repository) CheckExistUsername(ctx echo.Context, username string) (bool, error) {
	user := model.User{Username: username}
	rows, err := r.db.NamedQuery("SELECT * FROM profile WHERE username=:username", user)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, myErrors.ErrUserNotFound
	}

	return true, nil
}

func (r *repository) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	var session model.Session
	err := r.db.Get(&session, "SELECT * FROM session WHERE cookie=$1", cookie)

	if errors.Is(err, sql.ErrNoRows) {
		return session, myErrors.ErrSessionNotFound
	}

	return session, err
}

func (r *repository) CreateSession(ctx echo.Context, session model.Session) error {
	_, err := r.db.NamedQuery("INSERT INTO session (profile_id, cookie) VALUES (:profile_id, :cookie)", session)
	return err
}

func (r *repository) DeleteSession(ctx echo.Context, cookie string) error {
	_, err := r.db.Query("DELETE FROM session WHERE cookie=$1", cookie)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrSessionNotFound
	}

	return err
}
