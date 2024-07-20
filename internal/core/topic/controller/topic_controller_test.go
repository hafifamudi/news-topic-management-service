package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/request"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/resource"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/controller"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/model"
	"github.com/hafifamudi/news-topic-management-service/internal/general/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListTopic(t *testing.T) {
	mockService := new(mocks.MockTopicService)
	ctrl := controller.NewTopicController(mockService)

	topicID := uuid.New()
	expectedTopic := model.Topic{
		ID:   topicID,
		Name: "Sample Topic",
	}

	topicList := []model.Topic{expectedTopic}

	req := httptest.NewRequest("GET", "/v1/api/topics", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/v1/api/topics", ctrl.ListTopic)

	mockService.On("GetAll").Return(topicList, nil)
	mockService.On("Preload", &expectedTopic).Return(&expectedTopic, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data []resource.TopicResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response.Data, len(topicList))
	assert.Equal(t, resource.NewTopicResource(expectedTopic), response.Data[0])

	mockService.AssertExpectations(t)
}

func TestCreateTopic(t *testing.T) {
	mockService := new(mocks.MockTopicService)
	ctrl := controller.NewTopicController(mockService)

	topicID := uuid.New()
	topicRequest := request.CreateTopicRequest{
		Name: "New Topic",
	}
	jsonData, _ := json.Marshal(topicRequest)

	req := httptest.NewRequest("POST", "/v1/api/topics", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post("/v1/api/topics", ctrl.CreateTopic)

	expectedTopic := model.Topic{
		ID:   topicID,
		Name: "New Topic",
	}

	mockService.On("Create", topicRequest).Return(&expectedTopic, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.TopicResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewTopicResource(expectedTopic), response.Data)

	mockService.AssertExpectations(t)
}

func TestUpdateTopic(t *testing.T) {
	mockService := new(mocks.MockTopicService)
	ctrl := controller.NewTopicController(mockService)

	topicID := uuid.New()
	topicRequest := request.UpdateTopicRequest{
		Name: "Updated Topic",
	}
	jsonData, _ := json.Marshal(topicRequest)

	req := httptest.NewRequest("PUT", "/v1/api/topics/"+topicID.String(), bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Put("/v1/api/topics/{id}", ctrl.UpdateTopic)

	expectedTopic := common.Topic{
		ID:   topicID,
		Name: "Updated Topic",
	}

	mockService.On("Find", topicID).Return(&expectedTopic, nil)
	mockService.On("Update", topicRequest, topicID).Return(&expectedTopic, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.TopicResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewTopicResource((model.Topic)(expectedTopic)), response.Data)

	mockService.AssertExpectations(t)
}

func TestDeleteTopic(t *testing.T) {
	mockService := new(mocks.MockTopicService)
	ctrl := controller.NewTopicController(mockService)

	topicID := uuid.New()
	expectedTopic := model.Topic{
		ID:   topicID,
		Name: "Sample Topic",
	}

	req := httptest.NewRequest("DELETE", "/v1/api/topics/"+topicID.String(), nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Delete("/v1/api/topics/{id}", ctrl.DeleteTopic)

	mockService.On("Delete", topicID).Return(&expectedTopic, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Data resource.TopicResource `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, resource.NewTopicResource(expectedTopic), response.Data)

	mockService.AssertExpectations(t)
}
