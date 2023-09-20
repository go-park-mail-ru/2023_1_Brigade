package main

import (
	"github.com/centrifugal/centrifuge-go"
	"os"
	"os/signal"
	authUserRepository "project/internal/microservices/auth/repository"
	usecaseAuth "project/internal/microservices/auth/usecase"
	repositoryChat "project/internal/microservices/chat/repository"
	usecaseChat "project/internal/microservices/chat/usecase"
	repositoryMessages "project/internal/microservices/messages/repository"
	"project/internal/microservices/messages/usecase"
	repositoryUser "project/internal/microservices/user/repository"
	usecaseUser "project/internal/microservices/user/usecase"
	//clientAuth "project/internal/microservices/auth/delivery/grpc/client"
	//clientChat "project/internal/microservices/chat/delivery/grpc/client"
	//clientMessages "project/internal/microservices/messages/delivery/grpc/client"
	//clientUser "project/internal/microservices/user/delivery/grpc/client"
	"project/internal/pkg/serialization"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	defer func() {
		err = db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	centrifugoMessagesClient := centrifuge.NewJsonClient(config.Centrifugo.ConnAddr, centrifuge.Config{})
	centrifugoNotificationsClient := centrifuge.NewJsonClient(config.Centrifugo.ConnAddr, centrifuge.Config{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		centrifugoMessagesClient.Close()
		centrifugoNotificationsClient.Close()
		log.Fatal()
	}()

	err = centrifugoMessagesClient.Connect()
	if err != nil {
		log.Error(err)
	}

	err = centrifugoNotificationsClient.Connect()
	if err != nil {
		log.Error(err)
	}

	subMessages, err := centrifugoMessagesClient.NewSubscription(config.Centrifugo.ChannelName, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Error(err)
	}

	subNotifications, err := centrifugoNotificationsClient.NewSubscription(config.Centrifugo.ChannelName, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Error(err)
	}

	err = subMessages.Subscribe()
	if err != nil {
		log.Error(err)
	}

	err = subNotifications.Subscribe()
	if err != nil {
		log.Error(err)
	}

	imagesRepository := repositoryImages.NewImagesMemoryRepository(userAvatarsClient, chatAvatarsClient, chatImagesClient)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(db)
	authUserRepository := authUserRepository.NewAuthUserMemoryRepository(db)
	chatRepo := repositoryChat.NewChatMemoryRepository(db)
	userRepo := repositoryUser.NewUserMemoryRepository(db)
	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	imagesUsecase := usecaseImages.NewImagesUsecase(imagesRepository)
	authService := usecaseAuth.NewAuthUserUsecase(authUserRepository, userRepo, chatRepo, imagesUsecase)
	chatService := usecaseChat.NewChatUsecase(chatRepo, userRepo, messagesRepo, imagesUsecase)
	userService := usecaseUser.NewUserUsecase(userRepo, authUserRepository)
	messagesService := usecase.NewMessagesUsecase(chatRepo, centrifugoMessagesClient, config.Centrifugo.ChannelName, messagesRepo)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.Cors.AllowMethods,
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowCredentials: config.Cors.AllowCredentials,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    config.Cors.ExposeHeaders,
	}))

	e.Use(myMiddleware.LoggerMiddleware)
	e.Use(myMiddleware.CSRFMiddleware())
	e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))

	e.JSONSerializer = serialization.EasyJsonSerializer{}

	httpUser.NewUserHandler(e, userService)
	httpAuthUser.NewAuthHandler(e, authService, authSessionUsecase, userService)
	httpChat.NewChatHandler(e, chatService, userService)
	httpImages.NewImagesHandler(e, userService, imagesUsecase)

	_, err = wsMessages.NewMessagesHandler(e, messagesService, centrifugoMessagesClient, config.Centrifugo.ChannelName)
	if err != nil {
		log.Error(err)
	}

	_, err = wsNotifications.NewNotificationsHandler(e, chatService, userService, centrifugoNotificationsClient, config.Centrifugo.ChannelName)
	if err != nil {
		log.Error(err)
	}

	e.Logger.Fatal(e.Start(config.Server.Port))
}
