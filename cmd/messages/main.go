package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
	"os"
	"project/internal/config"
	repositoryChat "project/internal/microservices/chat/repository"
	consumer "project/internal/microservices/consumer/delivery/grpc/client"
	serverMessages "project/internal/microservices/messages/delivery/grpc/server"
	repositoryMessages "project/internal/microservices/messages/repository"
	usecaseMessages "project/internal/microservices/messages/usecase"
	producer "project/internal/microservices/producer/delivery/grpc/client"
	"project/internal/middleware"
	metrics "project/internal/pkg/metrics/prometheus"
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

	var config config.Config
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

	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	chatRepo := repositoryChat.NewChatMemoryRepository(db)

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
