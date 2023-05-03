package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	authSessionRepository "project/internal/auth/session/repository/postgres"
	authSessionUsecase "project/internal/auth/session/usecase"
	serverAuthUser "project/internal/auth/user/delivery/grpc"
	authUserRepository "project/internal/auth/user/repository"
	authUserUsecase "project/internal/auth/user/usecase"
	"project/internal/configs"
	"project/internal/middleware"
	metrics "project/internal/pkg/metrics/prometheus"
	userRepository "project/internal/user/repository"
)

func init() {
	envPath := ".env"
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

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	userRepository := userRepository.NewUserMemoryRepository(db)
	authUserRepository := authUserRepository.NewAuthUserMemoryRepository(db)
	authSessionRepository := authSessionRepository.NewAuthSessionMemoryRepository(db)

	authUserUsecase := authUserUsecase.NewAuthUserUsecase(authUserRepository, userRepository)
	authSessionUsecase := authSessionUsecase.NewAuthUserUsecase(authSessionRepository)

	metrics, err := metrics.NewMetricsGRPCServer(config.AuthService.ServiceName)
	if err != nil {
		log.Error(err)
	}

	grpcMidleware := middleware.NewGRPCMiddleware(metrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcMidleware.GRPCMetricsMiddleware),
	)

	go func() {
		if err = metrics.StartGRPCMetricsServer(config.AuthService.AddrMetrics); err != nil {
			log.Error(err)
		}
	}()

	service := serverAuthUser.NewAuthUserServiceGRPCServer(grpcServer, authUserUsecase, authSessionUsecase)

	err = service.StartGRPCServer(config.AuthService.Addr)
	if err != nil {
		log.Error(err)
	}
}
