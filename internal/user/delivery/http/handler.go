package http

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"project/internal/model"
	"project/internal/user"
	"strconv"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler(ctx echo.Context) error {
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

func (u *userHandler) GetCurrentUserHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	user, err := u.usecase.GetUserById(ctx, session.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u *userHandler) DeleteUserHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	err := u.usecase.DeleteUserById(ctx, session.UserId)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (u *userHandler) PutUserHandler(ctx echo.Context) error {
	var updateUser model.UpdateUser
	body := ctx.Get("body").([]byte)

	err := json.Unmarshal(body, &updateUser)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	user, err := u.usecase.PutUserById(ctx, updateUser, session.UserId)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u *userHandler) GetUserContactsHandler(ctx echo.Context) error {
	//session := ctx.Get("session").(model.Session)
	//contacts, err := u.usecase.GetUserContacts(ctx, session.UserId)
	//if err != nil {
	//	return err
	//}

	// test
	contacts := []model.Contact{
		{
			Username: "marcussss1",
			Nickname: "Marcus1",
			Status:   "Marcus1 is cool",
		},
		{
			Username: "marcussss2",
			Nickname: "Marcus2",
			Status:   "Marcus2 is cool",
		},
		{
			Username: "marcussss3",
			Nickname: "Marcus3",
			Status:   "Marcus3 is cool",
		},
	}

	return ctx.JSON(http.StatusOK, contacts)
}

func (u *userHandler) UserAddContactHandler(ctx echo.Context) error {
	contactID, err := strconv.ParseUint(ctx.Param("userID"), 10, 64)
	if err != nil {
		return err
	}

	session := ctx.Get("session").(model.Session)
	contact, err := u.usecase.AddUserContact(ctx, session.UserId, contactID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, contact)
}

func NewUserHandler(e *echo.Echo, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/:userID"
	deleteUserUrl := "/users/remove"
	currentUserUrl := "/users/settings"
	userContactsUrl := "/users/contacts"
	userAddContactUrl := "/users/{userID}/add"

	api := e.Group("api/v1")
	user := api.Group(userUrl)
	deleteUser := api.Group(deleteUserUrl)
	currentUser := api.Group(currentUserUrl)
	userContacts := api.Group(userContactsUrl)
	userAddContact := api.Group(userAddContactUrl)

	user.OPTIONS("", handler.GetUserHandler)
	user.GET("", handler.GetUserHandler)

	currentUser.OPTIONS("", handler.PutUserHandler)
	currentUser.POST("", handler.PutUserHandler)

	deleteUser.OPTIONS("", handler.DeleteUserHandler)
	currentUser.DELETE("", handler.DeleteUserHandler)

	currentUser.OPTIONS("", handler.GetCurrentUserHandler)
	currentUser.GET("", handler.GetCurrentUserHandler)

	userContacts.OPTIONS("", handler.GetUserContactsHandler)
	userContacts.GET("", handler.GetUserContactsHandler)

	userAddContact.OPTIONS("", handler.UserAddContactHandler)
	userAddContact.POST("", handler.UserAddContactHandler)

	return handler
}
