package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"project/internal/auth"
	"project/internal/model"
)

func NewAuthMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sqlx.DB
}

func (u *repositoryImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	_, err := u.db.Exec(
		"INSERT INTO profile (username, name, email, password) VALUES ($1, $2, $3, $4)",
		user.Username, user.Name, user.Email, user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *repositoryImpl) CheckExistUserByEmail(ctx context.Context, email string) bool {
	err := u.db.QueryRow("SELECT * FROM profile WHERE email=$1", email).Scan()

	if err == sql.ErrNoRows {
		return false
	}

	return true
}
