package chat

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"os"
	"project/internal/chat/usecase"
	"project/internal/configs"

	repositoryChat "project/internal/chat/repository"
	repositoryMessages "project/internal/messages/repository"
	repositoryUser "project/internal/user/repository"
)

func init() {
	envPath := "../../.env"
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	//// Config
	//var configPath string
	//
	//flag.StringVar(&configPath, "config-path", "cmd/auth/configs/debug.toml", "path to config file")
	//
	//flag.Parse()
	//
	//config := innerPKG.NewConfig()
	//
	//_, err := toml.DecodeFile(configPath, config)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//
	//// Logger
	//logger, closeResource := pkg.NewLogger(&config.Logger)
	//defer func(closer func() error, log *logrus.Logger) {
	//	err = closer()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(closeResource, logger)
	//
	//metrics := monitoring.NewPrometheusMetrics(config.ServerGRPCAuth.ServiceName)
	//err = metrics.SetupMonitoring()
	//if err != nil {
	//	logger.Fatal(err)
	//}

	// Middleware
	//md := middleware.NewGRPCMiddleware(logger, metrics)

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

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB) // ping
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	// Connections
	//postgres := sqltools.NewPostgresRepository(&config.DatabaseParams)

	// Auth repository
	//authStorage := repoAuth.NewAuthDatabase(postgres)
	//sessionStorage := repoSession.NewSessionCache()

	// Auth service
	//authService := serviceSession.NewAuthService(authStorage)
	//sessionService := serviceSession.NewSessionService(sessionStorage)

	chatRepo := repositoryChat.NewChatMemoryRepository(db)
	userRepo := repositoryUser.NewUserMemoryRepository(db)
	messagesRepo := repositoryMessages.NewMessagesMemoryRepository(db)

	chatUsecase := usecase.NewChatUsecase(chatRepo, userRepo, messagesRepo)

	//go monitoring.CreateNewMonitoringServer(config.Metrics.BindHTTPAddr)

	// Server
	grpcServer := grpc.NewServer()

	service := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)
	//New

	//service := server.New(grpcServer, authService, sessionService)

	logrus.Info(config.ServerGRPCAuth.ServiceName + " starting server at " + config.ServerGRPCAuth.BindAddr)

	err = service.StartGRPCServer(config.ServerGRPCAuth.BindAddr)
	if err != nil {
		logrus.Fatal(err)
	}
}
