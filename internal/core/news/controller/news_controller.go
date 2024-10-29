package controller

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	errorRequest "github.com/hafifamudi/news-topic-management-service/pkg/utils/errors"
	"github.com/hafifamudi/news-topic-management-service/pkg/utils/response"
	errorResponse "github.com/hafifamudi/news-topic-management-service/pkg/utils/validations"
	"go.opentelemetry.io/otel"
	"io"
	"net/http"
	"news-topic-management-service/internal/core/news/request"
	"news-topic-management-service/internal/core/news/resource"
	"news-topic-management-service/internal/core/news/service"
)

// NewsController Swagger annotations
// @title News API
// @version 1.0
// @description APIs for managing News
// @host localhost:8080
// @BasePath /api
// @schemes http
// @produce json
// @router /api/swagger.json [get]
type NewsController interface {
	ListNews(w http.ResponseWriter, r *http.Request)
	DetailNews(w http.ResponseWriter, r *http.Request)
	CreateNews(w http.ResponseWriter, r *http.Request)
	UpdateNews(w http.ResponseWriter, r *http.Request)
	DeleteNews(w http.ResponseWriter, r *http.Request)
}

type newsController struct {
	service service.NewsService
}

func NewNewsController(service service.NewsService) NewsController {
	return &newsController{
		service: service,
	}
}

func News() NewsController {
	return NewNewsController(service.News())
}

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

// ListNews @Summary Retrieve all News
// @Description Retrieve all News items with optional filtering by status or topic
// @Tags News
// @Accept json
// @Produce json
// @Param status query string false "Filter by status"
// @Param topicID query string false "Filter by Topic ID"
// @Success 200 {object} common.SuccessWithMessageResponse{data=[]common.NewsResource}
// @Router /news [get]
func (c *newsController) ListNews(w http.ResponseWriter, r *http.Request) {
	ctxT, span := tracer.Start(context.Background(), "newsController-ListNews")
	defer span.End()

	var status *string
	if s := r.URL.Query().Get("status"); s != "" {
		status = &s
	}

	var topicID *uuid.UUID
	if tid := r.URL.Query().Get("topicID"); tid != "" {
		id, err := uuid.Parse(tid)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Topic ID")
			return
		}
		topicID = &id
	}

	newsList, err := c.service.GetAll(ctxT, status, topicID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error fetching News")
		return
	}

	var newsResources []resource.NewsResource
	for _, topic := range newsList {
		preloadedNews, err := c.service.Preload(ctxT, &topic)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Error fetching related news")
			return
		}
		newsResources = append(newsResources, resource.NewNewsResource(*preloadedNews))
	}

	response.SuccessWithMessage(w, "List of News retrieved successfully", newsResources)
}

// DetailNews @Summary Detail data of a News
// @Description Detail News with the provided information
// @Tags News
// @Produce json
// @Param id path string true "News ID" Format(uuid)
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.NewsResource}
// @Router /news/{id} [get]
func (c *newsController) DetailNews(w http.ResponseWriter, r *http.Request) {
	ctxT, span := tracer.Start(context.Background(), "newsController-DetailNews")
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

	news, err := c.service.Find(ctxT, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if news.ID == uuid.Nil {
		response.Error(w, http.StatusNotFound, "News not found")
		return
	}

	var newsResources []resource.NewsResource
	preloadedNews, err := c.service.Preload(ctxT, news)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	newsResources = append(newsResources, resource.NewNewsResource(*preloadedNews))

	response.SuccessWithMessage(w, "News updated", resource.NewNewsResource(*news))
}

// CreateNews @Summary Create a new News
// @Description Create News with the provided information
// @Tags News
// @Accept json
// @Produce json
// @Param CreateNewsRequest body request.CreateNewsRequest true "Create News Request"
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.NewsResource}
// @Router /news [post]
func (c *newsController) CreateNews(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "newsController-CreateNews")
	defer span.End()

	var newsRequest request.CreateNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&newsRequest); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := errorRequest.ValidateStruct(newsRequest); err != nil {
		errorResponse.HandleHttpRequestValidationError(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	news, err := c.service.Create(ctx, newsRequest)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(w, "News created", resource.NewNewsResource(*news))
}

// UpdateNews @Summary Update an existing News
// @Description Update existing News with the provided information
// @Tags News
// @Accept json
// @Produce json
// @Param id path string true "News ID" Format(uuid)
// @Param product body request.UpdateNewsRequest true "Update News"
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.NewsResource}
// @Router /news/{id} [put]
func (c *newsController) UpdateNews(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "newsController-UpdateNews")
	defer span.End()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid News ID")
		return
	}

	var newsRequest request.UpdateNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&newsRequest); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := errorRequest.ValidateStruct(newsRequest); err != nil {
		errorResponse.HandleHttpRequestValidationError(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	news, err := c.service.Find(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if news.ID == uuid.Nil {
		response.Error(w, http.StatusNotFound, "News not found")
		return
	}

	news, err = c.service.Update(ctx, newsRequest, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(w, "News updated", resource.NewNewsResource(*news))
}

// DeleteNews @Summary Delete a News
// @Description Delete News with the provided News ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path string true "News ID" Format(uuid)
// @Success 200 {object} common.SuccessWithMessageResponse{data=common.NewsResource}
// @Router /news/{id} [delete]
func (c *newsController) DeleteNews(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "newsController-DeleteNews")
	defer span.End()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid News ID")
		return
	}

	deletedNews, err := c.service.Delete(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(w, "News deleted", resource.NewNewsResource(*deletedNews))
}
