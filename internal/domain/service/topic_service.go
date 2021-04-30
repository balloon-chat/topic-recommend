package service

import (
	"context"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"github.com/balloon-chat/topic-recommend/internal/domain/repository"
	"github.com/balloon-chat/topic-recommend/internal/infrastructure/firebase"
	"github.com/balloon-chat/topic-recommend/internal/infrastructure/firestore"
	"sort"
)

type TopicService struct {
	messageRepository   repository.MessageDatabase
	recommendRepository repository.RecommendTopicDatabase
	topicRepository     repository.TopicDatabase
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

	service := &TopicService{
		messageRepository:   messageRepository,
		recommendRepository: recommendTopicRepository,
		topicRepository:     topicRepository,
	}
	return service, nil
}

func (service *TopicService) GetPickupTopics() ([]*model.Topic, error) {
	topics, err := service.topicRepository.GetTopicsOrderByCreatedAt()
	if err != nil {
		return nil, err
	}

	var data []*model.TopicData
	for _, topic := range topics {
		count, err := service.messageRepository.GetMessageCountOf(topic.Id)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.TopicData{
			Topic:        *topic,
			MessageCount: *count,
		})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].MessageCount > data[j].MessageCount
	})

	var result []*model.Topic
	for _, d := range data {
		result = append(result, &d.Topic)
	}

	return result, nil
}

func (service *TopicService) GetNewestTopics() ([]*model.Topic, error) {
	return service.topicRepository.GetTopicsOrderByCreatedAt()
}

func (service *TopicService) SaveRecommendTopics() (*model.RecommendTopics, error) {
	pickup, err := service.GetPickupTopics()
	if err != nil {
		return nil, err
	}

	newest, err := service.GetNewestTopics()
	if err != nil {
		return nil, err
	}

	var pickupTopicIds []string
	var newestTopicIds []string

	for _, p := range pickup {
		pickupTopicIds = append(pickupTopicIds, p.Id)
	}

	for _, n := range newest {
		newestTopicIds = append(newestTopicIds, n.Id)
	}

	recommend := model.RecommendTopics{
		Pickup: pickupTopicIds,
		Newest: newestTopicIds,
	}

	err = service.recommendRepository.SaveRecommendTopics(recommend)
	if err != nil {
		return nil, err
	}

	return &recommend, nil
}
