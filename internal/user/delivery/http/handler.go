package http

import (
	log "github.com/sirupsen/logrus"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/model"
	"project/internal/user"
	"strconv"
)

type userHandler struct {
	usecase user.Usecase
}

func (u userHandler) GetUserHandler(ctx echo.Context) error {
	userID, err := strconv.ParseUint(ctx.Param("userID"), 10, 64)
	if err != nil {
		return err
	}

	user, err := u.usecase.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u userHandler) GetCurrentUserHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	user, err := u.usecase.GetUserById(ctx, session.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u userHandler) DeleteUserHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	err := u.usecase.DeleteUserById(ctx, session.UserId)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (u userHandler) PutUserHandler(ctx echo.Context) error {
	var updateUser model.UpdateUser
	err := ctx.Bind(&updateUser)
	if err != nil {
		log.Warn(err)
		return err
	}

	session := ctx.Get("session").(model.Session)
	user, err := u.usecase.PutUserById(ctx, updateUser, session.UserId)

	if err != nil {
		log.Warn(err)
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u userHandler) GetUserContactsHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	contacts, err := u.usecase.GetAllUsersExceptCurrentUser(ctx, session.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, contacts)
}

func (u userHandler) UserAddContactHandler(ctx echo.Context) error {
	contactID, err := strconv.ParseUint(ctx.Param("userID"), 10, 64)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	contacts, err := u.usecase.AddUserContact(ctx, session.UserId, contactID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, contacts)
}

func NewUserHandler(e *echo.Echo, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/:userID/"
	deleteUserUrl := "/users/remove/"
	currentUserUrl := "/users/settings/"
	userContactsUrl := "/users/contacts/"
	userAddContactUrl := "/users/:userID/add/"

	api := e.Group("api/v1")
	user := api.Group(userUrl)
	deleteUser := api.Group(deleteUserUrl)
	currentUser := api.Group(currentUserUrl)
	userContacts := api.Group(userContactsUrl)
	userAddContact := api.Group(userAddContactUrl)

	user.GET("", handler.GetUserHandler)
	currentUser.PUT("", handler.PutUserHandler)
	deleteUser.DELETE("", handler.DeleteUserHandler)
	currentUser.GET("", handler.GetCurrentUserHandler)
	userContacts.GET("", handler.GetUserContactsHandler)
	userAddContact.POST("", handler.UserAddContactHandler)

	return handler
}
