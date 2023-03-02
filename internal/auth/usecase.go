package auth

import (
	"context"
	"net/http"
)

type Usecase interface {
	Signup(ctx context.Context, r *http.Request) ([]byte, []error)
	Login(ctx context.Context, r *http.Request) ([]byte, error)
}
