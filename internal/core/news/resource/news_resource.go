package resource

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/model"
	"time"
)

type TopicResource struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewsResource struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Status    string          `json:"status"`
	Topics    []TopicResource `json:"topics"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func NewNewsResource(news model.News) NewsResource {
	var topicResources []TopicResource
	for _, topics := range news.Topics {
		topicResources = append(topicResources, TopicResource{
			ID:        topics.ID,
			Name:      topics.Name,
			CreatedAt: topics.CreatedAt,
			UpdatedAt: topics.UpdatedAt,
		})
	}
	return NewsResource{
		ID:        news.ID,
		Title:     news.Title,
		Content:   news.Content,
		Status:    news.Status,
		Topics:    topicResources,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}
}
