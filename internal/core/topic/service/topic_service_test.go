package service_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"news-topic-management-service/internal/core/topic/model"
	"news-topic-management-service/internal/core/topic/request"
	"news-topic-management-service/internal/core/topic/service"
	"news-topic-management-service/internal/general/mocks"
	"news-topic-management-service/internal/general/model/common"
	"testing"
)

func TestCreate(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topicID := uuid.New()
	createReq := request.CreateTopicRequest{Name: "New Topic"}
	expectedTopic := &model.Topic{ID: topicID, Name: createReq.Name}

	mockTopicRepo.EXPECT().
		Create(&model.Topic{Name: createReq.Name}).
		Return(expectedTopic, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.Create(createReq)

	assert.NoError(t, err)
	assert.Equal(t, expectedTopic, result)
}

func TestUpdate(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topicID := uuid.New()
	updateReq := request.UpdateTopicRequest{Name: "Updated Topic"}
	existingTopic := &common.Topic{ID: topicID, Name: "Old Topic"}
	updatedTopic := &common.Topic{ID: topicID, Name: updateReq.Name}

	mockTopicRepo.EXPECT().
		Find(topicID).
		Return(existingTopic, nil)
	mockTopicRepo.EXPECT().
		Update(topicID, existingTopic).
		Return(updatedTopic, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.Update(updateReq, topicID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTopic, result)
}

func TestFind(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topicID := uuid.New()
	expectedTopic := &common.Topic{ID: topicID, Name: "Test Topic"}

	mockTopicRepo.EXPECT().
		Find(topicID).
		Return(expectedTopic, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.Find(topicID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTopic, result)
}

func TestDelete(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topicID := uuid.New()
	expectedTopic := &model.Topic{ID: topicID, Name: "Test Topic"}
	expectedTopicFind := &common.Topic{ID: topicID, Name: "Test Topic"}

	mockTopicRepo.EXPECT().
		Find(topicID).
		Return(expectedTopicFind, nil)
	mockTopicRepo.EXPECT().
		Delete(topicID).
		Return(expectedTopic, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.Delete(topicID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTopic, result)
}

func TestGetAll(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	expectedTopics := []model.Topic{
		{ID: uuid.New(), Name: "Topic 1"},
		{ID: uuid.New(), Name: "Topic 2"},
	}

	mockTopicRepo.EXPECT().
		GetAll().
		Return(expectedTopics, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedTopics, result)
}

func TestPreload(t *testing.T) {
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topic := &model.Topic{ID: uuid.New(), Name: "Test Topic"}
	preloadedTopic := &model.Topic{ID: topic.ID, Name: topic.Name + " Preloaded"}

	mockTopicRepo.EXPECT().
		Preload(topic).
		Return(preloadedTopic, nil)

	svc := service.NewTopicService(mockTopicRepo)

	result, err := svc.Preload(topic)

	assert.NoError(t, err)
	assert.Equal(t, preloadedTopic, result)
}
