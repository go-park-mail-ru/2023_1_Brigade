package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/auth/session"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthSessionMemoryRepository(db *sqlx.DB) session.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	var session model.Session
	err := r.db.Get(&session, `SELECT * FROM session WHERE cookie = $1`, cookie)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r repository) CreateSession(ctx context.Context, session model.Session) error {
	rows, err := r.db.NamedQuery(`INSERT INTO session (cookie, profile_id) VALUES (:cookie, :profile_id)`, session)
	defer rows.Close()

	return err
}

func (r repository) DeleteSession(ctx context.Context, cookie string) error {
	rows, err := r.db.Query("DELETE FROM session WHERE cookie=$1", cookie)
	defer rows.Close()

	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrSessionNotFound
	}

	return err
}
