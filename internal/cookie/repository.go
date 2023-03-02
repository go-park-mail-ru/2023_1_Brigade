package cookie

import (
	"context"
	"project/internal/pkg/http_utils"
)

type Repository interface {
	InsertTokenInDB(ctx context.Context, token string) bool
	GetTokenInDB(ctx context.Context, token string) bool
}
