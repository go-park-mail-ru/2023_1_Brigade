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
	"gopkg.in/yaml.v2"
	"net"
	"os"
	"project/internal/configs"
	wsMessages "project/internal/messages/delivery/ws"
	usecaseMessages "project/internal/messages/usecase"
	myMiddleware "project/internal/middleware"

	log "github.com/sirupsen/logrus"

	grpcChat "project/internal/microservices/chat/grpc"
	protobufChat "project/internal/microservices/chat/protobuf"

	httpAuthUser "project/internal/auth/user/delivery/http"
	httpChat "project/internal/chat/delivery/http"
	httpImages "project/internal/images/delivery/http"
	httpUser "project/internal/user/delivery/http"

	usecaseAuthSession "project/internal/auth/session/usecase"
	usecaseAuthUser "project/internal/auth/user/usecase"
	usecaseChat "project/internal/chat/usecase"
	usecaseImages "project/internal/images/usecase"
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

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB) // ping
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	redis := redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
	})
	defer redis.Close()

	// Warehouse microservice
	//chatsService := grpcChats.NewChatsClient(grpcConnChats)

	userRepository := repositoryUser.NewUserMemoryRepository(db)
	chatRepository := repositoryChat.NewChatMemoryRepository(db)
	imagesRepostiory := repositoryImages.NewImagesMemoryRepository(db)
	messagesRepository := repositoryMessages.NewMessagesMemoryRepository(db)
	authUserRepository := repositoryAuthUser.NewAuthUserMemoryRepository(db)
	authSessionRepository := repositoryAuthSession.NewAuthSessionMemoryRepository(redis)

	/// Microservices
	//grpcConnChats, err := grpc.Dial(
	//	"localhost:8090",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//grpc.WithBlock(),
	//)
	//if err != nil {
	//	log.Fatal("cant connect to grpc ", err)
	//}
	//defer grpcConnChats.Close()
	listen, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal("%s:%s: %s", err)
	}

	microChat := grpcChat.New(chatRepository, userRepository, messagesRepository)
	grpcServer := grpc.NewServer()

	protobufChat.RegisterChatsServer(grpcServer, microChat)

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Fatal("posts: %s", "service image stopped")
		}
	}()

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: %v", err)
	}
	defer conn.Close()

	//microSeviceChat := NewChat
	//ьш
	//client := pb.NewGreeterClient(conn)

	//protobuf.RegisterChatsServer(grpcConnChats)

	userUsecase := usecaseUser.NewUserUsecase(userRepository, authUserRepository)
	authUserUsecase := usecaseAuthUser.NewAuthUserUsecase(authUserRepository, userRepository)
	authSessionUsecase := usecaseAuthSession.NewAuthUserUsecase(authSessionRepository)
	chatUsecase := usecaseChat.NewChatUsecase(chatRepository, userRepository, messagesRepository)
	messagesUsecase := usecaseMessages.NewMessagesUsecase(chatRepository, messagesRepository, config.Kafka)
	imagesUsecase := usecaseImages.NewChatUsecase(imagesRepostiory)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.Cors.AllowMethods,
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowCredentials: config.Cors.AllowCredentials,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    []string{"X-Csrf-Token"},
	}))
	//rpc.Register()
	//csrfMiddleware := csrf.Protect(
	//	[]byte("32-byte-long-auth-key"),
	//	csrf.Secure(true),
	//	csrf.HttpOnly(true),
	//	csrf.RequestHeader("X-CSRF-Token"),
	//)
	//e.Use(echo.WrapMiddleware(csrfMiddleware))
	//csrfMiddleware := csrf.Protect(
	//	[]byte("32-byte-long-auth-key"),
	//	csrf.Secure(true),
	//	csrf.HttpOnly(false),
	//	csrf.RequestHeader("X-CSRF-Token"),
	//)
	//e.Use(echo.WrapMiddleware(csrfMiddleware))
	//e.GET("/api/v1/csrf", func(c echo.Context) error {
	//	csrfToken := uuid.New().String()
	//	if err != nil {
	//		return err
	//	}
	//	type CSRF struct {
	//		csrf string `json:"csrf"`
	//	}
	//	// выставляем токен в хэдер X-CSRF-Token
	//	c.Response().Header().Set("X-Csrf-Token", csrfToken)
	//	a := CSRF{csrf: csrfToken}
	//	// возвращаем токен в качестве ответа
	//	return c.JSON(http.StatusOK, a)
	//})
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-Csrf-Token",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookiePath:     "/",
	}))
	//csrfMiddleware := csrf.Protect(
	//	[]byte("32-byte-long-auth-key"),
	//	csrf.Secure(false),
	//	csrf.HttpOnly(false),
	//)
	//
	//// создаем новый экземпляр Echo сервера
	//
	//// добавляем middleware для логирования запросов
	//e.Use(middleware.Logger())
	//
	//// добавляем middleware для защиты от CSRF
	//e.Use(echo.MiddlewareFunc(csrfMiddleware))
	//e.Use(myMiddleware.CSRFMiddleware())
	//csrfMiddleware := csrf.Protect(
	//	[]byte("32-byte-long-auth-key"),
	//	csrf.Secure(false),
	//	csrf.HttpOnly(false),
	//)
	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	Skipper: func(c echo.Context) bool {
	//		// пропускаем проверку CSRF-токена для GET-запросов
	//		if c.Request().Method == http.MethodGet {
	//			return true
	//		}
	//		return false
	//	},
	//	CookieName: "_csrf",
	//	//Head:  "X-CSRF-Token",
	//	ContextKey: "csrf",
	//	//FailOnError: true,
	//	//SigningKey:  []byte("32-byte-long-auth-key"),
	//}))
	//csrfMiddleware := csrf.Protect(
	//	[]byte("32-byte-long-auth-key"),
	//	csrf.Secure(false),
	//	csrf.HttpOnly(false),
	//)

	// добавляем middleware для защиты от CSRF
	//e.Use(csrfMiddleware)
	e.Use(myMiddleware.LoggerMiddleware)
	//e.Use(myMiddleware.AuthMiddleware(authSessionUsecase))

	httpUser.NewUserHandler(e, userUsecase)
	httpAuthUser.NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, userUsecase)
	wsMessages.NewMessagesHandler(e, messagesUsecase)
	httpImages.NewImagesHandler(e, userUsecase, imagesUsecase)

	e.Logger.Fatal(e.Start(config.Server.Port))
}

//curl -X POST http://localhost:8081/api/v1/signup -H 'Content-Type: application/json' -H 'X-CSRF-Token: <csrf_token>' -d '{"data": "some data"}'
