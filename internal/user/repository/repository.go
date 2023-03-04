package repository

import (
	"context"
	"database/sql"
	"errors"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/user"

	"github.com/jmoiron/sqlx"
)

func NewUserMemoryRepository(db *sqlx.DB) user.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) GetUserById(ctx context.Context, userID uint64) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(`
	SELECT 
		id,
		username,
		name,
		email,
		status
	FROM 
		profile 
	WHERE 
		id = $1`,
		userID).
		Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, myErrors.ErrUserNotFound
		} else {
			return user, err
		}
	}

	return user, myErrors.ErrEmailIsAlreadyRegistred
}
