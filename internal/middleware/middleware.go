package middleware

import (
	"encoding/json"
	"github.com/gorilla/csrf"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	authSession "project/internal/auth/session"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
)

type jsonError struct {
	Err error `json:"error"`
}

func (j jsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Err.Error())
}

func CSRFMiddleware() echo.MiddlewareFunc {
	csrfMiddleware := csrf.Protect(
		[]byte("random-secret-key"),
		csrf.Secure(false),
	)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Генерируем токен CSRF
			token := csrf.Token(c.Request())
			//token := "123"
			// Добавляем токен CSRF в заголовок ответа
			c.Response().Header().Set("X-CSRF-Token", token)

			// Продолжаем обработку запроса
			csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next(c)
			})).ServeHTTP(c.Response().Writer, c.Request())

			return nil
		}
	}
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestId := rand.Int63()
		log.Info("Incoming request: ", ctx.Request().URL, ", ip: ", ctx.RealIP(), ", method: ", ctx.Request().Method, ", request_id: ", requestId)

		//log.Warn(ctx.Request().Header)
		//curl -X 'POST' 'http://localhost:8081/api/v1/signup/' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{ "username": "<a onblur="alert(secret)" href="http://www.google.com">Google</a>", "email": "danssssddsila22om", "name": "string", "password": "tests", "status":"i am star" }'

		if err := next(ctx); err != nil {
			statusCode := httpUtils.StatusCode(err)
			log.Error("HTTP code: ", statusCode, ", Error: ", err, ", request_id: ", requestId)
			if statusCode == 500 {
				return ctx.JSON(statusCode, jsonError{Err: myErrors.ErrInternal})
			}

			return ctx.JSON(statusCode, jsonError{Err: err})
		}

		log.Info("HTTP code: ", ctx.Response().Status, ", request_id: ", requestId)
		return nil
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
				return ctx.JSON(httpUtils.StatusCode(myErrors.ErrCookieNotFound), jsonError{Err: myErrors.ErrCookieNotFound})
			}

			authSession, err := authSessionUsecase.GetSessionByCookie(ctx, session.Value)
			if err != nil {
				return ctx.JSON(httpUtils.StatusCode(err), jsonError{Err: err})
			}

			ctx.Set("session", authSession)
			return next(ctx)
		}
	}
}
