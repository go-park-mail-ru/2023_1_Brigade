package main

import (
	"github.com/centrifugal/centrifuge-go"
	"os"
	"os/signal"
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

	grpcConnChats, err := grpc.Dial(
		config.ChatsService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer func() {
		err = grpcConnChats.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	grpcConnUsers, err := grpc.Dial(
		config.UsersService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer func() {
		err = grpcConnUsers.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	grpcConnMessages, err := grpc.Dial(
		config.MessagesService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer func() {
		err = grpcConnMessages.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	grpcConnAuth, err := grpc.Dial(
		config.AuthService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer func() {
		err = grpcConnAuth.Close()
		if err != nil {
			log.Error(err)
		}
	}()

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

	DefaultCSRFConfig := middleware.CSRFConfig{
		Skipper:        middleware.DefaultSkipper,
		TokenLength:    32,
		TokenLookup:    "header:" + echo.HeaderXCSRFToken,
		ContextKey:     "csrf",
		CookieName:     "_csrf",
		CookieMaxAge:   86400,
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}
	e.Use(middleware.CSRFWithConfig(DefaultCSRFConfig))

	//e.GET("api/v1/csrf/", func(ctx echo.Context) error {
	//	csrf := ctx.Get("csrf")
	//	//ctx.Response().Header().Set("X-Csrf-Token", csrf.(string))
	//	bytes, err := json.Marshal(csrf)
	//	if err != nil {
	//		return err
	//	}
	//
	//	return ctx.JSONBlob(http.StatusOK, bytes)
	//})

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
