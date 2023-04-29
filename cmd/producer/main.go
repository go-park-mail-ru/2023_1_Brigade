package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	"project/internal/configs"
	serverProducer "project/internal/qaas/send_messages/producer/delivery/grpc"
	"project/internal/qaas/send_messages/producer/usecase"
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

	grpcServer := grpc.NewServer()

	producerUsecase, err := usecase.NewProducer(config.Kafka.BrokerList)
	if err != nil {
		log.Error(err)
	}

	service := serverProducer.NewProducerServiceGRPCServer(grpcServer, producerUsecase)

	err = service.StartGRPCServer(config.ProducerService.Addr)
	if err != nil {
		log.Error(err)
	}
}
