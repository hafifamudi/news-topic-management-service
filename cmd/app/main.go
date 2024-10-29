package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/hafifamudi/news-topic-management-service/news-topic-management-service/pkg/infrastructure/opentelemetry"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"github.com/joho/godotenv"
	"github.com/riandyrn/otelchi"
	httpSwagger "github.com/swaggo/http-swagger"
	"news-topic-management-service/config"
	_ "news-topic-management-service/docs"
	newsModel "news-topic-management-service/internal/core/news/model"
	newsRoute "news-topic-management-service/internal/core/news/route"
	topicModel "news-topic-management-service/internal/core/topic/model"
	topicRoute "news-topic-management-service/internal/core/topic/route"
	"os"

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

	/** Register routes */
	app := chi.NewRouter()
	app.Use(otelchi.Middleware("news-topic-app"))

	app.Mount("/swagger", httpSwagger.WrapHandler)

	// Create a sub-router for versioning
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
