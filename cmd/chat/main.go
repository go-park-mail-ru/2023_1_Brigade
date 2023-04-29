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
	serverChat "project/internal/chat/delivery/grpc"
	repositoryChat "project/internal/chat/repository"
	usecaseChat "project/internal/chat/usecase"
	"project/internal/configs"
	repositoryMessages "project/internal/messages/repository"
	repositoryUser "project/internal/user/repository"
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

	chatRepo := repositoryChat.NewChatMemoryRepository(db)
	userRepo := repositoryUser.NewUserMemoryRepository(db)
	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	chatUsecase := usecaseChat.NewChatUsecase(chatRepo, userRepo, messagesRepo)

	grpcServer := grpc.NewServer()

	service := serverChat.NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	err = service.StartGRPCServer(config.ChatsService.Addr)
	if err != nil {
		log.Error(err)
	}
}
