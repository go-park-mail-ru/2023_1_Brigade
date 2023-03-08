package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	httpUtils "project/internal/pkg/http_utils"
	"project/internal/user"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := httpUtils.ParsingIdUrl(r, "userID")

	if err != nil {
		httpUtils.JsonWriteErrors(w, []error{err})
		return
	}

	user, err := u.usecase.GetUserById(context.Background(), userID)
	if err == nil {
		httpUtils.JsonWriteUserGet(w, user)
	} else {
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func NewUserHandler(r *mux.Router, us user.Usecase) userHandler {
	handler := userHandler{usecase: us}
	userUrl := "/users/{userID:[0-9]+}"

	r.HandleFunc(userUrl, handler.GetUserHandler).
		Methods("GET")

	return handler
}
