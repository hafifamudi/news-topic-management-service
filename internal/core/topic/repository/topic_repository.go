package repository

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/model"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"gorm.io/gorm"
)

type TopicRepository interface {
	GetAll() ([]model.Topic, error)
	Create(topic *model.Topic) (*model.Topic, error)
	Update(id uuid.UUID, topic *common.Topic) (*common.Topic, error)
	Find(id uuid.UUID) (*common.Topic, error)
	Delete(topicID uuid.UUID) (*model.Topic, error)
	Preload(topic *model.Topic) (*model.Topic, error)
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

func (r *topicRepository) GetAll() ([]model.Topic, error) {
	var topics []model.Topic
	result := r.db.Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *topicRepository) Create(topic *model.Topic) (*model.Topic, error) {
	if topic.ID == uuid.Nil {
		topic.ID = uuid.New()
	}

	result := r.db.Create(&topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}

func (r *topicRepository) Find(id uuid.UUID) (*common.Topic, error) {
	var topic common.Topic
	result := r.db.First(&topic, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &topic, nil
}

func (r *topicRepository) Update(id uuid.UUID, topic *common.Topic) (*common.Topic, error) {
	result := r.db.Model(&topic).Where("id = ?", id).Updates(topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}

func (r *topicRepository) Delete(topicID uuid.UUID) (*model.Topic, error) {
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

func (r *topicRepository) Preload(topic *model.Topic) (*model.Topic, error) {
	result := r.db.Preload("News").Find(&topic)
	if result.Error != nil {
		return nil, result.Error
	}

	return topic, nil
}
