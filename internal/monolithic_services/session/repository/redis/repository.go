package redis

import (
	"context"
	"github.com/redis/go-redis/v9"

	"project/internal/model"
	"project/internal/monolithic_services/session"
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
	result, err := r.db.Get(context.TODO(), cookie).Result()
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
	err := r.db.Set(context.TODO(), session.Cookie, session.UserId, 0).Err()
	return err
}

func (r repository) DeleteSession(ctx context.Context, cookie string) error {
	err := r.db.Del(context.TODO(), cookie).Err()
	if err == redis.Nil {
		return myErrors.ErrSessionNotFound
	}

	return err
}
