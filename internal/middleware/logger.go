package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RequestResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method, r.Header, r.Body)
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

//func Cors(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "*")
//		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
//		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
//		next.ServeHTTP(w, r)
//	})
//	//c := cors.New(cors.Options{
//	//	AllowedMethods:   []string{"POST", "GET", "DELETE"},
//	//	AllowedOrigins:   []string{"*"},
//	//	AllowCredentials: true,
//	//	AllowedHeaders:   []string{"*"},
//	//	Debug:            true,
//	//})
//	//return c.Handler(next)
//}

func SetupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

//func Cors(h http.Handler) http.Handler {
//	c := cors.New(cors.Options{
//		AllowedMethods:   []string{"*"},
//		AllowedOrigins:   []string{"*"},
//		AllowCredentials: true,
//		AllowedHeaders:   []string{"*"},
//		Debug:            true,
//	})
//	return c.Handler(h)
//}

//func Cors(next http.Handler) http.Handler {
//	//handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	//	log.Info(r.Header, r.Host, r.Body)
//	//	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
//	//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, Content-Length, User-Agent, X-CSRF-Token")
//	//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
//	//	w.Header().Set("Access-Control-Allow-Credentials", "true")
//	//})
//
//	//return handler
//}
