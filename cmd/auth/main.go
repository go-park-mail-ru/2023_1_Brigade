package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	"project/internal/config"
	serverAuthUser "project/internal/microservices/auth/delivery/grpc/server"
	authUserRepository "project/internal/microservices/auth/repository"
	authUserUsecase "project/internal/microservices/auth/usecase"
	chatRepository "project/internal/microservices/chat/repository"
	userRepository "project/internal/microservices/user/repository"
	"project/internal/middleware"
	repositoryImages "project/internal/monolithic_services/images/repository"
	usecaseImages "project/internal/monolithic_services/images/usecase"
	authSessionRepository "project/internal/monolithic_services/session/repository/postgres"
	authSessionUsecase "project/internal/monolithic_services/session/usecase"
	metrics "project/internal/pkg/metrics/prometheus"
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

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)

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

	imagesRepository := repositoryImages.NewImagesMemoryRepository(userAvatarsClient, chatAvatarsClient, chatImagesClient)
	userRepository := userRepository.NewUserMemoryRepository(db)
	authUserRepository := authUserRepository.NewAuthUserMemoryRepository(db)
	authSessionRepository := authSessionRepository.NewAuthSessionMemoryRepository(db)
	chatRepository := chatRepository.NewChatMemoryRepository(db)

	imagesUsecase := usecaseImages.NewImagesUsecase(imagesRepository)
	authSessionUsecase := authSessionUsecase.NewAuthUserUsecase(authSessionRepository)
	authUserUsecase := authUserUsecase.NewAuthUserUsecase(authUserRepository, userRepository, chatRepository, imagesUsecase)

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
		log.Fatal(err)
	}
}
