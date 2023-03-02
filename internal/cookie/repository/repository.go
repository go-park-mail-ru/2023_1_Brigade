package repository

import (
	"context"
	"project/internal/cookie"
)

type repositoryImpl struct {
	db map[string]bool
}

func NewTokenRepository() cookie.Repository {
	result := &repositoryImpl{}
	result.db = map[string]bool{}
	return result
}

func (r *repositoryImpl) InsertTokenInDB(ctx context.Context, token string) bool {
	if _, exists := r.db[token]; exists {
		return false
	}
	r.db[token] = true
	return true
}

func (r *repositoryImpl) GetTokenInDB(ctx context.Context, token string) bool {
	_, exists := r.db[token]
	return exists
}
