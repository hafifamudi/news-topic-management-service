package service

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"news-topic-management-service/internal/core/news/model"
	newsRepo "news-topic-management-service/internal/core/news/repository"
	"news-topic-management-service/internal/core/news/request"
	topicRepo "news-topic-management-service/internal/core/topic/repository"
	"news-topic-management-service/internal/general/model/common"
)

type NewsService interface {
	GetAll(ctx context.Context, status *string, topicID *uuid.UUID) ([]model.News, error)
	Create(ctx context.Context, req request.CreateNewsRequest) (*model.News, error)
	Update(ctx context.Context, req request.UpdateNewsRequest, newsID uuid.UUID) (*model.News, error)
	Find(ctx context.Context, newsID uuid.UUID) (*model.News, error)
	Delete(ctx context.Context, newsID uuid.UUID) (*model.News, error)
	Preload(ctx context.Context, news *model.News) (*model.News, error)
}

type newsService struct {
	NewsRepository  newsRepo.NewsRepository
	TopicRepository topicRepo.TopicRepository
}

func NewNewsService(
	NewsRepository newsRepo.NewsRepository,
	TopicRepository topicRepo.TopicRepository,
) *newsService {
	return &newsService{
		NewsRepository:  NewsRepository,
		TopicRepository: TopicRepository,
	}
}

func News() NewsService {
	return NewNewsService(
		newsRepo.News(),
		topicRepo.Topic(),
	)
}

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

func (n newsService) Create(ctx context.Context, req request.CreateNewsRequest) (*model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	news := &model.News{
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
		Topics:  nil,
	}

	for _, topicID := range req.TopicIDs {
		parsedTopicID, err := uuid.Parse(topicID)
		if err != nil {
			return nil, err
		}

		topic, err := n.TopicRepository.Find(ctxT, parsedTopicID)
		if err != nil {
			return nil, err
		}

		news.Topics = append(news.Topics, *topic)
	}

	createdNews, err := n.NewsRepository.Create(ctxT, news)
	if err != nil {
		return nil, err
	}

	return createdNews, nil
}

func (n newsService) Update(ctx context.Context, req request.UpdateNewsRequest, newsID uuid.UUID) (*model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	existingNews, err := n.NewsRepository.Find(ctxT, newsID)
	if err != nil {
		return nil, err
	}

	existingNews.Title = req.Title
	existingNews.Content = req.Content
	existingNews.Status = req.Status

	existingNews.Topics = []common.Topic{}

	for _, topicID := range req.TopicIDs {
		parsedTopicID, err := uuid.Parse(topicID)
		if err != nil {
			return nil, err
		}

		topic, err := n.TopicRepository.Find(ctxT, parsedTopicID)
		if err != nil {
			return nil, err
		}

		existingNews.Topics = append(existingNews.Topics, *topic)
	}

	updatedNews, err := n.NewsRepository.Update(ctxT, newsID, existingNews)
	if err != nil {
		return nil, err
	}

	return updatedNews, nil
}

func (n newsService) Find(ctx context.Context, newsID uuid.UUID) (*model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	news, err := n.NewsRepository.Find(ctxT, newsID)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n newsService) Preload(ctx context.Context, news *model.News) (*model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	preloadedNews, err := n.NewsRepository.Preload(ctxT, news)
	if err != nil {
		return nil, err
	}

	return preloadedNews, nil
}

func (n newsService) Delete(ctx context.Context, newsID uuid.UUID) (*model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	_, err := n.NewsRepository.Find(ctxT, newsID)
	if err != nil {
		return nil, err
	}

	deletedNews, err := n.NewsRepository.Delete(ctxT, newsID)
	if err != nil {
		return nil, err
	}

	return deletedNews, nil
}

func (n newsService) GetAll(ctx context.Context, status *string, topicID *uuid.UUID) ([]model.News, error) {
	ctxT, span := tracer.Start(ctx, "newsService-ListTopic")
	defer span.End()

	data, err := n.NewsRepository.GetAll(ctxT, status, topicID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
