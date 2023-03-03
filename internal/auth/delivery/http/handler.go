package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/auth"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	httpUtils "project/internal/pkg/http_utils"
	"time"
)

type authHandler struct {
	usecase auth.Usecase
}

func (u *authHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Error(err)
		httpUtils.JsonWriteInternalError(w)
		return
	}

	user, errors := u.usecase.Signup(context.Background(), user)

	if len(errors) == 0 {
		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)

		if err != nil {
			log.Error(err)
			httpUtils.JsonWriteErrors(w, []error{err})
		}

		httpUtils.SetCookie(w, session)
		httpUtils.JsonWriteUser(w, user)
	} else {
		log.Error(err)
		httpUtils.JsonWriteErrors(w, errors)
	}
}

func (u *authHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Error(err)
		httpUtils.JsonWriteInternalError(w)
		return
	}

	user, err = u.usecase.Login(context.Background(), user)

	if err == nil {
		session, err := u.usecase.CreateSessionById(context.Background(), user.Id)

		if err != nil {
			httpUtils.JsonWriteErrors(w, []error{err})
		}

		httpUtils.SetCookie(w, session)
		httpUtils.JsonWriteUser(w, user)
	} else {
		log.Error(err)
		httpUtils.JsonWriteErrors(w, []error{err})
	}
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
		log.Error(err)
		httpUtils.JsonWriteErrors(w, []error{err})
	}
}

func (u *authHandler) DeleteLogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		log.Error(err)
		httpUtils.JsonWriteErrors(w, []error{myErrors.ErrCookieNoFound})
		return
	}

	err = u.usecase.DeleteSessionByCookie(context.Background(), session.Value)
	if err == nil {
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		httpUtils.JsonWriteErrors(w, []error{myErrors.CookieSuccessDeleted})
	} else {
		log.Error(err)
		httpUtils.JsonWriteErrors(w, []error{myErrors.CookieSuccessDeleted})
	}
}

func NewAuthHandler(r *mux.Router, us auth.Usecase) {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"
	loginUrl := "/login/"
	logoutUrl := "/logout/"
	authUrl := "/auth/"

	r.HandleFunc(logoutUrl, handler.DeleteLogoutHandler).
		Methods("DELETE")
	r.HandleFunc(authUrl, handler.GetAuthHandler).
		Methods("GET")
	r.HandleFunc(signupUrl, handler.SignupHandler).
		Methods("POST")
	r.HandleFunc(loginUrl, handler.LoginHandler).
		Methods("POST")
}
