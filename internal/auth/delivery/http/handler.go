package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
)

type authHandler struct {
	usecase auth.Usecase
}

func (u *authHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpUtils.JsonWriteErrors(w, []error{err})
		return
	}

	user, errors := u.usecase.Signup(context.Background(), user)
	if len(errors) == 0 {

		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)
		if err != nil {
			httpUtils.JsonWriteErrors(w, []error{err})
		}

		httpUtils.SetCookie(w, session)
		httpUtils.JsonWriteUserCreated(w, user)
	} else {
		httpUtils.JsonWriteErrors(w, errors)
	}
}

func (u *authHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpUtils.JsonWriteErrors(w, []error{err})
		return
	}

	user, err := u.usecase.Login(context.Background(), user)
	if err == nil {

		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)
		if err != nil {
			httpUtils.JsonWriteErrors(w, []error{err})
		}

		httpUtils.SetCookie(w, session)
		httpUtils.JsonWriteUserGet(w, user)
	} else {
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func (u *authHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		httpUtils.JsonWriteErrors(w, []error{myErrors.ErrCookieNotFound})
		return
	}

	authSession, err := u.usecase.GetSessionByCookie(context.Background(), session.Value)
	if err == nil {
		user, err := u.usecase.GetUserById(context.Background(), authSession.UserId)
		if err != nil {
			httpUtils.JsonWriteErrors(w, []error{err})
			return
		}

		httpUtils.JsonWriteUserGet(w, user)
	} else {
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func (u *authHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		httpUtils.JsonWriteErrors(w, []error{myErrors.ErrCookieNotFound})
		return
	}

	err = u.usecase.DeleteSessionByCookie(context.Background(), session.Value)
	if err == nil {
		httpUtils.DeleteCookie(w)
		httpUtils.JsonWriteErrors(w, []error{myErrors.SessionSuccessDeleted})
	} else {
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func NewAuthHandler(r *mux.Router, us auth.Usecase) authHandler {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	r.HandleFunc(logoutUrl, handler.LogoutHandler).
		Methods("DELETE", "OPTIONS")
	r.HandleFunc(authUrl, handler.AuthHandler).
		Methods("GET", "OPTIONS")
	r.HandleFunc(signupUrl, handler.SignupHandler).
		Methods("POST", "OPTIONS")
	r.HandleFunc(loginUrl, handler.LoginHandler).
		Methods("POST", "OPTIONS")

	return handler
}
