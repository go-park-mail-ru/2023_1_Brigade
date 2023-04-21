package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
	"os"
	clientChat "project/internal/clients/chat"
	clientUser "project/internal/clients/user"

	"project/internal/configs"
	wsMessages "project/internal/messages/delivery/ws"
	usecaseMessages "project/internal/messages/usecase"
	myMiddleware "project/internal/middleware"

	log "github.com/sirupsen/logrus"

	httpAuthUser "project/internal/auth/user/delivery/http"
	httpChat "project/internal/chat/delivery/http"
	httpImages "project/internal/images/delivery/http"
	httpUser "project/internal/user/delivery/http"

	usecaseAuthSession "project/internal/auth/session/usecase"
	usecaseAuthUser "project/internal/auth/user/usecase"
	usecaseImages "project/internal/images/usecase"

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
		log.Error("Yaml path not found")
	}

	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Error(err)
	}

	var config configs.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Error(err)
	}

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	redis := redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
	})
	defer redis.Close()

	grpcConnChats, err := grpc.Dial(
		config.ChatsService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Error("cant connect to grpc ", err)
	}
	defer grpcConnChats.Close()

	grpcConnUsers, err := grpc.Dial(
		config.UsersService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Error("cant connect to grpc ", err)
	}
	defer grpcConnUsers.Close()

	chatService := clientChat.NewChatServiceGRPSClient(grpcConnChats)
	userService := clientUser.NewUserServiceGRPSClient(grpcConnUsers)

	userRepository := repositoryUser.NewUserMemoryRepository(db)
	chatRepository := repositoryChat.NewChatMemoryRepository(db)
	imagesRepostiory := repositoryImages.NewImagesMemoryRepository(db)
	messagesRepository := repositoryMessages.NewMessagesMemoryRepository(db)
	authUserRepository := repositoryAuthUser.NewAuthUserMemoryRepository(db)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(redis)

	authUserUsecase := usecaseAuthUser.NewAuthUserUsecase(authUserRepository, userRepository)
	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	messagesUsecase := usecaseMessages.NewMessagesUsecase(chatRepository, messagesRepository, config.Kafka)
	imagesUsecase := usecaseImages.NewChatUsecase(imagesRepostiory)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.Cors.AllowMethods,
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowCredentials: config.Cors.AllowCredentials,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    config.Cors.ExposeHeaders,
	}))

	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup:    "header:X-Csrf-Token",
	//	CookieSecure:   true,
	//	CookieHTTPOnly: true,
	//	CookiePath:     "/",
	//}))

	e.Use(myMiddleware.LoggerMiddleware)
	e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))

	httpUser.NewUserHandler(e, userService)
	httpAuthUser.NewAuthHandler(e, authUserUsecase, authSessionUsecase, userService)
	httpChat.NewChatHandler(e, chatService, userService)
	wsMessages.NewMessagesHandler(e, messagesUsecase)
	httpImages.NewImagesHandler(e, userService, imagesUsecase)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
