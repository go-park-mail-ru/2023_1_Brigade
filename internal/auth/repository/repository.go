package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	_, err := r.db.Exec(
		"INSERT INTO profile (username, name, email, status, password) VALUES ($1, $2, $3, $4, $5)",
		user.Username, user.Name, user.Email, user.Status, user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CheckCorrectPassword(ctx context.Context, hashedPassword string) (bool, error) {
	err := r.db.QueryRow("SELECT * FROM profile WHERE password=$1", hashedPassword).Scan()

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return true, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow("SELECT * FROM profile WHERE email=$1", email).
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

func (r *repository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow("SELECT * FROM profile WHERE username=$1", username).
		Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Status, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, myErrors.ErrNoUserFound
		} else {
			return user, err
		}
	}

	return user, myErrors.ErrUsernameIsAlreadyRegistred
}

func (r *repository) GetSessionById(ctx context.Context, userId uint64) (model.Session, error) {
	session := model.Session{}
	err := r.db.QueryRow("SELECT * FROM session WHERE user_id=$1", userId).
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

func (r *repository) CreateSession(ctx context.Context, session model.Session) (model.Session, error) {
	_, err := r.db.Exec(
		"INSERT INTO session (user_id, cookie) VALUES ($1, $2)",
		session.UserId, session.Cookie)

	if err != nil {
		return session, err
	}

	return session, nil
}

func (r *repository) DeleteSession(ctx context.Context, session model.Session) error {
	_, err := r.db.Exec(
		"DELETE FROM session WHERE cookie=$1",
		session.Cookie)

	if err != nil {
		return err
	}

	return nil
}
