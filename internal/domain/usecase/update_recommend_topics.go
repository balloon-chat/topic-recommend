package usecase

import (
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
)

func (u *TopicUseCase) UpdateRecommendTopics() (*model.RecommendTopics, error) {
	pickup, err := u.GetPickupTopics()
	if err != nil {
		return nil, err
	}

	var pickupTopicIds []string

	for _, p := range pickup {
		pickupTopicIds = append(pickupTopicIds, p.Id)
	}

	recommend := model.RecommendTopics{
		Pickup: pickupTopicIds,
	}

	err = u.recommendRepository.SaveRecommendTopics(recommend)
	if err != nil {
		return nil, err
	}

	return &recommend, nil
}
