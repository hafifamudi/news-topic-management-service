package controller

import (
	"context"
	"encoding/json"
	errorRequest "github.com/hafifamudi/news-topic-management-service/pkg/utils/errors"
	errorResponse "github.com/hafifamudi/news-topic-management-service/pkg/utils/validations"
	"go.opentelemetry.io/otel"
	"io"
	"net/http"
	"news-topic-management-service/internal/core/topic/model"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/pkg/utils/response"
	"news-topic-management-service/internal/core/topic/request"
	"news-topic-management-service/internal/core/topic/resource"
	"news-topic-management-service/internal/core/topic/service"
)

type TopicController interface {
	ListTopic(w http.ResponseWriter, r *http.Request)
	DetailTopic(w http.ResponseWriter, r *http.Request)
	CreateTopic(w http.ResponseWriter, r *http.Request)
	UpdateTopic(w http.ResponseWriter, r *http.Request)
	DeleteTopic(w http.ResponseWriter, r *http.Request)
}

type topicController struct {
	service service.TopicService
}

func NewTopicController(service service.TopicService) TopicController {
	return &topicController{
		service: service,
	}
}

func Topic() TopicController {
	return NewTopicController(service.Topic())
}

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

// ListTopic @Summary List all topics
// @Description List all topics with related topic
// @Tags Topics
// @Accept  json
// @Produce  json
// @Success 200 {object} common.SuccessWithMessageResponse{data=[]common.TopicResource}
// @Router /topics [get]
func (c *topicController) ListTopic(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "topicController-ListTopic")
	defer span.End()

	topicList, err := c.service.GetAll(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error fetching topics")
		return
	}

	var topicResources []resource.TopicResource
	for _, topic := range topicList {
		preloadedTopic, err := c.service.Preload(ctx, &topic)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Error fetching related topic")
			return
		}
		topicResources = append(topicResources, resource.NewTopicResource(*preloadedTopic))
	}

	response.SuccessWithMessage(w, "List of topics retrieved successfully", topicResources)
}

// DetailTopic @Summary Detail data of a Topic
// @Description Detail Topic with the provided information
// @Tags Topics
// @Produce json
// @Param id path string true "Topic ID" Format(uuid)
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.TopicResource}
// @Router /topics/{id} [get]
func (c *topicController) DetailTopic(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "topicController-DetailTopic")
	defer span.End()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid News ID")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	topic, err := c.service.Find(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if topic.ID == uuid.Nil {
		response.Error(w, http.StatusNotFound, "News not found")
		return
	}

	var NewsResources []resource.TopicResource
	preloadedNews, err := c.service.Preload(ctx, (*model.Topic)(topic))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	NewsResources = append(NewsResources, resource.NewTopicResource(*preloadedNews))

	response.SuccessWithMessage(w, "News updated", resource.NewTopicResource(model.Topic(*topic)))
}

// CreateTopic @Summary Create a new Topic
// @Description Create a new Topic with the provided information
// @Tags Topics
// @Accept json
// @Produce json
// @Param topic body request.CreateTopicRequest true "Create Topic"
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.TopicResource}
// @Router /topics [post]
func (c *topicController) CreateTopic(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "topicController-CreateTopic")
	defer span.End()

	var topicRequest request.CreateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&topicRequest); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := errorRequest.ValidateStruct(topicRequest); err != nil {
		errorResponse.HandleHttpRequestValidationError(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	topic, err := c.service.Create(ctx, topicRequest)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(w, "Topic created", resource.NewTopicResource(*topic))
}

// UpdateTopic @Summary Update an existing Topic
// @Description Update an existing Topic with the provided information
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID" Format(uuid)
// @Param topic body request.UpdateTopicRequest true "Update Topic"
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.TopicResource}
// @Router /topics/{id} [put]
func (c *topicController) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "topicController-UpdateTopic")
	defer span.End()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Topic ID")
		return
	}

	var topicRequest request.UpdateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&topicRequest); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := errorRequest.ValidateStruct(topicRequest); err != nil {
		errorResponse.HandleHttpRequestValidationError(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	topic, err := c.service.Find(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error fetching Topic")
		return
	}

	if topic.ID == uuid.Nil {
		response.Error(w, http.StatusNotFound, "Topic not found")
		return
	}

	topic, err = c.service.Update(ctx, topicRequest, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(w, "Topic updated", resource.NewTopicResource((model.Topic)(*topic)))
}

// DeleteTopic @Summary Delete a Topic
// @Description Delete a Topic with the provided Topic ID
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID" Format(uuid)
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.TopicResource}
// @Router /topics/{id} [delete]
func (c *topicController) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "topicController-DeleteTopic")
	defer span.End()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Topic ID")
		return
	}

	deletedTopic, err := c.service.Delete(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error deleting Topic")
		return
	}

	response.SuccessWithMessage(w, "Topic deleted", resource.NewTopicResource(*deletedTopic))
}
