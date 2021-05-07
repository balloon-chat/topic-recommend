package usecase

import (
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"sort"
)

func (u *TopicUseCase) GetPickupTopics() ([]*model.Topic, error) {
	topics, err := u.topicRepository.GetTopicsOrderByCreatedAt()
	if err != nil {
		return nil, err
	}

	var data []*model.TopicData
	for _, topic := range topics {
		count, err := u.messageRepository.GetMessageCountOf(topic.Id)
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
