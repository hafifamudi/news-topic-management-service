package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/controller"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/model"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/request"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/resource"
	"github.com/hafifamudi/news-topic-management-service/internal/general/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDetailNews(t *testing.T) {
	mockService := new(mocks.MockNewsService)
	ctrl := controller.NewNewsController(mockService)

	newsID := uuid.New()
	topicID := uuid.New()
	expectedTopic := common.Topic{ID: topicID, Name: "Sample Topic"}
	expectedNews := model.News{
		ID:      newsID,
		Title:   "Some Title",
		Content: "Some Content",
		Status:  "Draft",
		Topics:  []common.Topic{expectedTopic},
	}

	req := httptest.NewRequest("GET", "/v1/api/news/"+newsID.String(), nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/v1/api/news/{id}", ctrl.DetailNews)

	mockService.On("Find", newsID).Return(&expectedNews, nil)
	mockService.On("Preload", &expectedNews).Return(&expectedNews, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.NewsResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewNewsResource(expectedNews), response.Data)

	mockService.AssertExpectations(t)
}

func TestListNews(t *testing.T) {
	mockService := new(mocks.MockNewsService)
	ctrl := controller.NewNewsController(mockService)

	topicID := uuid.New()
	expectedTopic := common.Topic{ID: topicID, Name: "Sample Topic"}
	expectedNews := model.News{
		Title:   "Some Title",
		Content: "Some Content",
		Status:  "Draft",
		Topics:  []common.Topic{expectedTopic},
	}

	newsList := []model.News{expectedNews}

	req := httptest.NewRequest("GET", "/v1/api/news", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/v1/api/news", ctrl.ListNews)

	mockService.On("GetAll").Return(newsList, nil)
	mockService.On("Preload", &expectedNews).Return(&expectedNews, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data []resource.NewsResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response.Data, len(newsList))
	assert.Equal(t, resource.NewNewsResource(expectedNews), response.Data[0])

	mockService.AssertExpectations(t)
}

func TestCreateNews(t *testing.T) {
	mockService := new(mocks.MockNewsService)
	ctrl := controller.NewNewsController(mockService)

	topicID := uuid.New()
	newsRequest := request.CreateNewsRequest{
		Title:    "New Title",
		Content:  "New Content",
		Status:   "draft",
		TopicIDs: []string{topicID.String()},
	}
	jsonData, _ := json.Marshal(newsRequest)

	req := httptest.NewRequest("POST", "/v1/api/news", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post("/v1/api/news", ctrl.CreateNews)

	// Set up the mock expectation before serving the request
	expectedTopic := common.Topic{ID: topicID, Name: "Sample Topic"}
	expectedNews := model.News{
		Title:   "New Title",
		Content: "New Content",
		Status:  "Draft",
		Topics:  []common.Topic{expectedTopic},
	}

	// Use mock.Anything to handle any input for the Create method
	mockService.On("Create", mock.Anything).Return(&expectedNews, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.NewsResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewNewsResource(expectedNews), response.Data)

	mockService.AssertExpectations(t)
}

func TestUpdateNews(t *testing.T) {
	mockService := new(mocks.MockNewsService)
	ctrl := controller.NewNewsController(mockService)

	newsID := uuid.New()
	topicID := uuid.New()
	newsRequest := request.UpdateNewsRequest{
		Title:    "Updated Title",
		Content:  "Updated Content",
		Status:   "published",
		TopicIDs: []string{topicID.String()},
	}
	jsonData, _ := json.Marshal(newsRequest)

	req := httptest.NewRequest("PUT", "/v1/api/news/"+newsID.String(), bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Put("/v1/api/news/{id}", ctrl.UpdateNews)

	expectedTopic := common.Topic{ID: topicID, Name: "Sample Topic"}
	expectedNews := model.News{
		ID:      newsID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Status:  "Published",
		Topics:  []common.Topic{expectedTopic},
	}

	mockService.On("Find", newsID).Return(&expectedNews, nil)
	mockService.On("Update", newsRequest, newsID).Return(&expectedNews, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.NewsResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewNewsResource(expectedNews), response.Data)

	mockService.AssertExpectations(t)
}

func TestDeleteNews(t *testing.T) {
	mockService := new(mocks.MockNewsService)
	ctrl := controller.NewNewsController(mockService)

	newsID := uuid.New()
	topicID := uuid.New()
	expectedTopic := common.Topic{ID: topicID, Name: "Sample Topic"}
	expectedNews := model.News{
		ID:      newsID,
		Title:   "Some Title",
		Content: "Some Content",
		Status:  "Draft",
		Topics:  []common.Topic{expectedTopic},
	}

	req := httptest.NewRequest("DELETE", "/v1/api/news/"+newsID.String(), nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Delete("/v1/api/news/{id}", ctrl.DeleteNews)

	mockService.On("Delete", newsID).Return(&expectedNews, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.NewsResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewNewsResource(expectedNews), response.Data)

	mockService.AssertExpectations(t)
}
