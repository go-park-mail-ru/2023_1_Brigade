package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/user"
)

type userHandler struct {
	usecase user.Usecase
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//userID, err := http_utils.ParsingIdUrl(r, "userID")
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//user, err := u.usecase.GetUserById(r.Context(), userID)
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//jsonUser, err := json.Marshal(user)
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//w.Write(jsonUser)
}

func (u *userHandler) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	//userID, err := http_utils.ParsingIdUrl(r, "userID")
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//user, err := u.usecase.ChangeUserById(r.Context(), userID, []byte(""))
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//jsonUser, err := json.Marshal(user)
	//
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//w.Write(jsonUser)
}

func (u *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//	userID, err := http_utils.ParsingIdUrl(r, "userID")
	//
	//	if err != nil {
	//		w.Write([]byte(err.Error()))
	//		return
	//	}
	//
	//	err = u.usecase.DeleteUserById(r.Context(), userID)
	//
	//	if err != nil {
	//		w.Write([]byte(err.Error()))
	//		return
	//	}
	//
	//	jsonError, err := json.Marshal(err)
	//
	//	if err != nil {
	//		w.Write([]byte(err.Error()))
	//		return
	//	}
	//
	//	w.Write(jsonError)
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
