package repository

import "github.com/balloon-chat/topic-recommend/internal/domain/model"

type RecommendTopicDatabase interface {
	SaveRecommendTopics(recommend model.RecommendTopics) error
}
