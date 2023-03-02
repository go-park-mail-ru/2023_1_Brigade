package cookie

import (
	"context"
	"project/internal/pkg/http_utils"
)

type Usecase interface {
	CreateNewToken(ctx context.Context) http_utils.Response
	GetToken(ctx context.Context, token string) http_utils.Response
}
