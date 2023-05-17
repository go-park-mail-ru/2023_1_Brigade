package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/model"
	"project/internal/monolithic_services/session"
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
	err := r.db.GetContext(ctx, &session, `SELECT * FROM session WHERE cookie = $1`, cookie)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r repository) CreateSession(ctx context.Context, session model.Session) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO session (cookie, profile_id) VALUES ($1, $2)`, session.Cookie, session.UserId)
	return err
}

func (r repository) DeleteSession(ctx context.Context, cookie string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM session WHERE cookie=$1", cookie)

	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrSessionNotFound
	}

	return err
}
