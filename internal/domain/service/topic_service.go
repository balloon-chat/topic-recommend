package service

import (
	"context"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"github.com/balloon-chat/topic-recommend/internal/domain/usecase"
	"github.com/balloon-chat/topic-recommend/internal/infrastructure/firebase"
	"github.com/balloon-chat/topic-recommend/internal/infrastructure/firestore"
)

type TopicServiceInterface interface {
	GetPickupTopics() ([]*model.Topic, error)
	SaveRecommendTopics() (*model.RecommendTopics, error)
}

type TopicService struct {
	topicUseCase *usecase.TopicUseCase
}

func NewTopicService(ctx context.Context) (*TopicService, error) {
	messageRepository, err := firebase.NewFirebaseMessageRepository(ctx)
	if err != nil {
		return nil, err
	}
	recommendTopicRepository, err := firebase.NewFirebaseRecommendTopicRepository(ctx)
	if err != nil {
		return nil, err
	}
	topicRepository, err := firestore.NewFirestoreTopicRepository(ctx)
	if err != nil {
		return nil, err
	}

	u := usecase.NewTopicUsecase(messageRepository, recommendTopicRepository, topicRepository)

	service := &TopicService{topicUseCase: u}
	return service, nil
}

func (service *TopicService) GetPickupTopics() ([]*model.Topic, error) {
	return service.topicUseCase.GetPickupTopics()
}

func (service *TopicService) SaveRecommendTopics() (*model.RecommendTopics, error) {
	return service.topicUseCase.UpdateRecommendTopics()
}
