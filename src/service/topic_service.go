package service

import (
	"context"
	"github.com/balloon-chat/topic-recommend/src/database"
	"github.com/balloon-chat/topic-recommend/src/model"
	"sort"
)

type TopicService struct {
	messageDatabase   database.MessageDatabase
	recommendDatabase database.RecommendTopicDatabase
	topicDatabase     database.TopicDatabase
}

func NewTopicService(ctx context.Context) (*TopicService, error) {
	messageDB, err := database.NewFirebaseMessageDatabase(ctx)
	if err != nil {
		return nil, err
	}
	recommendDB, err := database.NewFirebaseRecommendTopicDatabase(ctx)
	if err != nil {
		return nil, err
	}
	topicDB, err := database.NewFirebaseTopicDatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &TopicService{
		messageDatabase:   messageDB,
		recommendDatabase: recommendDB,
		topicDatabase:     topicDB,
	}, nil
}

func (service *TopicService) GetPickupTopics() ([]*model.Topic, error) {
	topics, err := service.topicDatabase.GetTopicsOrderByCreatedAt()
	if err != nil {
		return nil, err
	}

	var data []*model.TopicData
	for _, topic := range topics {
		count, err := service.messageDatabase.GetMessageCountOf(topic.Id)
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
	return service.topicDatabase.GetTopicsOrderByCreatedAt()
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

	err = service.recommendDatabase.SaveRecommendTopics(recommend)
	if err != nil {
		return nil, err
	}

	return &recommend, nil
}
