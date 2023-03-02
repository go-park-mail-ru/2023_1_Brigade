package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/pkg/http_utils"
	"project/internal/user"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := http_utils.ParsingIdUrl(r, "userID")
	jsonUser, err := u.usecase.GetUserById(context.Background(), userID)
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

func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := userHandler{usecase: us}
	userUrl := "/users/{userID:[0-9]+}"

	r.HandleFunc(userUrl, handler.GetUserHandler).
		Methods("GET")
}
