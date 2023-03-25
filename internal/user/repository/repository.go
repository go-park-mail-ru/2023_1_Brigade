package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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

func (r *repository) DeleteUserById(ctx echo.Context, userID uint64) error {
	_, err := r.db.Query("DELETE FROM profile WHERE id=$1", userID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r *repository) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return user, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r *repository) GetUserByEmail(ctx echo.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE email=$1", email)

	if errors.Is(err, sql.ErrNoRows) {
		return user, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r *repository) GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error) {
	var contacts []model.User
	err := r.db.Select(&contacts, "SELECT * FROM user_friends WHERE id_user=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return contacts, myErrors.ErrUserNotFound
	}

	return contacts, err
}

func (r *repository) UpdateUserById(ctx echo.Context, user model.User) (model.User, error) {
	rows, err := r.db.NamedQuery(`UPDATE profile SET username=:username, email=:email, status=:status, password=:password  WHERE :id = $1`, user)

	if err != nil {
		return user, err
	}
	if rows.Next() {
		rows.Scan(&user)
	}

	return user, nil
}

func (r *repository) CheckUserIsContact(ctx echo.Context, contact model.UserContact) (bool, error) {
	rows, err := r.db.NamedQuery("SELECT * FROM user_contacts WHERE id_user=:id_user, id_contact:=id_contact", contact)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, myErrors.ErrUserNotFound
	}

	return true, nil
}

func (r *repository) AddUserInContact(ctx echo.Context, contact model.UserContact) error {
	_, err := r.db.NamedQuery("INSERT INTO user_contacts (id_user, id_contact) VALUES (:id_user, :id_contact)", contact)
	if err != nil {
		return err
	}

	return nil
}
