package http

import (
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

func (u *authHandler) SignupHandler(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)

	if err != nil {
		return err
	}

	user, err = u.usecase.Signup(ctx, user)
	if err != nil {
		return err
	}

	session, err := u.usecase.CreateSessionById(ctx, user.Id)
	if err != nil {
		return err
	}

	httpUtils.SetCookie(ctx, session)
	return ctx.JSON(http.StatusCreated, user)
}

func (u *authHandler) LoginHandler(ctx echo.Context) error {
	user := model.User{}
	err := ctx.Bind(&user)

	if err != nil {
		return err
	}

	user, err = u.usecase.Login(ctx, user)
	if err != nil {
		return err
	}

	session, err := u.usecase.CreateSessionById(ctx, user.Id)
	if err != nil {
		return err
	}

	httpUtils.SetCookie(ctx, session)
	return ctx.JSON(http.StatusOK, user)
}

func (u *authHandler) AuthHandler(ctx echo.Context) error {
	session, err := ctx.Cookie("session_id")
	if err != nil {
		return myErrors.ErrCookieNotFound
	}

	authSession, err := u.usecase.GetSessionByCookie(ctx, session.Value)
	if err != nil {
		return err
	}

	user, err := u.usecase.GetUserById(ctx, authSession.UserId)
	if err != nil {
		return err
	}

	httpUtils.SetCookie(ctx, authSession)
	return ctx.JSON(http.StatusOK, user)
}

func (u *authHandler) LogoutHandler(ctx echo.Context) error {
	session, err := ctx.Cookie("session_id")
	if err != nil {
		return myErrors.ErrCookieNotFound
	}

	err = u.usecase.DeleteSessionByCookie(ctx, session.Value)
	if err != nil {
		return err
	}

	httpUtils.DeleteCookie(ctx)
	return ctx.NoContent(http.StatusNoContent)
}

func NewAuthHandler(e *echo.Echo, us auth.Usecase) authHandler {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	e.OPTIONS(signupUrl, handler.SignupHandler)
	e.OPTIONS(loginUrl, handler.LoginHandler)
	e.OPTIONS(authUrl, handler.AuthHandler)
	e.OPTIONS(logoutUrl, handler.LogoutHandler)

	e.POST(signupUrl, handler.SignupHandler)
	e.POST(loginUrl, handler.LoginHandler)
	e.GET(authUrl, handler.AuthHandler)
	e.DELETE(logoutUrl, handler.LogoutHandler)

	return handler
}
