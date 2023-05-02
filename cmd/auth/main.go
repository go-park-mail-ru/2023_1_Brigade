package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	"os"
	authSessionRepository "project/internal/auth/session/repository/postgres"
	authSessionUsecase "project/internal/auth/session/usecase"
	serverAuthUser "project/internal/auth/user/delivery/grpc"
	authUserRepository "project/internal/auth/user/repository"
	authUserUsecase "project/internal/auth/user/usecase"
	"project/internal/configs"
	prometheus "project/internal/pkg/metrics"
	userRepository "project/internal/user/repository"
	"time"
)

func init() {
	envPath := ".env"
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found")
	}
}

type GRPCMiddleware struct {
	//cfg    *config.Config
	metric *prometheus.MetricsGRPC
}

func NewGRPCMiddleware(m *prometheus.MetricsGRPC) *GRPCMiddleware {
	return &GRPCMiddleware{m}
}

func (m *GRPCMiddleware) MetricsGRPCUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, uHandler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := uHandler(ctx, req)

	errStatus, _ := status.FromError(err)
	code := errStatus.Code()

	if code != codes.OK {
		m.metric.Errors.WithLabelValues(code.String(), info.FullMethod).Inc()
	}

	m.metric.Timings.WithLabelValues(code.String(), info.FullMethod).Observe(time.Since(start).Seconds())
	m.metric.Hits.Inc()

	return resp, err
}

func (m *GRPCMiddleware) LoggerGRPCUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, uHandler grpc.UnaryHandler) (interface{}, error) {
	//grpcLogger := log.WithField(logrus.Fields{
	//"method": info.FullMethod,
	//})
	grpcLogger := log.WithField("method", info.FullMethod)

	ctx = context.WithValue(ctx, "handler-logger-ctx", grpcLogger)
	resp, err := uHandler(ctx, req)

	return resp, err
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

	//grpcServer := grpc.NewServer()

	metrics, err := prometheus.NewMetricsGRPCServer("auth")
	if err != nil {
		log.Fatal("auth - failed create metrics server", err)
	}
	middleware := NewGRPCMiddleware(metrics)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.LoggerGRPCUnaryInterceptor, middleware.MetricsGRPCUnaryInterceptor),
	)
	//authGRPCServer := _authServer.NewAuthServerGRPC(grpcServer, authUC)

	go func() {
		if err = metrics.RunGRPCMetricsServer(":" + "9005"); err != nil {
			log.Fatal("auth - failed run metrics server", err)
		}
	}()

	service := serverAuthUser.NewAuthUserServiceGRPCServer(grpcServer, authUserUsecase, authSessionUsecase)

	err = service.StartGRPCServer(config.AuthService.Addr)
	if err != nil {
		log.Error(err)
	}
}
