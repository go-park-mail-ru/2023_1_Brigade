package middleware

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
)

func HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestId := rand.Int63()
		log.Info("Incoming request: ", ctx.Request().URL, ", ip: ", ctx.RealIP(), ", method: ", ctx.Request().Method, ", request_id: ", requestId)
		if err := next(ctx); err != nil {
			log.Error("Error: ", err, ", request_id: ", requestId)
			statusCode := httpUtils.StatusCode(err)
			if statusCode == 500 {
				return ctx.JSON(statusCode, myErrors.ErrInternal)
			}

			return ctx.JSON(statusCode, err)
		}

		log.Info("HTTP code: ", ctx.Response().Status, ", request_id: ", requestId)
		return nil
	}
}
