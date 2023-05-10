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
	serverChat "project/internal/chat/delivery/grpc"
	repositoryChat "project/internal/chat/repository"
	usecaseChat "project/internal/chat/usecase"
	"project/internal/configs"
	repositoryImages "project/internal/images/repository"
	usecaseImages "project/internal/images/usecase"
	repositoryMessages "project/internal/messages/repository"
	"project/internal/middleware"
	metrics "project/internal/pkg/metrics/prometheus"
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

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	user_avatars_client, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.UserAvatarsAccessKey, config.VkCloud.UserAvatarsSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Error(err)
	}

	chat_avatars_client, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.ChatAvatarsAccessKey, config.VkCloud.ChatAvatarsSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Error(err)
	}

	chat_images_client, err := minio.New(config.VkCloud.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.VkCloud.ChatImagesAccessKey, config.VkCloud.ChatImagesSecretKey, ""),
		Secure: config.VkCloud.Ssl,
	})
	if err != nil {
		log.Error(err)
	}

	imagesRepository := repositoryImages.NewImagesMemoryRepository(user_avatars_client, chat_avatars_client, chat_images_client)
	chatRepo := repositoryChat.NewChatMemoryRepository(db)
	userRepo := repositoryUser.NewUserMemoryRepository(db, imagesRepository)
	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	imagesUsecase := usecaseImages.NewImagesUsecase(imagesRepository)
	chatUsecase := usecaseChat.NewChatUsecase(chatRepo, userRepo, messagesRepo, imagesUsecase)

	metrics, err := metrics.NewMetricsGRPCServer(config.ChatsService.ServiceName)
	if err != nil {
		log.Error(err)
	}

	grpcMidleware := middleware.NewGRPCMiddleware(metrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcMidleware.GRPCMetricsMiddleware),
	)

	go func() {
		if err = metrics.StartGRPCMetricsServer(config.ChatsService.AddrMetrics); err != nil {
			log.Error(err)
		}
	}()

	service := serverChat.NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	err = service.StartGRPCServer(config.ChatsService.Addr)
	if err != nil {
		log.Error(err)
	}
}
