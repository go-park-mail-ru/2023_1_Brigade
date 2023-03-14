package repository

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/user"
)

func NewUserMemoryRepository(db *sql.DB) user.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetUserById(ctx echo.Context, userID uint64) (user model.User, err error) {
	err = r.db.QueryRow("SELECT * FROM profile WHERE id=$1", userID).
		Scan(&user.Id, &user.Username, &user.Email, &user.Status, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrUserNotFound
	}
	return
}
