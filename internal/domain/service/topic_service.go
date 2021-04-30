package service

import (
	"context"
	model2 "github.com/balloon-chat/topic-recommend/internal/domain/model"
	"github.com/balloon-chat/topic-recommend/internal/infrastructure/firebase"
	"sort"
)

type TopicService struct {
	messageDatabase   firebase.MessageDatabase
	recommendDatabase firebase.RecommendTopicDatabase
	topicDatabase     firebase.TopicDatabase
}

func NewTopicService(ctx context.Context) (*TopicService, error) {
	messageDB, err := firebase.NewFirebaseMessageDatabase(ctx)
	if err != nil {
		return nil, err
	}
	recommendDB, err := firebase.NewFirebaseRecommendTopicDatabase(ctx)
	if err != nil {
		return nil, err
	}
	topicDB, err := firebase.NewFirebaseTopicDatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &TopicService{
		messageDatabase:   messageDB,
		recommendDatabase: recommendDB,
		topicDatabase:     topicDB,
	}, nil
}

func (service *TopicService) GetPickupTopics() ([]*model2.Topic, error) {
	topics, err := service.topicDatabase.GetTopicsOrderByCreatedAt()
	if err != nil {
		return nil, err
	}

	var data []*model2.TopicData
	for _, topic := range topics {
		count, err := service.messageDatabase.GetMessageCountOf(topic.Id)
		if err != nil {
			return nil, err
		}
		data = append(data, &model2.TopicData{
			Topic:        *topic,
			MessageCount: *count,
		})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].MessageCount > data[j].MessageCount
	})

	var result []*model2.Topic
	for _, d := range data {
		result = append(result, &d.Topic)
	}

	return result, nil
}

func (service *TopicService) GetNewestTopics() ([]*model2.Topic, error) {
	return service.topicDatabase.GetTopicsOrderByCreatedAt()
}

func (service *TopicService) SaveRecommendTopics() (*model2.RecommendTopics, error) {
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

	recommend := model2.RecommendTopics{
		Pickup: pickupTopicIds,
		Newest: newestTopicIds,
	}

	err = service.recommendDatabase.SaveRecommendTopics(recommend)
	if err != nil {
		return nil, err
	}

	return &recommend, nil
}
