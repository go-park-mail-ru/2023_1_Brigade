package repository

import (
	"context"
	"github.com/redis/go-redis/v9"

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

func (r repository) GetSessionByCookie(ctx context.Context, cookie string) (model.Session, error) {
	var session model.Session
	result, err := r.db.Get(context.Background(), cookie).Result()
	if err == redis.Nil {
		return model.Session{}, myErrors.ErrSessionNotFound
	}

	userId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return model.Session{}, err
	}

	session.UserId = userId
	session.Cookie = cookie

	return session, nil
}

func (r repository) CreateSession(ctx context.Context, session model.Session) error {
	err := r.db.Set(context.Background(), session.Cookie, session.UserId, 0).Err()
	return err
}

func (r repository) DeleteSession(ctx context.Context, cookie string) error {
	err := r.db.Del(context.Background(), cookie).Err()
	if err == redis.Nil {
		return myErrors.ErrSessionNotFound
	}

	return err
}
