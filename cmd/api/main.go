package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"project/cmd/configs"
	httpauth "project/internal/auth/delivery/http"
	authrepository "project/internal/auth/repository"
	authusecase "project/internal/auth/usecase"
	httpuser "project/internal/user/delivery/http"
	userrepository "project/internal/user/repository"
	userusecase "project/internal/user/usecase"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	yamlPath := "config.yaml"
	yamlFile, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		log.Fatal(err)
	}

	var config configs.Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(config.DB, config.ConnectionToDB)

	if err != nil {
		log.Fatal(err)
	}

	repositoryAuth := authrepository.NewAuthMemoryRepository(db)
	repositoryUser := userrepository.NewUserMemoryRepository(db)

	usecaseAuth := authusecase.NewAuthUsecase(repositoryAuth)
	usecaseUser := userusecase.NewUserUsecase(repositoryUser)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     config.AllowMethods,
		AllowOrigins:     config.AllowOrigins,
		AllowCredentials: config.AllowCredentials,
		AllowHeaders:     config.AllowHeaders,
	}))

	httpauth.NewAuthHandler(e, usecaseAuth)
	httpuser.NewUserHandler(e, usecaseUser)

	e.Logger.Fatal(e.Start(config.Port))
}
