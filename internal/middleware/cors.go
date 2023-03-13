package middleware

//next echo.HandlerFunc) echo.HandlerFunc

//func Cors(next echo.HandlerFunc) echo.HandlerFunc {
//	c := cors.New(cors.Options{
//		AllowedMethods:   []string{"POST", "GET", "DELETE", "OPTIONS"},
//		AllowedOrigins:   []string{"http://95.163.249.116"},
//		AllowCredentials: true,
//		AllowedHeaders:   []string{"Content-Type", "Content-Length", "X-Csrf-Token", "SameSite=None"},
//		Debug:            true,
//	})
//	return c.Handler(next)
//}

//func Cors(next http.Handler) http.Handler {
//	c := cors.New(cors.Options{
//		AllowedMethods:   []string{"POST", "GET", "DELETE", "OPTIONS"},
//		AllowedOrigins:   []string{"http://95.163.249.116"},
//		AllowCredentials: true,
//		AllowedHeaders:   []string{"Content-Type", "Content-Length", "X-Csrf-Token", "SameSite=None"},
//		Debug:            true,
//	})
//	return c.Handler(next)
//}
