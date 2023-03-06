package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

//func RequestResponseMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Info(r.Method, r.Header, r.Body)
//		w.Header().Set("Content-Type", "application/json")
//		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
//		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		next.ServeHTTP(w, r)
//		//r.Header.Get()
//	})
//}

//func RequestResponseMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Info(r.Method, r.Header, r.Body)
//		w.Header().Set("Content-Type", "application/json")
//		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
//		next.ServeHTTP(w, r)
//		//r.Header.Get()
//	})
//}

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

//func SetupCorsResponse(w *http.ResponseWriter, req *http.Request) {
//	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
//}

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

func Cors(next http.Handler) http.Handler {
	//handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Info(r.Header, r.Host, r.Body)
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Access-Control-Allow-Headers", "*")
	//	w.Header().Set("Access-Control-Allow-Methods", "*")
	//	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//})
	//
	//return handler
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "GET", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{"http://127.0.0.1", "http://95.163.249.116"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "X-Csrf-Token, SameSite=None"},
		Debug:            true,
	})
	return c.Handler(next)
}
