package repository

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"project/internal/auth/session"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strconv"
)

func NewAuthSessionMemoryRepository(db *redis.Client) session.Repository {
	return &repository{db: db}
}

type repository struct {
	db *redis.Client
}

func (r *repository) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	var session model.Session
	result, err := r.db.Get(cookie).Result()
	if err == redis.Nil {
		return session, myErrors.ErrSessionNotFound
	}

	userId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return session, err
	}

	session.UserId = userId
	session.Cookie = cookie

	return session, nil
}

func (r *repository) CreateSession(ctx echo.Context, session model.Session) error {
	err := r.db.Set(session.Cookie, session.UserId, 0).Err()
	return err
}

func (r *repository) DeleteSession(ctx echo.Context, cookie string) error {
	err := r.db.Del(cookie).Err()
	if err == redis.Nil {
		return myErrors.ErrSessionNotFound
	}

	return err
}
