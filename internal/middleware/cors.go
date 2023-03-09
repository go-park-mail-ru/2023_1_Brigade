package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "GET", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{"http://127.0.0.1", "http://localhost:63343", "http://localhost:63343/frontend/2023_1_Brigade/src/127.0.0.1:8081/signup/", "http://95.163.249.116", "http://127.0.0.1:5500", "http://95.163.249.116:8081/logout/", "http://95.163.249.116:8081/auth/", "http://95.163.249.116:8081/signup/", "http://95.163.249.116:8081/login/"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "X-Csrf-Token", "SameSite=None"},
		Debug:            true,
	})
	return c.Handler(next)
}
