package repository_test

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/model"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) (*gorm.DB, repository.NewsRepository) {
	// Create an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate your models
	err = db.AutoMigrate(&model.News{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewNewsRepository(db)
	return db, repo
}

func TestNewsRepository_Create(t *testing.T) {
	_, repo := setupTestDB(t)

	newsID := uuid.New()
	news := &model.News{ID: newsID, Title: "Test News", Content: "News Content"}

	// Create the news item in the database
	result, err := repo.Create(news)
	assert.NoError(t, err)
	assert.Equal(t, news, result)
}

func TestNewsRepository_Delete(t *testing.T) {
	_, repo := setupTestDB(t)

	newsID := uuid.New()
	news := &model.News{ID: newsID}
	// Create the news item first
	_, err := repo.Create(news)
	if err != nil {
		t.Fatalf("failed to create news item: %v", err)
	}

	// Delete the news item
	result, err := repo.Delete(newsID)
	assert.NoError(t, err)
	assert.Equal(t, news.ID, result.ID)
}

func TestNewsRepository_Find(t *testing.T) {
	_, repo := setupTestDB(t)

	newsID := uuid.New()
	news := &model.News{ID: newsID, Title: "Test News", Content: "News Content"}

	resultCreate, err := repo.Create(news)
	if err != nil {
		t.Fatalf("failed to create news item: %v", err)
	}

	result, err := repo.Find(newsID)
	assert.NoError(t, err)
	assert.Equal(t, resultCreate.Title, result.Title)
}

func TestNewsRepository_GetAll(t *testing.T) {
	_, repo := setupTestDB(t)

	newsList := []model.News{
		{ID: uuid.New(), Title: "News 1", Content: "Content 1"},
		{ID: uuid.New(), Title: "News 2", Content: "Content 2"},
	}

	// Create the news items first
	for _, news := range newsList {
		_, err := repo.Create(&news)
		if err != nil {
			t.Fatalf("failed to create news item: %v", err)
		}
	}

	// Get all news items
	result, err := repo.GetAll()
	assert.NoError(t, err)

	// Helper function to compare relevant fields
	compareNews := func(a, b model.News) bool {
		return a.ID == b.ID && a.Title == b.Title && a.Content == b.Content
	}

	// Create maps for comparison
	expectedMap := make(map[uuid.UUID]model.News)
	for _, news := range newsList {
		expectedMap[news.ID] = news
	}

	resultMap := make(map[uuid.UUID]model.News)
	for _, news := range result {
		resultMap[news.ID] = news
	}

	// Check if all expected items are in the result
	for _, expected := range newsList {
		actual, exists := resultMap[expected.ID]
		if !exists || !compareNews(expected, actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}

	// Check if all result items are in the expected
	for _, actual := range result {
		expected, exists := expectedMap[actual.ID]
		if !exists || !compareNews(expected, actual) {
			t.Errorf("unexpected result %v", actual)
		}
	}
}

func TestNewsRepository_Update(t *testing.T) {
	_, repo := setupTestDB(t)

	newsID := uuid.New()
	originalNews := &model.News{ID: newsID, Title: "Original News", Content: "Original Content"}
	updatedNews := &model.News{ID: newsID, Title: "Updated News", Content: "Updated Content"}

	// Create the news item first
	_, err := repo.Create(originalNews)
	if err != nil {
		t.Fatalf("failed to create news item: %v", err)
	}

	// Update the news item
	result, err := repo.Update(newsID, updatedNews)
	assert.NoError(t, err)
	assert.Equal(t, updatedNews, result)
}

func TestNewsRepository_Preload(t *testing.T) {
	_, repo := setupTestDB(t)

	news := &model.News{ID: uuid.New(), Title: "Preloaded News", Content: "Preloaded Content"}
	preloadedNews := &model.News{ID: news.ID, Title: news.Title + " - Preloaded", Content: news.Content}

	// Create the news item first
	_, err := repo.Create(news)
	if err != nil {
		t.Fatalf("failed to create news item: %v", err)
	}

	// Preload the news item
	result, err := repo.Preload(news)
	assert.NoError(t, err)
	assert.Equal(t, preloadedNews.ID, result.ID)
}
