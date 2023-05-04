package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	"project/internal/configs"
	"project/internal/middleware"
	metrics "project/internal/pkg/metrics/prometheus"
	serverConsumer "project/internal/qaas/send_messages/consumer/delivery/grpc"
	"project/internal/qaas/send_messages/consumer/usecase"
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

	consumerUsecase, err := usecase.NewConsumer(config.Kafka.BrokerList, config.Kafka.GroupID)
	if err != nil {
		log.Fatal(err)
	}

	metrics, err := metrics.NewMetricsGRPCServer(config.ConsumerService.ServiceName)
	if err != nil {
		log.Error(err)
	}

	grpcMidleware := middleware.NewGRPCMiddleware(metrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcMidleware.GRPCMetricsMiddleware),
	)

	go func() {
		if err = metrics.StartGRPCMetricsServer(config.ConsumerService.AddrMetrics); err != nil {
			log.Error(err)
		}
	}()

	service := serverConsumer.NewConsumerServiceGRPCServer(grpcServer, consumerUsecase)

	err = service.StartGRPCServer(config.ConsumerService.Addr)
	if err != nil {
		log.Fatal(err)
	}
}
