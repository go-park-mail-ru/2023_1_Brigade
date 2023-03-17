package main

import (
	httpAuth "project/internal/auth/delivery/http"
	httpChat "project/internal/chat/delivery/http"
	myMiddleware "project/internal/middleware"
	httpUser "project/internal/user/delivery/http"

	usecaseAuth "project/internal/auth/usecase"
	usecaseChat "project/internal/chat/usecase"
	usecaseUser "project/internal/user/usecase"

	repositoryAuth "project/internal/auth/repository"
	repositoryChat "project/internal/chat/repository"
	repositoryUser "project/internal/user/repository"

	log "github.com/sirupsen/logrus"
	"project/internal/configs"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	yamlPath, exists := os.LookupEnv("YAML_PATH")
	if !exists {
		log.Fatal("Yaml path not found")
	}

	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	var config configs.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(config.DB, config.ConnectionToDB)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositoryUser.NewUserMemoryRepository(db)
	authRepository := repositoryAuth.NewAuthMemoryRepository(db)
	chatRepository := repositoryChat.NewChatMemoryRepository(db)

	userUsecase := usecaseUser.NewUserUsecase(userRepository)
	authUsecase := usecaseAuth.NewAuthUsecase(authRepository, userRepository)
	chatUsecase := usecaseChat.NewChatUsecase(chatRepository)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.AllowMethods,
		AllowOrigins:     config.AllowOrigins,
		AllowCredentials: config.AllowCredentials,
		AllowHeaders:     config.AllowHeaders,
	}))
	e.Use(myMiddleware.LoggerMiddleware)
	e.Use(myMiddleware.AuthMiddleware(authUsecase))

	httpUser.NewUserHandler(e, userUsecase)
	httpAuth.NewAuthHandler(e, authUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, authUsecase)

	e.Logger.Fatal(e.Start(config.Port))
}
