package service

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/model"
	topicRepo "github.com/hafifamudi/news-topic-management-service/internal/core/topic/repository"
	"github.com/hafifamudi/news-topic-management-service/internal/core/topic/request"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
)

type TopicService interface {
	GetAll() ([]model.Topic, error)
	Create(req request.CreateTopicRequest) (*model.Topic, error)
	Update(req request.UpdateTopicRequest, topicID uuid.UUID) (*common.Topic, error)
	Find(topicID uuid.UUID) (*common.Topic, error)
	Delete(topicID uuid.UUID) (*model.Topic, error)
	Preload(topic *model.Topic) (*model.Topic, error)
}

type topicService struct {
	TopicRepository topicRepo.TopicRepository
}

func NewTopicService(
	TopicRepository topicRepo.TopicRepository,
) *topicService {
	return &topicService{
		TopicRepository: TopicRepository,
	}
}

func Topic() TopicService {
	return NewTopicService(
		topicRepo.Topic(),
	)
}

func (n topicService) Create(req request.CreateTopicRequest) (*model.Topic, error) {
	topic := &model.Topic{
		Name: req.Name,
	}

	createdTopic, err := n.TopicRepository.Create(topic)
	if err != nil {
		return nil, err
	}

	return createdTopic, nil
}

func (n topicService) Update(req request.UpdateTopicRequest, topicID uuid.UUID) (*common.Topic, error) {
	existingTopic, err := n.TopicRepository.Find(topicID)
	if err != nil {
		return nil, err
	}

	existingTopic.Name = req.Name
	updatedTopic, err := n.TopicRepository.Update(topicID, existingTopic)
	if err != nil {
		return nil, err
	}

	return updatedTopic, nil
}

func (n topicService) Find(topicID uuid.UUID) (*common.Topic, error) {
	topic, err := n.TopicRepository.Find(topicID)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (n topicService) Preload(topic *model.Topic) (*model.Topic, error) {
	preloadedTopic, err := n.TopicRepository.Preload(topic)
	if err != nil {
		return nil, err
	}

	return preloadedTopic, nil
}

func (n topicService) Delete(topicID uuid.UUID) (*model.Topic, error) {
	_, err := n.TopicRepository.Find(topicID)
	if err != nil {
		return nil, err
	}

	deletedTopic, err := n.TopicRepository.Delete(topicID)
	if err != nil {
		return nil, err
	}

	return deletedTopic, nil
}

func (n topicService) GetAll() ([]model.Topic, error) {
	data, err := n.TopicRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
