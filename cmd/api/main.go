package main

import (
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
	envPath := "../../.env"
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

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

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB) // ping
	if err != nil {
		log.Fatal(err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
	})

	minioClient, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(config.Minio.Username, config.Minio.Password, config.Minio.Token),
	})

	if err != nil {
		log.Fatal(err)
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
	messagesUsecase := usecaseMessages.NewMessagesUsecase(messagesRepository, config.Kafka)
	imagesUsecase := usecaseImages.NewChatUsecase(imagesRepostiory)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.Cors.AllowMethods,
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowCredentials: config.Cors.AllowCredentials,
		AllowHeaders:     config.Cors.AllowHeaders,
	}))
	e.Use(myMiddleware.LoggerMiddleware)
	//e.Use(myMiddleware.XSSMidlleware)
	//e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))

	httpUser.NewUserHandler(e, userUsecase)
	httpAuthUser.NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, userUsecase)
	wsMessages.NewMessagesHandler(e, messagesUsecase)
	httpImages.NewImagesHandler(e, imagesUsecase)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
