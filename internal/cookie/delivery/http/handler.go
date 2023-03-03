package http

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/cookie"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
)

type authHandler struct {
	usecase cookie.Usecase
}

func (u *authHandler) GetAuthHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		httpUtils.JsonWriteErrors(w, []error{myErrors.ErrCookieNoFound})
		return
	}

	authSession, err := u.usecase.GetSessionByCookie(context.Background(), session.Value)
	if err == nil {
		httpUtils.JsonWriteUserId(w, authSession.UserId)
	} else {
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func NewAuthHandler(r *mux.Router, us cookie.Usecase) {
	handler := authHandler{usecase: us}
	authUrl := "/auth/"

	r.HandleFunc(authUrl, handler.GetAuthHandler).
		Methods("GET")
}
