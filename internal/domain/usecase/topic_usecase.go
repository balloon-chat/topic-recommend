package usecase

import "github.com/balloon-chat/topic-recommend/internal/domain/repository"

type TopicUseCase struct {
	messageRepository   repository.MessageDatabase
	recommendRepository repository.RecommendTopicDatabase
	topicRepository     repository.TopicDatabase
}

func NewTopicUsecase(m repository.MessageDatabase, r repository.RecommendTopicDatabase, t repository.TopicDatabase) *TopicUseCase {
	return &TopicUseCase{
		messageRepository:   m,
		recommendRepository: r,
		topicRepository:     t,
	}
}
