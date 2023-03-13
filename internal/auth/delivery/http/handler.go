package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
)

type authHandler struct {
	usecase auth.Usecase
}

func (u *authHandler) SignupHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := model.User{}
		err := ctx.Bind(&user)

		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		user, err = u.usecase.Signup(context.Background(), user)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		httpUtils.SetCookie(ctx, session)
		return httpUtils.SendJsonUser(ctx, user, myErrors.UserCreated)
	}
}

func (u *authHandler) LoginHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := model.User{}
		err := ctx.Bind(&user)

		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		user, err = u.usecase.Login(context.Background(), user)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		httpUtils.SetCookie(ctx, session)
		return httpUtils.SendJsonUser(ctx, user, myErrors.UserGetting)
	}
}

func (u *authHandler) AuthHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Cookie("session_id")
		if err != nil {
			return httpUtils.SendJsonError(ctx, myErrors.ErrCookieNotFound)
		}

		authSession, err := u.usecase.GetSessionByCookie(context.Background(), session.Value)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		user, err := u.usecase.GetUserById(context.Background(), authSession.UserId)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		httpUtils.SetCookie(ctx, authSession)
		return httpUtils.SendJsonUser(ctx, user, myErrors.UserGetting)
	}
}

func (u *authHandler) LogoutHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Cookie("session_id")
		if err != nil {
			return httpUtils.SendJsonError(ctx, myErrors.ErrCookieNotFound)
		}

		err = u.usecase.DeleteSessionByCookie(context.Background(), session.Value)
		if err != nil {
			return httpUtils.SendJsonError(ctx, err)
		}

		httpUtils.DeleteCookie(ctx)
		return ctx.NoContent(http.StatusNoContent)
	}
}

func NewAuthHandler(e *echo.Echo, us auth.Usecase) authHandler {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	e.OPTIONS(signupUrl, handler.SignupHandler())
	e.OPTIONS(loginUrl, handler.LoginHandler())
	e.OPTIONS(authUrl, handler.AuthHandler())
	e.OPTIONS(logoutUrl, handler.LogoutHandler())

	e.POST(signupUrl, handler.SignupHandler())
	e.POST(loginUrl, handler.LoginHandler())
	e.GET(authUrl, handler.AuthHandler())
	e.DELETE(logoutUrl, handler.LogoutHandler())

	return handler
}
