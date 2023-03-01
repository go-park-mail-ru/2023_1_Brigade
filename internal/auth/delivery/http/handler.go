package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/auth"
	"project/internal/pkg/http_utils"
)

type authHandler struct {
	usecase auth.Usecase
}

func (u *authHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	response := u.usecase.Signup(context.Background(), r)
	http_utils.SendJsonResponse(w, response)
}

func NewChatHandler(r *mux.Router, us auth.Usecase) {
	handler := authHandler{usecase: us}
	signupUrl := "/signup/"

	r.HandleFunc(signupUrl, handler.SignupHandler).
		Methods("POST")
}
