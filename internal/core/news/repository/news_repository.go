package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"news-topic-management-service/internal/core/news/model"
)

type NewsRepository interface {
	GetAll(ctx context.Context, status *string, topicID *uuid.UUID) ([]model.News, error)
	Create(ctx context.Context, news *model.News) (*model.News, error)
	Update(ctx context.Context, id uuid.UUID, news *model.News) (*model.News, error)
	Find(ctx context.Context, id uuid.UUID) (*model.News, error)
	Delete(ctx context.Context, newsID uuid.UUID) (*model.News, error)
	Preload(ctx context.Context, news *model.News) (*model.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

func News() NewsRepository {
	return NewNewsRepository(db.DB)
}

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

func (r *newsRepository) GetAll(ctx context.Context, status *string, topicID *uuid.UUID) ([]model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	var news []model.News
	query := r.db.Preload("Topics")

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if topicID != nil {
		query = query.Joins("JOIN news_topics ON news.id = news_topics.news_id").Where("news_topics.topic_id = ?", *topicID)
	}

	result := query.Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

func (r *newsRepository) Create(ctx context.Context, news *model.News) (*model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	if news.ID == uuid.Nil {
		news.ID = uuid.New()
	}

	result := r.db.Create(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

func (r *newsRepository) Find(ctx context.Context, id uuid.UUID) (*model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	var news model.News
	result := r.db.First(&news, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &news, nil
}

func (r *newsRepository) Update(ctx context.Context, id uuid.UUID, news *model.News) (*model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	result := r.db.Model(&model.News{}).Where("id = ?", id).Updates(news)
	if result.Error != nil {
		return nil, result.Error
	}

	err := r.db.Model(news).Association("Topics").Replace(news.Topics)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *newsRepository) Delete(ctx context.Context, newsID uuid.UUID) (*model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	var news model.News
	result := r.db.Where("id = ?", newsID).First(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Delete(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return &news, nil
}

func (r *newsRepository) Preload(ctx context.Context, news *model.News) (*model.News, error) {
	_, span := tracer.Start(ctx, "newsRepository-ListTopic")
	defer span.End()

	if err := r.db.Preload("Topics").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
