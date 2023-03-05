package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	httpauth "project/internal/auth/delivery/http"
	authrepository "project/internal/auth/repository"
	authusecase "project/internal/auth/usecase"
	"project/internal/middleware"
	httpuser "project/internal/user/delivery/http"
	userrepository "project/internal/user/repository"
	userusecase "project/internal/user/usecase"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	connStr := "user=brigade password=123 dbname=brigade sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	repositoryAuth := authrepository.NewAuthMemoryRepository(db)
	repositoryUser := userrepository.NewUserMemoryRepository(db)

	usecaseAuth := authusecase.NewAuthUsecase(repositoryAuth)
	usecaseUser := userusecase.NewUserUsecase(repositoryUser)

	r := mux.NewRouter()
	r.Use(middleware.RequestResponseMiddleware)
	//r.Use(middleware.Cors)

	httpauth.NewAuthHandler(r, usecaseAuth)
	httpuser.NewUserHandler(r, usecaseUser)

	//r.Handle()
	//corsRouter := middleware.Cors()

	//server := http.Server{
	//	Addr: ":" + config.Port,
	//	middleware.Cors(ht),
	//	Handler: corsRouter,
	//}

	http.ListenAndServe(":8081", r)
}
