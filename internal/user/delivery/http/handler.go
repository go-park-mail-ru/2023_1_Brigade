package http

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
		log.Error(err)
		httpUtils.JsonWriteInternalError(w)
		return
	}

	user, err := u.usecase.GetUserById(context.Background(), userID)

	if err == nil {
		httpUtils.JsonWriteUserLogin(w, user)
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
