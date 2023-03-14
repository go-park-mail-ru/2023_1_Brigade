package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	httpauth "project/internal/auth/delivery/http"
	authrepository "project/internal/auth/repository"
	authusecase "project/internal/auth/usecase"
	"project/internal/configs"
	myMiddleware "project/internal/middleware"
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

	repositoryAuth := authrepository.NewAuthMemoryRepository(db)
	usecaseAuth := authusecase.NewAuthUsecase(repositoryAuth)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(myMiddleware.HandlerMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.AllowMethods,
		AllowOrigins:     config.AllowOrigins,
		AllowCredentials: config.AllowCredentials,
		AllowHeaders:     config.AllowHeaders,
	}))

	httpauth.NewAuthHandler(e, usecaseAuth)

	e.Logger.Fatal(e.Start(config.Port))
}
