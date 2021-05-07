package repository

import "github.com/balloon-chat/topic-recommend/internal/domain/model"

type TopicDatabase interface {
	GetTopicsOrderByCreatedAt() ([]*model.Topic, error)
}
