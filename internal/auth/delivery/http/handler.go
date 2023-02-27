package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type emptyBody struct {
}

// @Summary Auth
// @Security ApiKeyAuth
// @Tags auth
// @Description auth with cookie
// @ID auth-cookie
// @Accept  json
// @Produce  json
// @Success 200 {object} emptyBody
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/ [put]
func Auth(w http.ResponseWriter, r *http.Request) {
}

// @Summary Sign In
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func SignIn(w http.ResponseWriter, r *http.Request) {
}

// @Summary Logout
// @Security ApiKeyAuth
// @Tags auth
// @Description logout
// @ID logout
// @Accept  json
// @Produce  json
// @Param input body User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
}

// @Summary Sign Up
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func SignUp(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler(r *mux.Router) {
	r.HandleFunc("/auth", Auth).Methods("PUT")
	r.HandleFunc("/auth/sign-in", SignIn).Methods("POST")
	r.HandleFunc("/auth/logout", Logout).Methods("POST")
	r.HandleFunc("/auth/sign-up", SignUp).Methods("POST")
}
