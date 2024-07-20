package service

import (
	"github.com/google/uuid"
	"news-topic-management-service/internal/core/news/model"
	newsRepo "news-topic-management-service/internal/core/news/repository"
	"news-topic-management-service/internal/core/news/request"
	topicRepo "news-topic-management-service/internal/core/topic/repository"
	"news-topic-management-service/internal/general/model/common"
)

type NewsService interface {
	GetAll() ([]model.News, error)
	Create(req request.CreateNewsRequest) (*model.News, error)
	Update(req request.UpdateNewsRequest, newsID uuid.UUID) (*model.News, error)
	Find(newsID uuid.UUID) (*model.News, error)
	Delete(newsID uuid.UUID) (*model.News, error)
	Preload(news *model.News) (*model.News, error)
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

func (n newsService) Create(req request.CreateNewsRequest) (*model.News, error) {
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

		topic, err := n.TopicRepository.Find(parsedTopicID)
		if err != nil {
			return nil, err
		}

		news.Topics = append(news.Topics, *topic)
	}

	createdNews, err := n.NewsRepository.Create(news)
	if err != nil {
		return nil, err
	}

	return createdNews, nil
}

func (n newsService) Update(req request.UpdateNewsRequest, newsID uuid.UUID) (*model.News, error) {
	existingNews, err := n.NewsRepository.Find(newsID)
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

		topic, err := n.TopicRepository.Find(parsedTopicID)
		if err != nil {
			return nil, err
		}

		existingNews.Topics = append(existingNews.Topics, *topic)
	}

	updatedNews, err := n.NewsRepository.Update(newsID, existingNews)
	if err != nil {
		return nil, err
	}

	return updatedNews, nil
}

func (n newsService) Find(newsID uuid.UUID) (*model.News, error) {
	news, err := n.NewsRepository.Find(newsID)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n newsService) Preload(news *model.News) (*model.News, error) {
	preloadedNews, err := n.NewsRepository.Preload(news)
	if err != nil {
		return nil, err
	}

	return preloadedNews, nil
}

func (n newsService) Delete(newsID uuid.UUID) (*model.News, error) {
	// Find the news item to delete
	_, err := n.NewsRepository.Find(newsID)
	if err != nil {
		return nil, err
	}

	// Delete the news item from the database
	deletedNews, err := n.NewsRepository.Delete(newsID)
	if err != nil {
		return nil, err
	}

	return deletedNews, nil
}

func (n newsService) GetAll() ([]model.News, error) {
	data, err := n.NewsRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
