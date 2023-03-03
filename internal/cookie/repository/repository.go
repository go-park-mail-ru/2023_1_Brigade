package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/cookie"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthMemoryRepository(db *sqlx.DB) cookie.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) GetSessionById(ctx context.Context, userId uint64) (model.Session, error) {
	session := model.Session{}
	err := r.db.QueryRow("SELECT * FROM session WHERE id=$1", userId).
		Scan(&session.UserId, &session.Cookie)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return session, myErrors.ErrNoSessionFound
		} else {
			return session, err
		}
	}

	return session, myErrors.ErrSessionIsAlreadyCrated
}

func (r *repository) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	session := model.Session{}
	err := r.db.QueryRow("SELECT * FROM session WHERE cookie=$1", cookie).
		Scan(&session.UserId, &session.Cookie)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return session, myErrors.ErrNoSessionFound
		} else {
			return session, err
		}
	}

	return session, myErrors.ErrSessionIsAlreadyCrated
}

func (r *repository) CreateCookie(ctx context.Context, session model.Session) (model.Session, error) {
	_, err := r.db.Exec(
		"INSERT INTO profile (user_id, cookie) VALUES ($1, $2)",
		session.UserId, session.Cookie)

	if err != nil {
		return session, err
	}

	return session, nil
}
