package repository

import (
	"context"
	"database/sql"
	"errors"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthMemoryRepository(db *sql.DB) auth.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	_, err := r.db.Exec(
		"INSERT INTO profile (username, email, status, password) VALUES ($1, $2, $3, $4)",
		user.Username, user.Email, user.Status, user.Password)

	if err != nil {
		return user, err
	}

	user, err = r.GetUserByEmail(ctx, user.Email) // для получения нормального айдишника
	return user, err
}

func (r *repository) CheckCorrectPassword(ctx context.Context, hashedPassword string) (bool, error) {
	err := r.db.QueryRow("SELECT * FROM profile WHERE password=$1", hashedPassword).Scan()

	return err == nil, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	err = r.db.QueryRow("SELECT * FROM profile WHERE email=$1", email).
		Scan(&user.Id, &user.Username, &user.Email, &user.Status, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrUserNotFound
	}

	return
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	err = r.db.QueryRow("SELECT * FROM profile WHERE username=$1", username).
		Scan(&user.Id, &user.Username, &user.Email, &user.Status, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrUserNotFound
	}
	return
}

func (r *repository) GetUserById(ctx context.Context, userID uint64) (user model.User, err error) {
	err = r.db.QueryRow("SELECT * FROM profile WHERE id=$1", userID).
		Scan(&user.Id, &user.Username, &user.Email, &user.Status, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrUserNotFound
	}
	return
}

func (r *repository) GetSessionById(ctx context.Context, userId uint64) (session model.Session, err error) {
	err = r.db.QueryRow("SELECT * FROM session WHERE user_id=$1", userId).
		Scan(&session.UserId, &session.Cookie)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrSessionNotFound
	}
	return
}

func (r *repository) GetSessionByCookie(ctx context.Context, cookie string) (session model.Session, err error) {
	err = r.db.QueryRow("SELECT * FROM session WHERE cookie=$1", cookie).
		Scan(&session.UserId, &session.Cookie)

	if errors.Is(err, sql.ErrNoRows) {
		err = myErrors.ErrSessionNotFound
	}
	return
}

func (r *repository) CreateSession(ctx context.Context, session model.Session) error {
	_, err := r.db.Exec(
		"INSERT INTO session (user_id, cookie) VALUES ($1, $2)",
		session.UserId, session.Cookie)

	return err
}

func (r *repository) DeleteSession(ctx context.Context, session model.Session) error {
	_, err := r.db.Exec(
		"DELETE FROM session WHERE cookie=$1",
		session.Cookie)

	return err
}
