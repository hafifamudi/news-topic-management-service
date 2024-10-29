package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"news-topic-management-service/internal/core/topic/model"
	"news-topic-management-service/internal/general/model/common"
)

type TopicRepository interface {
	GetAll(ctx context.Context) ([]model.Topic, error)
	Create(ctx context.Context, topic *model.Topic) (*model.Topic, error)
	Update(ctx context.Context, id uuid.UUID, topic *common.Topic) (*common.Topic, error)
	Find(ctx context.Context, id uuid.UUID) (*common.Topic, error)
	Delete(ctx context.Context, topicID uuid.UUID) (*model.Topic, error)
	Preload(ctx context.Context, topic *model.Topic) (*model.Topic, error)
}

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{
		db: db,
	}
}

func Topic() TopicRepository {
	return NewTopicRepository(db.DB)
}

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

func (r *topicRepository) GetAll(ctx context.Context) ([]model.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-ListTopic")
	defer span.End()

	var topics []model.Topic
	result := r.db.Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *topicRepository) Create(ctx context.Context, topic *model.Topic) (*model.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-Create")
	defer span.End()

	if topic.ID == uuid.Nil {
		topic.ID = uuid.New()
	}

	result := r.db.Create(&topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}

func (r *topicRepository) Find(ctx context.Context, id uuid.UUID) (*common.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-Find")
	defer span.End()

	var topic common.Topic
	result := r.db.First(&topic, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &topic, nil
}

func (r *topicRepository) Update(ctx context.Context, id uuid.UUID, topic *common.Topic) (*common.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-Update")
	defer span.End()

	result := r.db.Model(&topic).Where("id = ?", id).Updates(topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}

func (r *topicRepository) Delete(ctx context.Context, topicID uuid.UUID) (*model.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-Delete")
	defer span.End()

	var topic model.Topic
	result := r.db.First(&topic, "id = ?", topicID)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Delete(&topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return &topic, nil
}

func (r *topicRepository) Preload(ctx context.Context, topic *model.Topic) (*model.Topic, error) {
	_, span := tracer.Start(ctx, "topicRepository-Preload")
	defer span.End()

	result := r.db.Preload("News").Find(&topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}
