package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"news-topic-management-service/config"
	newsModel "news-topic-management-service/internal/core/news/model"
	topicModel "news-topic-management-service/internal/core/topic/model"
	"news-topic-management-service/internal/general/model/common"
	"news-topic-management-service/pkg/infrastructure/db"
)

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
	db, err := db.InitDB(db.Config{
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
	db.AutoMigrate(
		&topicModel.Topic{},
		&newsModel.News{},
	)

	topic := []common.Topic{
		{
			ID:   uuid.New(),
			Name: "Tech",
		},
	}
	if err := db.Create(topic).Error; err != nil {
		log.Fatalf("failed to seed topic data: %v", err)
	}

	news := &newsModel.News{
		ID:      uuid.New(),
		Title:   "Hacking News",
		Content: "Hacking incident that happened in a country",
		Status:  "Draft",
		Topics:  topic,
	}
	if err := db.Create(news).Error; err != nil {
		log.Fatalf("failed to seed news data: %v", err)
	}

	log.Println("Successfully seed the news and topic data")
}
