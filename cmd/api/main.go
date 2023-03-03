package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	httpauth "project/internal/auth/delivery/http"
	authrepository "project/internal/auth/repository"
	authusecase "project/internal/auth/usecase"
	"project/internal/middleware"
	httpuser "project/internal/user/delivery/http"
	userrepository "project/internal/user/repository"
	userusecase "project/internal/user/usecase"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	connStr := "user=golang password=golang dbname=golang sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	repositoryAuth := authrepository.NewAuthMemoryRepository(db)
	repositoryUser := userrepository.NewUserMemoryRepository(db)

	usecaseAuth := authusecase.NewAuthUsecase(repositoryAuth)
	usecaseUser := userusecase.NewUserUsecase(repositoryUser)

	r := mux.NewRouter()
	r.Use(middleware.RequestResponseMiddleware)

	httpauth.NewAuthHandler(r, usecaseAuth)
	httpuser.NewUserHandler(r, usecaseUser)

	http.ListenAndServe(":8081", r)
}
