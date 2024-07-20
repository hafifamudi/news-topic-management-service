package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"news-topic-management-service/config"
	_ "news-topic-management-service/docs"
	newsModel "news-topic-management-service/internal/core/news/model"
	newsRoute "news-topic-management-service/internal/core/news/route"
	topicModel "news-topic-management-service/internal/core/topic/model"
	topicRoute "news-topic-management-service/internal/core/topic/route"

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
	/**  Load environment variables */
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**  Load configuration */
	cfg := config.Instance()

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
	dbInstance.AutoMigrate(
		&topicModel.Topic{},
		&newsModel.News{},
	)

	/** Register routes */
	app := chi.NewRouter()

	app.Mount("/swagger", httpSwagger.WrapHandler)

	// Create a sub-router for versioning
	v1Router := chi.NewRouter()
	topicRoute.Register(v1Router)
	newsRoute.Register(v1Router)

	// Mount the versioned router
	app.Mount("/v1/api", v1Router)

	log.Println("Server is starting on port 3333")

	err = http.ListenAndServe(":3333", app)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
