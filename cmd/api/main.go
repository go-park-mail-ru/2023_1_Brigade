package main

import (
	"os"
	clientAuth "project/internal/microservices/auth/delivery/grpc/client"
	clientChat "project/internal/microservices/chat/delivery/grpc/client"
	clientMessages "project/internal/microservices/messages/delivery/grpc/client"
	clientUser "project/internal/microservices/user/delivery/grpc/client"
	"project/internal/pkg/serialization"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	httpUser "project/internal/microservices/user/delivery/http"

	"project/internal/config"
	wsMessages "project/internal/microservices/messages/delivery/ws"
	myMiddleware "project/internal/middleware"

	log "github.com/sirupsen/logrus"

	httpAuthUser "project/internal/microservices/auth/delivery/http"
	httpChat "project/internal/microservices/chat/delivery/http"
	httpImages "project/internal/monolithic_services/images/delivery/http"
	usecaseImages "project/internal/monolithic_services/images/usecase"
	wsNotifications "project/internal/monolithic_services/notifications/delivery/ws"
	usecaseAuthSession "project/internal/monolithic_services/session/usecase"

	repositoryImages "project/internal/monolithic_services/images/repository"
	repositoryAuthSession "project/internal/monolithic_services/session/repository/postgres"
)

func init() {
	envPath := ".env"
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("No .env file found")
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

	var config config.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	userAvatarsClient, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.UserAvatarsAccessKey, config.VkCloud.UserAvatarsSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Fatal(err)
	}

	chatAvatarsClient, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.ChatAvatarsAccessKey, config.VkCloud.ChatAvatarsSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Fatal(err)
	}

	chatImagesClient, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.ChatImagesAccessKey, config.VkCloud.ChatImagesSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	grpcConnChats, err := grpc.Dial(
		config.ChatsService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnChats.Close()

	grpcConnUsers, err := grpc.Dial(
		config.UsersService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnUsers.Close()

	grpcConnMessages, err := grpc.Dial(
		config.MessagesService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnMessages.Close()

	grpcConnAuth, err := grpc.Dial(
		config.AuthService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnAuth.Close()

	authService := clientAuth.NewAuthUserServiceGRPSClient(grpcConnAuth)
	chatService := clientChat.NewChatServiceGRPSClient(grpcConnChats)
	userService := clientUser.NewUserServiceGRPSClient(grpcConnUsers)
	messagesService := clientMessages.NewMessagesServiceGRPSClient(grpcConnMessages)

	imagesRepository := repositoryImages.NewImagesMemoryRepository(userAvatarsClient, chatAvatarsClient, chatImagesClient)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(db)

	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	imagesUsecase := usecaseImages.NewImagesUsecase(imagesRepository)

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

	p := prometheus.NewPrometheus("echo", nil)
	eProtheus := echo.New()

	e.Use(p.HandlerFunc)
	eProtheus.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	go func() {
		err := eProtheus.Start(":5555")
		if err != nil {
			log.Error(err)
		}
	}()

	e.JSONSerializer = serialization.EasyJsonSerializer{}

	httpUser.NewUserHandler(e, userService)
	httpAuthUser.NewAuthHandler(e, authService, authSessionUsecase, userService)
	httpChat.NewChatHandler(e, chatService, userService)
	wsMessages.NewMessagesHandler(e, messagesService, config.Centrifugo)
	httpImages.NewImagesHandler(e, userService, imagesUsecase)
	wsNotifications.NewNotificationsHandler(e, chatService, userService, config.Centrifugo)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
