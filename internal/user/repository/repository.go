package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
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

func (r *repository) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow("SELECT * FROM profile WHERE id=$1", userID).
		Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Status, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, myErrors.ErrNoUserFound
		} else {
			return user, err
		}
	}

	return user, myErrors.ErrEmailIsAlreadyRegistred
}
