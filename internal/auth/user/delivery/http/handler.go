package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	authSession "project/internal/auth/session"
	authUser "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/user"
)

type authHandler struct {
	userUsecase        user.Usecase
	authUserUsecase    authUser.Usecase
	authSessionUsecase authSession.Usecase
}

func (u authHandler) SignupHandler(ctx echo.Context) error {
	var registrationUser model.RegistrationUser
	err := ctx.Bind(&registrationUser)
	if err != nil {
		return err
	}

	registrationUser = httpUtils.SanitizeStruct(registrationUser).(model.RegistrationUser)

	user, err := u.authUserUsecase.Signup(context.TODO(), registrationUser)
	if err != nil {
		return err
	}

	session, err := u.authSessionUsecase.CreateSessionById(context.TODO(), user.Id)
	if err != nil {
		return err
	}

	httpUtils.SetCookie(ctx, session)
	return ctx.JSON(http.StatusCreated, user)
}

func (u authHandler) LoginHandler(ctx echo.Context) error {
	var loginUser model.LoginUser
	err := ctx.Bind(&loginUser)
	if err != nil {
		return err
	}

	loginUser = httpUtils.SanitizeStruct(loginUser).(model.LoginUser)

	user, err := u.authUserUsecase.Login(context.TODO(), loginUser)
	if err != nil {
		return err
	}

	session, err := u.authSessionUsecase.CreateSessionById(context.TODO(), user.Id)
	if err != nil {
		return err
	}

	httpUtils.SetCookie(ctx, session)
	return ctx.JSON(http.StatusOK, user)
}

func (u authHandler) AuthHandler(ctx echo.Context) error {
	session, err := ctx.Cookie("session_id")
	if err != nil {
		return myErrors.ErrCookieNotFound
	}

	authSession, err := u.authSessionUsecase.GetSessionByCookie(context.TODO(), session.Value)
	if err != nil {
		return err
	}

	user, err := u.userUsecase.GetUserById(context.TODO(), authSession.UserId)
	if err != nil {
		return err
	}

	user = httpUtils.SanitizeStruct(user).(model.User)

	httpUtils.SetCookie(ctx, authSession)
	return ctx.JSON(http.StatusOK, user)
}

func (u authHandler) LogoutHandler(ctx echo.Context) error {
	session, err := ctx.Cookie("session_id")
	if err != nil {
		return myErrors.ErrCookieNotFound
	}

	err = u.authSessionUsecase.DeleteSessionByCookie(context.TODO(), session.Value)
	if err != nil {
		return err
	}

	httpUtils.DeleteCookie(ctx)
	return ctx.NoContent(http.StatusNoContent)
}

func NewAuthHandler(e *echo.Echo, authUserUsecase authUser.Usecase, authSessionUsecase authSession.Usecase, userUsecase user.Usecase) authHandler {
	handler := authHandler{authUserUsecase: authUserUsecase, authSessionUsecase: authSessionUsecase, userUsecase: userUsecase}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	api := e.Group("api/v1")

	signup := api.Group(signupUrl)
	login := api.Group(loginUrl)
	logout := api.Group(logoutUrl)
	auth := api.Group(authUrl)

	signup.POST("", handler.SignupHandler)
	login.POST("", handler.LoginHandler)
	auth.GET("", handler.AuthHandler)
	logout.DELETE("", handler.LogoutHandler)

	return handler
}
