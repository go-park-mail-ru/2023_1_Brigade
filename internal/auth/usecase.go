package auth

import (
	"context"
	"net/http"
	"project/internal/pkg/http_utils"
)

type Usecase interface {
	Signup(ctx context.Context, r *http.Request) http_utils.Response
}
