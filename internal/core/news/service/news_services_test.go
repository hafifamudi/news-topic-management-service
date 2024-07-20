package service_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	newsModel "news-topic-management-service/internal/core/news/model"
	"news-topic-management-service/internal/core/news/request"
	"news-topic-management-service/internal/core/news/service"
	"news-topic-management-service/internal/general/mocks"
	"news-topic-management-service/internal/general/model/common"
	"testing"
)

func TestCreate(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	topicID := uuid.New()
	createReq := request.CreateNewsRequest{
		Title:    "New News",
		Content:  "Content of the new news",
		Status:   "Draft",
		TopicIDs: []string{topicID.String()},
	}
	topic := &common.Topic{ID: topicID}
	expectedNews := &newsModel.News{
		Title:   createReq.Title,
		Content: createReq.Content,
		Status:  createReq.Status,
		Topics:  []common.Topic{*topic},
	}

	mockTopicRepo.EXPECT().
		Find(topicID).
		Return(topic, nil)

	mockNewsRepo.EXPECT().
		Create(&newsModel.News{
			Title:   createReq.Title,
			Content: createReq.Content,
			Status:  createReq.Status,
			Topics:  []common.Topic{*topic},
		}).
		Return(expectedNews, nil)

	svc := service.NewNewsService(mockNewsRepo, mockTopicRepo)

	result, err := svc.Create(createReq)

	assert.NoError(t, err)
	assert.Equal(t, expectedNews, result)
}

func TestUpdate(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)
	mockTopicRepo := mocks.NewMockTopicRepository(t)

	newsID := uuid.New()
	topicID := uuid.New()
	updateReq := request.UpdateNewsRequest{
		Title:    "Updated News",
		Content:  "Updated content",
		Status:   "Published",
		TopicIDs: []string{topicID.String()},
	}
	existingNews := &newsModel.News{ID: newsID, Title: "Old News", Content: "Old content", Status: "Draft", Topics: []common.Topic{}}
	updatedNews := &newsModel.News{
		ID:      newsID,
		Title:   updateReq.Title,
		Content: updateReq.Content,
		Status:  updateReq.Status,
		Topics:  []common.Topic{*&common.Topic{ID: topicID}},
	}

	mockNewsRepo.EXPECT().
		Find(newsID).
		Return(existingNews, nil)

	mockTopicRepo.EXPECT().
		Find(topicID).
		Return(&common.Topic{ID: topicID}, nil)

	mockNewsRepo.EXPECT().
		Update(newsID, updatedNews).
		Return(updatedNews, nil)

	svc := service.NewNewsService(mockNewsRepo, mockTopicRepo)

	result, err := svc.Update(updateReq, newsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedNews, result)
}

func TestFind(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)

	newsID := uuid.New()
	expectedNews := &newsModel.News{ID: newsID, Title: "Test News", Content: "News content", Status: "Published"}

	mockNewsRepo.EXPECT().
		Find(newsID).
		Return(expectedNews, nil)

	svc := service.NewNewsService(mockNewsRepo, nil)

	result, err := svc.Find(newsID)

	assert.NoError(t, err)
	assert.Equal(t, expectedNews, result)
}

func TestDelete(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)

	newsID := uuid.New()
	expectedNews := &newsModel.News{ID: newsID, Title: "Test News", Content: "News content", Status: "Published"}

	mockNewsRepo.EXPECT().
		Find(newsID).
		Return(expectedNews, nil)
	mockNewsRepo.EXPECT().
		Delete(newsID).
		Return(expectedNews, nil)

	svc := service.NewNewsService(mockNewsRepo, nil)

	result, err := svc.Delete(newsID)

	assert.NoError(t, err)
	assert.Equal(t, expectedNews, result)
}

func TestGetAll(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)

	expectedNewsList := []newsModel.News{
		{ID: uuid.New(), Title: "News 1", Content: "Content 1", Status: "Published"},
		{ID: uuid.New(), Title: "News 2", Content: "Content 2", Status: "Draft"},
	}

	mockNewsRepo.EXPECT().
		GetAll(nil, nil).
		Return(expectedNewsList, nil)

	// Create the NewsService with the mocks
	svc := service.NewNewsService(mockNewsRepo, nil)

	result, err := svc.GetAll(nil, nil)

	assert.NoError(t, err)
	assert.Equal(t, expectedNewsList, result)
}

func TestPreload(t *testing.T) {
	mockNewsRepo := mocks.NewMockNewsRepository(t)

	news := &newsModel.News{ID: uuid.New(), Title: "Test News", Content: "News content", Status: "Published"}
	preloadedNews := &newsModel.News{
		ID:      news.ID,
		Title:   news.Title,
		Content: news.Content,
		Status:  news.Status,
		Topics:  []common.Topic{{ID: uuid.New()}}, // Example preloaded topics
	}

	mockNewsRepo.EXPECT().
		Preload(news).
		Return(preloadedNews, nil)

	svc := service.NewNewsService(mockNewsRepo, nil)

	result, err := svc.Preload(news)

	assert.NoError(t, err)
	assert.Equal(t, preloadedNews, result)
}
