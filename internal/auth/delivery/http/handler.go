package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/auth"
	"project/internal/pkg/http_utils"
)

type authHandler struct {
	usecase auth.Usecase
}

func (u *authHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonUser, sliceErrors := u.usecase.Signup(context.Background(), r)
	jsonResponse := []byte("")

	if len(sliceErrors) == 0 {
		w.WriteHeader(http.StatusCreated)
		jsonResponse = jsonUser
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		var validateErrors []http_utils.JsonErrors
		for _, err := range sliceErrors {
			validateErrors = append(validateErrors, http_utils.JsonErrors{Err: err})
		}

		jsonResponse, _ = json.Marshal(validateErrors) // TODO ERROR
	}
	_, _ = w.Write(jsonResponse) // TODO ERROR
}

func (u *authHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonUser, err := u.usecase.Login(context.Background(), r)
	jsonResponse := []byte("")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		inJsonError := http_utils.JsonErrors{Err: err}
		jsonResponse, _ = json.Marshal(inJsonError) // TODO ERROR
	} else {
		w.WriteHeader(http.StatusOK)
		jsonResponse = jsonUser
	}

	_, _ = w.Write(jsonResponse) // TODO ERROR
}

func NewAuthHandler(r *mux.Router, us auth.Usecase) {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"
	loginUrl := "/login/"

	r.HandleFunc(signupUrl, handler.SignupHandler).
		Methods("POST")
	r.HandleFunc(loginUrl, handler.LoginHandler).
		Methods("POST")
}
