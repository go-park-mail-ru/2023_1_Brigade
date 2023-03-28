package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	myMiddleware "project/internal/middleware"

	"project/internal/configs"

	log "github.com/sirupsen/logrus"
	httpAuthUser "project/internal/auth/user/delivery/http"
	httpChat "project/internal/chat/delivery/http"
	httpImages "project/internal/images/delivery/http"
	wsMessages "project/internal/messages/delivery/ws"
	httpUser "project/internal/user/delivery/http"

	usecaseAuthSession "project/internal/auth/session/usecase"
	usecaseAuthUser "project/internal/auth/user/usecase"
	usecaseChat "project/internal/chat/usecase"
	usecaseImages "project/internal/images/usecase"
	usecaseMessages "project/internal/messages/usecase"
	usecaseUser "project/internal/user/usecase"

	repositoryAuthSession "project/internal/auth/session/repository"
	repositoryAuthUser "project/internal/auth/user/repository"
	repositoryChat "project/internal/chat/repository"
	repositoryImages "project/internal/images/repository"
	repositoryMessages "project/internal/messages/repository"
	repositoryUser "project/internal/user/repository"
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

	defer func() {
		if err := recover(); err != nil {
			log.WithField("error", err).Error("Panic occurred")
		}
	}()

	yamlPath, exists := os.LookupEnv("YAML_PATH")
	if !exists {
		log.Fatal("Yaml path not found")
	}

	yamlFile, err := os.ReadFile(yamlPath)
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

	endpoint := "172.20.0.1:9000"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4("minio", "minio123", ""),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	userRepository := repositoryUser.NewUserMemoryRepository(db)
	chatRepository := repositoryChat.NewChatMemoryRepository(db)
	imagesRepostiory := repositoryImages.NewImagesMemoryRepository(minioClient)
	messagesRepository := repositoryMessages.NewMessagesMemoryRepository(db)
	authUserRepository := repositoryAuthUser.NewAuthUserMemoryRepository(db)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(redis)

	userUsecase := usecaseUser.NewUserUsecase(userRepository, authUserRepository)
	authUserUsecase := usecaseAuthUser.NewAuthUserUsecase(authUserRepository, userRepository)
	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	chatUsecase := usecaseChat.NewChatUsecase(chatRepository, userRepository)
	messagesUsecase := usecaseMessages.NewMessagesUsecase(messagesRepository)
	imagesUsecase := usecaseImages.NewChatUsecase(imagesRepostiory)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.AllowMethods,
		AllowOrigins:     config.AllowOrigins,
		AllowCredentials: config.AllowCredentials,
		AllowHeaders:     config.AllowHeaders,
	}))
	e.Use(myMiddleware.LoggerMiddleware)
	//e.Use(myMiddleware.XSSMidlleware)
	//e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))
	//curl -X POST \ -H "Content-Type: multipart/form-data" \ -F "image=@/home/marcussss1/Downloads/1avatara_ru_3D001.jpg" \ http://localhost:8081/api/v1/images
	httpUser.NewUserHandler(e, userUsecase)
	httpAuthUser.NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, userUsecase)
	wsMessages.NewMessagesHandler(e, messagesUsecase)
	httpImages.NewImagesHandler(e, imagesUsecase)

	e.Logger.Fatal(e.Start(config.Port))
}
