package http

import (
	user "example.com/m/user"
	"fmt"
	"net/http"
)

func NewUserHandler(usecaseImpl user.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			fmt.Println("HANDLER GET SUCCESS")
			usecaseImpl.GetUserById(1)
			break
		case http.MethodPut:
			fmt.Println("HANDLER PUT SUCCESS")
			usecaseImpl.EdidUserById(1, []byte("new data"))
			break
		case http.MethodDelete:
			fmt.Println("HANDLER DELETE SUCCESS")
			usecaseImpl.DeleteUserById(1)
			break
		default:
			break
		}
	}
}
