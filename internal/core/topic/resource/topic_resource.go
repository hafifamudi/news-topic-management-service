package resource

import (
	"github.com/google/uuid"
	"news-topic-management-service/internal/core/topic/model"
	"time"
)

type NewsResource struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TopicResource struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	News      []NewsResource `json:"news"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func NewTopicResource(topic model.Topic) TopicResource {
	var newsResources []NewsResource
	for _, news := range topic.News {
		newsResources = append(newsResources, NewsResource{
			ID:        news.ID,
			Title:     news.Title,
			Content:   news.Content,
			Status:    news.Status,
			CreatedAt: news.CreatedAt,
			UpdatedAt: news.UpdatedAt,
		})
	}
	return TopicResource{
		ID:        topic.ID,
		Name:      topic.Name,
		News:      newsResources,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}
}
