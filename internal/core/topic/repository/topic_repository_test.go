package repository_test

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/model"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/repository"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) (*gorm.DB, repository.TopicRepository) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate your models
	err = db.AutoMigrate(&model.Topic{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	repo := repository.NewTopicRepository(db)
	return db, repo
}

func TestTopicRepository_Create(t *testing.T) {
	_, repo := setupTestDB(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Test Topic"}
	createdTopic, err := repo.Create(topic)
	assert.NoError(t, err)
	assert.Equal(t, topic.Name, createdTopic.Name)
}

func TestTopicRepository_Delete(t *testing.T) {
	_, repo := setupTestDB(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Test Topic"}
	_, err := repo.Create(topic) // First, create the topic
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	deletedTopic, err := repo.Delete(topic.ID)
	assert.NoError(t, err)
	assert.Equal(t, topic.ID, deletedTopic.ID)

	// Verify that the topic has been deleted
	_, err = repo.Find(topic.ID)
	assert.Error(t, err)
}

func TestTopicRepository_Find(t *testing.T) {
	_, repo := setupTestDB(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Test Topic"}
	_, err := repo.Create(topic) // First, create the topic
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	foundTopic, err := repo.Find(topic.ID)
	assert.NoError(t, err)
	assert.Equal(t, topic.Name, foundTopic.Name)
}

func TestTopicRepository_GetAll(t *testing.T) {
	_, repo := setupTestDB(t)

	// Create topics in the database
	topics := []model.Topic{
		{ID: uuid.New(), Name: "Topic 1"},
		{ID: uuid.New(), Name: "Topic 2"},
	}
	for _, topic := range topics {
		_, err := repo.Create(&topic)
		if err != nil {
			t.Fatalf("failed to create topic: %v", err)
		}
	}

	// Retrieve all topics from the database
	allTopics, err := repo.GetAll()
	assert.NoError(t, err)

	// Check that the number of topics is as expected
	assert.Len(t, allTopics, len(topics))

	// Compare topics by ID and Name only
	expectedTopics := make(map[uuid.UUID]string)
	for _, topic := range topics {
		expectedTopics[topic.ID] = topic.Name
	}

	actualTopics := make(map[uuid.UUID]string)
	for _, topic := range allTopics {
		actualTopics[topic.ID] = topic.Name
	}

	assert.Equal(t, expectedTopics, actualTopics)
}

func TestTopicRepository_Update(t *testing.T) {
	_, repo := setupTestDB(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Old Name"}
	_, err := repo.Create(topic) // First, create the topic
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	topic.Name = "Updated Name"
	updatedTopic, err := repo.Update(topic.ID, (*common.Topic)(topic))
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedTopic.Name)
}

func TestTopicRepository_Preload(t *testing.T) {
	_, repo := setupTestDB(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Test Topic"}
	_, err := repo.Create(topic) // First, create the topic
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	preloadedTopic, err := repo.Preload(topic)
	assert.NoError(t, err)
	assert.Equal(t, topic.ID, preloadedTopic.ID)
}
