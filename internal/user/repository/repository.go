package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/model"
	my_errors "project/internal/pkg/errors"
	"project/internal/user"
)

func NewUserMemoryRepository(db *sqlx.DB) user.Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sqlx.DB
}

func (u *repositoryImpl) GetUserById(ctx context.Context, userID int) (model.User, error) {
	user := model.User{}
	err := u.db.QueryRow("SELECT * FROM profile WHERE id=$1", userID).
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
