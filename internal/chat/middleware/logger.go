package middleware

import (
	"fmt"
	"net/http"
)

func RequestResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware begin")
		next.ServeHTTP(w, r)
		fmt.Println("Middleware End")
	})
}
