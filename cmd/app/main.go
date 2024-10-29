package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/opentelemetry"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/riandyrn/otelchi"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yarlson/chiprom"
	"news-topic-management-service/config"
	_ "news-topic-management-service/docs"
	newsModel "news-topic-management-service/internal/core/news/model"
	newsRoute "news-topic-management-service/internal/core/news/route"
	topicModel "news-topic-management-service/internal/core/topic/model"
	topicRoute "news-topic-management-service/internal/core/topic/route"
	"os"
	"path/filepath"

	"log"
	"net/http"
)

// @title News Topic Management API
// @version 1.0
// @description This News Topic Management service API.

// @contact.name Hafif Nur Muhammad
// @contact.url https://hafifamudi.github.io/
// @contact.email hafifcyber@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1/api
func main() {
	/**  Specify and Load environment variables */
	var cfg config.Config
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "DEVELOPMENT_DOCKER" {
		/**  Load configuration */
		cfg = config.Config{
			App: config.App{
				Name: os.Getenv("APP_NAME"),
				Env:  os.Getenv("APP_ENV"),
			},
			DB: config.DBConfig{
				Client:   os.Getenv("DB_CLIENT"),
				Host:     os.Getenv("DB_HOST"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				Port:     os.Getenv("DB_PORT"),
				Database: os.Getenv("DB_DATABASE"),
			},
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		/**  Load configuration */
		cfg = config.Instance()
	}

	/** Close database if application exit */
	defer db.CloseDB()

	/**  Initialize database */
	dbInstance, err := db.InitDB(db.Config{
		Client:   cfg.DB.Client,
		Database: cfg.DB.Database,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
	})

	if err != nil {
		panic("failed to connect database")
	}

	/** Migrate the schema */
	err = dbInstance.AutoMigrate(
		&topicModel.Topic{},
		&newsModel.News{},
	)
	if err != nil {
		return
	}

	/** Start Software Instrumentation */
	cleanup, err := opentelemetry.InitOpenTelemetry()
	if err != nil {
		log.Fatalf("failed to initialize OpenTelemetry: %v", err)
	}
	defer cleanup(context.Background())

	/** Init Logger */
	baseDir, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("failed to get current working directory: %v", err)
	}

	logPath := filepath.Join(baseDir, "logs", "app.log")
	if err != nil {
		logrus.Fatalf("failed to open log file: %v", err)
	}

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("failed to open log file: %v", err)
	}

	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	/** Register routes */
	app := chi.NewRouter()
	app.Use(otelchi.Middleware("news-topic-app", otelchi.WithChiRoutes(app)))
	app.Use(chiprom.NewMiddleware("news-topic-app"))

	app.Mount("/swagger", httpSwagger.WrapHandler)

	/** Expose Metrics based on promHTTP */
	app.Handle("/metrics", promhttp.Handler())

	/** Create a sub-router for versioning */
	v1Router := chi.NewRouter()
	topicRoute.Register(v1Router)
	newsRoute.Register(v1Router)

	/** Mount the versioned router */
	app.Mount("/v1/api", v1Router)

	log.Println("Server is starting on port 3333")

	err = http.ListenAndServe(":3333", app)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
