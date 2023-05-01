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
	repositoryChat "project/internal/chat/repository"
	"project/internal/clients/consumer"
	"project/internal/clients/producer"
	"project/internal/configs"
	serverMessages "project/internal/messages/delivery/grpc"
	repositoryMessages "project/internal/messages/repository"
	usecaseMessages "project/internal/messages/usecase"
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

	grpcServer := grpc.NewServer()

	chatRepo := repositoryChat.NewChatMemoryRepository(db)
	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	grpcConnConsumer, err := grpc.Dial(
		config.ConsumerService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnConsumer.Close()

	grpcConnProducer, err := grpc.Dial(
		config.ProducerService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnProducer.Close()

	consumerService := consumer.NewConsumerServiceGRPCClient(grpcConnConsumer)
	producerService := producer.NewProducerServiceGRPCClient(grpcConnProducer)

	//producerService, err := usecase2.NewProducer(config.Kafka.BrokerList)
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//consumerService, err := usecase.NewConsumer(config.Kafka.BrokerList, config.Kafka.GroupID)
	//if err != nil {
	//	log.Error(err)
	//}

	messagesUsecase := usecaseMessages.NewMessagesUsecase(chatRepo, messagesRepo, consumerService, producerService)

	messagesService := serverMessages.NewMessagesServiceGRPCServer(grpcServer, messagesUsecase)

	err = messagesService.StartGRPCServer(config.MessagesService.Addr)
	if err != nil {
		log.Fatal(err)
	}
}
