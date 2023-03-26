package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	"project/internal/configs"

	log "github.com/sirupsen/logrus"
	myMiddleware "project/internal/middleware"

	httpAuthUser "project/internal/auth/user/delivery/http"
	httpChat "project/internal/chat/delivery/http"
	wsMessages "project/internal/messages/delivery/ws"
	httpUser "project/internal/user/delivery/http"

	usecaseAuthSession "project/internal/auth/session/usecase"
	usecaseAuthUser "project/internal/auth/user/usecase"
	usecaseChat "project/internal/chat/usecase"
	usecaseMessages "project/internal/messages/usecase"
	usecaseUser "project/internal/user/usecase"

	repositoryAuthSession "project/internal/auth/session/repository"
	repositoryAuthUser "project/internal/auth/user/repository"
	repositoryChat "project/internal/chat/repository"
	repositoryMessages "project/internal/messages/repository"
	repositoryUser "project/internal/user/repository"
)

func init() {
	// убрать не забывать при docker
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	defer func() {
		if err := recover(); err != nil {
			log.WithField("error", err).Error("Panic occurred")
		}
	}()

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

	db, err := sqlx.Open(config.DB, config.ConnectionToDB) // ping
	if err != nil {
		log.Fatal(err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	userRepository := repositoryUser.NewUserMemoryRepository(db)
	chatRepository := repositoryChat.NewChatMemoryRepository(db)
	messagesRepository := repositoryMessages.NewMessagesMemoryRepository(db)
	authUserRepository := repositoryAuthUser.NewAuthUserMemoryRepository(db)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(redis)

	userUsecase := usecaseUser.NewUserUsecase(userRepository, authUserRepository)
	authUserUsecase := usecaseAuthUser.NewAuthUserUsecase(authUserRepository, userRepository)
	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	chatUsecase := usecaseChat.NewChatUsecase(chatRepository, userRepository)
	messagesUsecase := usecaseMessages.NewMessagesUsecase(messagesRepository)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.AllowMethods,
		AllowOrigins:     config.AllowOrigins,
		AllowCredentials: config.AllowCredentials,
		AllowHeaders:     config.AllowHeaders,
	}))
	e.Use(myMiddleware.LoggerMiddleware)
	e.Use(myMiddleware.XSSMidlleware)
	e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))

	httpUser.NewUserHandler(e, userUsecase)
	httpAuthUser.NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, userUsecase)
	wsMessages.NewMessagesHandler(e, messagesUsecase)

	e.Logger.Fatal(e.Start(config.Port))
	e.Logger.Fatal(e.Start(":8081"))
}
