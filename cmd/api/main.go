package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	httpAuth "project/internal/auth/delivery/http"
	repositoryAuth "project/internal/auth/repository"
	usecaseAuth "project/internal/auth/usecase"
	httpChat "project/internal/chat/delivery/http"
	repositoryChat "project/internal/chat/repository"
	usecaseChat "project/internal/chat/usecase"
	"project/internal/configs"
	myMiddleware "project/internal/middleware"
	httpUser "project/internal/user/delivery/http"
	repositoryUser "project/internal/user/repository"
	usecaseUser "project/internal/user/usecase"
	"project/internal/s3"
)

func init() {
	if err := godotenv.Load(); err != nil {
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

	imageStorage := s3.NewImageStorage(config.ImagesBucket)

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

	//e.Use(middleware.CSRF())
	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup: "header:X-XSRF-TOKEN",
	//}))
	//e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	//	XSSProtection:         "",
	//	ContentTypeNosniff:    "",
	//	XFrameOptions:         "",
	//	HSTSMaxAge:            3600,
	//	ContentSecurityPolicy: "default-src 'self'",
	//}))

	e.Use(myMiddleware.LoggerMiddleware)
	e.Use(myMiddleware.AuthMiddleware(authUsecase))

	httpUser.NewUserHandler(e, userUsecase)
	httpAuth.NewAuthHandler(e, authUsecase, userUsecase)
	httpChat.NewChatHandler(e, chatUsecase, authUsecase)

	e.Logger.Fatal(e.Start(config.Port))
	e.Logger.Fatal(e.Start(":8081"))
}
