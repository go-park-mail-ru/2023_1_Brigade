package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/user"
)

type authHandler struct {
	authUsecase auth.Usecase
	userUsecase user.Usecase
}

func (u *authHandler) SignupHandler(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)

	if err != nil {
		return err
	}

	user, err = u.authUsecase.Signup(ctx, user)
	if err != nil {
		return err
	}

	session, err := u.authUsecase.CreateSessionById(ctx, user.Id)
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

	user, err = u.authUsecase.Login(ctx, user)
	if err != nil {
		return err
	}

	session, err := u.authUsecase.CreateSessionById(ctx, user.Id)
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

	authSession, err := u.authUsecase.GetSessionByCookie(ctx, session.Value)
	if err != nil {
		return err
	}

	user, err := u.userUsecase.GetUserById(ctx, authSession.UserId)
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

	err = u.authUsecase.DeleteSessionByCookie(ctx, session.Value)
	if err != nil {
		return err
	}

	httpUtils.DeleteCookie(ctx)
	return ctx.NoContent(http.StatusNoContent)
}

func NewAuthHandler(e *echo.Echo, authUsecase auth.Usecase, userUsecase user.Usecase) authHandler {
	handler := authHandler{authUsecase: authUsecase, userUsecase: userUsecase}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	api := e.Group("api/v1")

	signup := api.Group(signupUrl)
	login := api.Group(loginUrl)
	logout := api.Group(logoutUrl)
	auth := api.Group(authUrl)

	signup.OPTIONS("", handler.SignupHandler)
	login.OPTIONS("", handler.LoginHandler)
	logout.OPTIONS("", handler.AuthHandler)
	auth.OPTIONS("", handler.LogoutHandler)

	signup.POST("", handler.SignupHandler)
	login.POST("", handler.LoginHandler)
	auth.GET("", handler.AuthHandler)
	logout.DELETE("", handler.LogoutHandler)

	return handler
}
