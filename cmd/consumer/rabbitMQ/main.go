package main

import (
	"github.com/centrifugal/centrifuge-go"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	"os/signal"
	"project/internal/config"
	serverConsumer "project/internal/microservices/consumer/delivery/grpc/server"
	"project/internal/middleware"
	"project/internal/monolithic_services/qaas/send_messages/rabbitMQ/consumer/usecase"
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

	centrifugo := centrifuge.NewJsonClient(config.Centrifugo.ConnAddr, centrifuge.Config{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		centrifugo.Close()
		log.Fatal()
	}()

	err = centrifugo.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := centrifugo.NewSubscription(config.Centrifugo.ChannelName, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Error(err)
	}

	err = sub.Subscribe()
	if err != nil {
		log.Error(err)
	}

	consumerUsecase, err := usecase.NewConsumer(config.RabbitMQ.ConnAddr, config.RabbitMQ.QueueName, centrifugo, config.Centrifugo.ChannelName)
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
