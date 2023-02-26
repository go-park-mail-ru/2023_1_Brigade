package main

import (
	"github.com/gorilla/mux"
	"net/http"
	chatrepository "project/internal/chat/repository"
	"project/internal/middleware"
	userrepository "project/internal/user/repository"

	chatusecase "project/internal/chat/usecase"
	userusecase "project/internal/user/usecase"

	httpchat "project/internal/chat/delivery/http"
	httpuser "project/internal/user/delivery/http"
)

func main() {

	repositoryUserImpl := userrepository.NewUserMemoryRepository()
	repositoryChatImpl := chatrepository.NewChatMemoryRepository()

	userImpl := userusecase.NewUserUsecase(repositoryUserImpl)
	chatImpl := chatusecase.NewChatUsecase(repositoryChatImpl)

	r := mux.NewRouter()

	r.Use(middleware.RequestResponseMiddleware)

	httpuser.NewUserHandler(r, userImpl)
	httpchat.NewChatHandler(r, chatImpl)

	http.ListenAndServe(":8081", r)
}
