package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"project/internal/config"
	"project/internal/model"
	authSession "project/internal/monolithic_services/session"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	metrics "project/internal/pkg/metrics/prometheus"
	"strconv"
	"strings"
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

func GenerateCSRFToken(userID string) (string, error) {
	h := hmac.New(sha256.New, []byte("csrf"))

	t := time.Second * 86400
	timeNow := time.Now().Add(t).Unix()

	data := fmt.Sprintf("%s:%d", userID, timeNow)
	h.Write([]byte(data))

	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(timeNow, 10)

	return token, nil
}

func RefreshIfNeededCSRFToken(token string, userID string) (string, error) {
	tokenData := strings.Split(token, ":")

	if len(tokenData) != 2 {
		return "", errors.New("неверный csrf токен")
	}

	//tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	//if err != nil {
	//	return "", errors.New("неверный csrf токен")
	//}

	//if tokenExp > time.Now().Unix()+viper.GetInt64(constants.ViperCSRFTTLKey)/2 {
	//	return "", nil
	//}

	return GenerateCSRFToken(userID)
}

func CSRFMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookieCSRF, err := ctx.Cookie(config.CsrfCookie)
			if err != nil || len(cookieCSRF.Value) == 0 {
				return errors.New("осутствует csrf токен")
			}
			tokenCSRF := ctx.QueryParam(config.CsrfCookie)

			if tokenCSRF != cookieCSRF.Value {
				log.Error("Cookie token: %s; Query token: %s", cookieCSRF.Value, tokenCSRF)
				return errors.New("неверный csrf токен")
			}

			session := ctx.Get("session").(model.Session)
			ctx.Set("session", session)

			newTokenCSRF, err := RefreshIfNeededCSRFToken(tokenCSRF, session.Cookie)
			if err != nil {
				return err
			}

			if len(newTokenCSRF) != 0 {
				cookie := &http.Cookie{
					Name:     "session_id",
					Value:    session.Cookie,
					HttpOnly: false,
					Path:     "/",
					Expires:  time.Now().Add(24 * time.Hour * 30),
					SameSite: http.SameSiteNoneMode,
					Secure:   true,
				}
				ctx.SetCookie(cookie)
			}
			//ctx.SetCookie(utils.CreateCookie(constants.CookieKeyCSRFToken, newTokenCSRF, viper.GetInt64(constants.ViperCSRFTTLKey)))

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
