package main

import (
	handler "example.com/m/user/delivery/http"
	"example.com/m/user/repository"
	"example.com/m/user/usecase"
	"net/http"
)

func main() {

	repositoryImpl := repository.NewUserMemoryRepository()
	userImpl := usecase.NewUserUsecase(repositoryImpl)
	handl := handler.NewUserHandler(userImpl)

	http.HandleFunc("/", handl)
	http.HandleFunc("/user/", handl)
	http.HandleFunc("/chat/", handl)

	http.ListenAndServe(":8081", nil)
}
