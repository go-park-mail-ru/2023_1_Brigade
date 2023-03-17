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

func (r *repository) GetUserByEmail(ctx echo.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE email=$1", email)

	if errors.Is(err, sql.ErrNoRows) {
		return user, myErrors.ErrUserNotFound
	}

	return user, err
}

func (r *repository) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", userID)

	if errors.Is(err, sql.ErrNoRows) {
		return user, myErrors.ErrUserNotFound
	}

	return user, err
}
