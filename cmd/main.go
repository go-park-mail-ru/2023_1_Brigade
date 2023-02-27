package main

import (
	"net/http"
	chatrepository "project/internal/chat/repository"
	"project/internal/middleware"
	userrepository "project/internal/user/repository"

	"github.com/gorilla/mux"

	chatusecase "project/internal/chat/usecase"
	userusecase "project/internal/user/usecase"

	httpauth "project/internal/auth/delivery/http"
	httpchat "project/internal/chat/delivery/http"
	httpuser "project/internal/user/delivery/http"

	_ "project/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Technogram
// @version 1.0.0
// @description Telegram wannabe

// @host localhost:8081
// @BasePath /

// securityDefinitions.apikey ApikeyAuth
// @in header
// @name Authorization

func main() {
	repositoryUserImpl := userrepository.NewUserMemoryRepository()
	repositoryChatImpl := chatrepository.NewChatMemoryRepository()

	userImpl := userusecase.NewUserUsecase(repositoryUserImpl)
	chatImpl := chatusecase.NewChatUsecase(repositoryChatImpl)

	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	r.Use(middleware.RequestResponseMiddleware)

	httpuser.NewUserHandler(r, userImpl)
	httpchat.NewChatHandler(r, chatImpl)
	httpauth.NewAuthHandler(r)

	http.ListenAndServe(":8081", r)
}
