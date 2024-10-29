package service

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"news-topic-management-service/internal/core/topic/model"
	topicRepo "news-topic-management-service/internal/core/topic/repository"
	"news-topic-management-service/internal/core/topic/request"
	"news-topic-management-service/internal/general/model/common"
)

type TopicService interface {
	GetAll(ctx context.Context) ([]model.Topic, error)
	Create(ctx context.Context, req request.CreateTopicRequest) (*model.Topic, error)
	Update(ctx context.Context, req request.UpdateTopicRequest, topicID uuid.UUID) (*common.Topic, error)
	Find(ctx context.Context, topicID uuid.UUID) (*common.Topic, error)
	Delete(ctx context.Context, topicID uuid.UUID) (*model.Topic, error)
	Preload(ctx context.Context, topic *model.Topic) (*model.Topic, error)
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

var tracer = otel.Tracer("github.com/Salaton/tracing/pkg/infrastructure/database/postgres")

func (n topicService) Create(ctx context.Context, req request.CreateTopicRequest) (*model.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-Create")
	defer span.End()

	topic := &model.Topic{
		Name: req.Name,
	}

	createdTopic, err := n.TopicRepository.Create(ctxT, topic)
	if err != nil {
		return nil, err
	}

	return createdTopic, nil
}

func (n topicService) Update(ctx context.Context, req request.UpdateTopicRequest, topicID uuid.UUID) (*common.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-Update")
	defer span.End()

	existingTopic, err := n.TopicRepository.Find(ctxT, topicID)
	if err != nil {
		return nil, err
	}

	existingTopic.Name = req.Name
	updatedTopic, err := n.TopicRepository.Update(ctxT, topicID, existingTopic)
	if err != nil {
		return nil, err
	}

	return updatedTopic, nil
}

func (n topicService) Find(ctx context.Context, topicID uuid.UUID) (*common.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-Find")
	defer span.End()

	topic, err := n.TopicRepository.Find(ctxT, topicID)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (n topicService) Preload(ctx context.Context, topic *model.Topic) (*model.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-Preload")
	defer span.End()

	preloadedTopic, err := n.TopicRepository.Preload(ctxT, topic)
	if err != nil {
		return nil, err
	}

	return preloadedTopic, nil
}

func (n topicService) Delete(ctx context.Context, topicID uuid.UUID) (*model.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-Delete")
	defer span.End()

	_, err := n.TopicRepository.Find(ctxT, topicID)
	if err != nil {
		return nil, err
	}

	deletedTopic, err := n.TopicRepository.Delete(ctxT, topicID)
	if err != nil {
		return nil, err
	}

	return deletedTopic, nil
}

func (n topicService) GetAll(ctx context.Context) ([]model.Topic, error) {
	ctxT, span := tracer.Start(ctx, "topicService-GetAll")
	defer span.End()

	data, err := n.TopicRepository.GetAll(ctxT)
	if err != nil {
		return nil, err
	}

	return data, nil
}
