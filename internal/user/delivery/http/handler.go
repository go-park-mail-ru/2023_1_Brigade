package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"project/internal/user"
	"project/pkg"
)

type userHandler struct {
	usecase user.Usecase
}

func writeInLogAndWriter(w http.ResponseWriter, message []byte) {
	log.Printf(string(message))
	w.Write(message)
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := pkg.ParsingIdUrl(r, "userID")

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	user, err := u.usecase.GetUserById(userID)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonUser, err := json.Marshal(user)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonUser)
}

func (u *userHandler) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := pkg.ParsingIdUrl(r, "userID")

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	user, err := u.usecase.ChangeUserById(userID, []byte(""))

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonUser, err := json.Marshal(user)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonUser)
}

func (u *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := pkg.ParsingIdUrl(r, "userID")

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	err = u.usecase.DeleteUserById(userID)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonError, err := json.Marshal(err)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonError)
}

func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := userHandler{usecase: us}
	userUrl := "/user/{userID:[0-9]+}"

	r.HandleFunc(userUrl, handler.GetUserHandler).
		Methods("GET")
	r.HandleFunc(userUrl, handler.PutUserHandler).
		Methods("PUT")
	r.HandleFunc(userUrl, handler.DeleteUserHandler).
		Methods("DELETE")
}
