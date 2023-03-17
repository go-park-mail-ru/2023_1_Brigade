package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

func NewUserHandler(e *echo.Echo, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/:userID"

	api := e.Group("api/v1")
	signup := api.Group(userUrl)

	signup.OPTIONS("", handler.GetUserHandler)
	signup.GET("", handler.GetUserHandler)

	return handler
}
