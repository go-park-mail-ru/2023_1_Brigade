package http

import (
	"github.com/labstack/echo/v4"
	"project/internal/user"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
		//userID, err := httpUtils.ParsingIdUrl(ctx, "userID")
		//
		//if err != nil {
		//	return httpUtils.SendJsonError(ctx, err)
		//}
		//
		//user, err := u.usecase.GetUserById(ctx, userID)
		//if err != nil {
		//	return httpUtils.SendJsonError(ctx, err)
		//}
		//
		//return httpUtils.SendJsonUser(ctx, user, myErrors.UserGetting)
	}
}

func NewUserHandler(e *echo.Echo, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/{userID:[0-9]+}"

	e.OPTIONS(userUrl, handler.GetUserHandler())
	e.GET(userUrl, handler.GetUserHandler())

	return handler
}
