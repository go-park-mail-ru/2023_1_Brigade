// func (u *usecaseImpl) CheckAuth(cts context.Context, r *http.Request) http_utils.Response {
// 	header, exists := r.Header["Auth"]
// 	if !exists || len(header) == 0 {
// 		return http_utils.Response{Status: http_utils.STATUS_UNAUTHORIZED}
// 	}
// 	token := header[0]
// 	if u
// 	if token != "ok" {
// 		return http_utils.Response{Status: http_utils.STATUS_UNAUTHORIZED}
// 	}
// 	return http_utils.Response{Status: http_utils.STATUS_OK}
// }


package http

import (
	"project/internal/cookie"
	"github.com/gorilla/mux"
)

type cookieHandler struct {
	usecase cookie.Usecase
}

func NewCookieHandler(r *mux.Router, us cookie.Usecase) {
	handler := cookieHandler{usecase: us}
	authUrl := "/auth"

	r.HandleFunc(authUrl, handler).Methods()
}