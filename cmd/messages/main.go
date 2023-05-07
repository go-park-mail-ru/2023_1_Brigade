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
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
	"os"
	repositoryChat "project/internal/chat/repository"
	"project/internal/clients/consumer"
	"project/internal/clients/producer"
	"project/internal/configs"
	repositoryImages "project/internal/images/repository"
	serverMessages "project/internal/messages/delivery/grpc"
	repositoryMessages "project/internal/messages/repository"
	usecaseMessages "project/internal/messages/usecase"
	"project/internal/middleware"
	metrics "project/internal/pkg/metrics/prometheus"
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

	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)
	imagesRepository := repositoryImages.NewImagesMemoryRepository(user_avatars_client, chat_avatars_client, chat_images_client)

	chatRepo := repositoryChat.NewChatMemoryRepository(db, imagesRepository)

	grpcConnConsumer, err := grpc.Dial(
		config.ConsumerService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnConsumer.Close()

	grpcConnProducer, err := grpc.Dial(
		config.ProducerService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnProducer.Close()

	consumerService := consumer.NewConsumerServiceGRPCClient(grpcConnConsumer)
	producerService := producer.NewProducerServiceGRPCClient(grpcConnProducer)

	messagesUsecase := usecaseMessages.NewMessagesUsecase(chatRepo, consumerService, producerService, messagesRepo)

	metrics, err := metrics.NewMetricsGRPCServer(config.MessagesService.ServiceName)
	if err != nil {
		log.Error(err)
	}

	grpcMidleware := middleware.NewGRPCMiddleware(metrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcMidleware.GRPCMetricsMiddleware),
	)

	go func() {
		if err = metrics.StartGRPCMetricsServer(config.MessagesService.AddrMetrics); err != nil {
			log.Error(err)
		}
	}()

	messagesService := serverMessages.NewMessagesServiceGRPCServer(grpcServer, messagesUsecase)

	err = messagesService.StartGRPCServer(config.MessagesService.Addr)
	if err != nil {
		log.Fatal(err)
	}
}
