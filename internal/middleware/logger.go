package middleware

import (
	"net/http"
)

func RequestResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//func CheckAuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		next.ServeHTTP(w, r)
//	})
//}

func Cors(next http.Handler) http.Handler {
	//handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Info(r.Header, r.Host, r.Body)
	//	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, Content-Length, User-Agent, X-CSRF-Token")
	//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	//	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//})

	return handler
}
