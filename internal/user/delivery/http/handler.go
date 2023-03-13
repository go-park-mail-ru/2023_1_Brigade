package http

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/user"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Error("\nOK\n")
		userID, err := httpUtils.ParsingIdUrl(ctx, "userID")

		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		user, err := u.usecase.GetUserById(context.Background(), userID)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		return httpUtils.SendJsonUser(ctx, user, myErrors.UserGetting)
	}
}

//func (u *authHandler) AuthHandler() echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//	session, err := ctx.Cookie("session_id")
//	if err != nil {
//	return httpUtils.SendJsonError(ctx, err)
//}
//
//	authSession, err := u.usecase.GetSessionByCookie(context.Background(), session.Value)
//	if err != nil {
//	return httpUtils.SendJsonError(ctx, err)
//}
//
//	user, err := u.usecase.GetUserById(context.Background(), authSession.UserId)
//	if err != nil {
//	return httpUtils.SendJsonError(ctx, err)
//}
//
//	httpUtils.SetCookie(ctx, authSession)
//	return httpUtils.SendJsonUser(ctx, user, myErrors.UserGetting)
//}
//}
//}

func NewUserHandler(e *echo.Echo, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/{userID:[0-9]+}"

	e.OPTIONS(userUrl, handler.GetUserHandler())
	e.GET(userUrl, handler.GetUserHandler())

	return handler
}
