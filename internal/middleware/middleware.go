package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	authSession "project/internal/monolithic_services/session"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	metrics "project/internal/pkg/metrics/prometheus"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type jsonError struct {
	Err error `json:"error"`
}

func (j jsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Err.Error())
}

type GRPCMiddleware struct {
	metric *metrics.MetricsGRPC
}

func NewGRPCMiddleware(metric *metrics.MetricsGRPC) *GRPCMiddleware {
	return &GRPCMiddleware{metric: metric}
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestId := rand.Int63()
		log.Info("Incoming request: ", ctx.Request().URL, ", ip: ", ctx.RealIP(), ", method: ", ctx.Request().Method, ", request_id: ", requestId)

		if err := next(ctx); err != nil {
			statusCode := httpUtils.StatusCode(err)
			log.Error("HTTP code: ", statusCode, ", Error: ", err, ", request_id: ", requestId)
			if statusCode == 500 {
				jsonErr, err := json.Marshal(jsonError{Err: myErrors.ErrInternal})
				if err != nil {
					log.Error(err)
				}

				return ctx.JSONBlob(statusCode, jsonErr)
			}

			jsonErr, err := json.Marshal(jsonError{Err: err})
			if err != nil {
				log.Error(err)
			}

			return ctx.JSONBlob(statusCode, jsonErr)
		}

		log.Info("HTTP code: ", ctx.Response().Status, ", request_id: ", requestId)
		return nil
	}
}

func CSRFMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			clientCsrf := ctx.Request().Header.Values("X-CSRF-Token")
			if clientCsrf == nil || len(clientCsrf) == 1 {
				cookie := &http.Cookie{
					Name:     "_csrf",
					Value:    uuid.NewString(),
					HttpOnly: false,
					Path:     "/login",
					Expires:  time.Now().Add(60 * time.Second),
					SameSite: http.SameSiteNoneMode,
					Secure:   true,
				}

				ctx.SetCookie(cookie)
				return next(ctx)
			}

			log.Info(clientCsrf)

			csrf := ctx.Get("_csrf").(string)

			log.Info(csrf)

			if clientCsrf[0] != csrf {
				return errors.New("неправильный токен")
			}

			return next(ctx)
		}
	}
}

func AuthMiddleware(authSessionUsecase authSession.Usecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			signupUrl := "/api/v1/signup/"
			loginUrl := "/api/v1/login/"
			logoutUrl := "/api/v1/logout/"
			authUrl := "/api/v1/auth/"
			url := ctx.Request().URL.String()

			if url == signupUrl || url == loginUrl || url == logoutUrl || url == authUrl {
				return next(ctx)
			}

			session, err := ctx.Cookie("session_id")
			if err != nil {
				jsonErr, err := json.Marshal(jsonError{Err: myErrors.ErrCookieNotFound})
				if err != nil {
					log.Error(err)
				}

				return ctx.JSONBlob(httpUtils.StatusCode(myErrors.ErrCookieNotFound), jsonErr)
			}

			authSession, err := authSessionUsecase.GetSessionByCookie(context.TODO(), session.Value)
			if err != nil {
				jsonErr, err := json.Marshal(jsonError{Err: err})
				if err != nil {
					log.Error(err)
				}

				return ctx.JSONBlob(httpUtils.StatusCode(err), jsonErr)
			}

			ctx.Set("session", authSession)
			return next(ctx)
		}
	}
}

func (m *GRPCMiddleware) GRPCMetricsMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, uHandler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := uHandler(ctx, req)
	if err != nil {
		return nil, err
	}

	errStatus, _ := status.FromError(err)
	code := errStatus.Code()

	m.metric.ResponseTime.WithLabelValues(code.String(), info.FullMethod).Observe(time.Since(start).Seconds())
	m.metric.Hits.Inc()

	return resp, err
}
