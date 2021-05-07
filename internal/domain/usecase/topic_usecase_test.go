package usecase

import (
	"fmt"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"sort"
	"testing"
)

var (
	messageDB = fakeMessageRepository{
		data: make(map[string][]model.Message),
	}
	topicDB = fakeTopicRepository{
		data: make(map[string]*model.Topic),
	}
	recommendDB = fakeRecommendTopicDatabase{
		data: nil,
	}
)

func TestMain(m *testing.M) {
	messageDB.data = make(map[string][]model.Message)
	topicDB.data = make(map[string]*model.Topic)
	recommendDB.data = nil
	m.Run()
}

type fakeMessageRepository struct {
	data map[string][]model.Message
}

func (f fakeMessageRepository) GetMessageCountOf(topicId string) (*int, error) {
	m, ok := f.data[topicId]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	c := len(m)
	return &c, nil
}

type fakeTopicRepository struct {
	data map[string]*model.Topic
}

func (f fakeTopicRepository) GetTopicsOrderByCreatedAt() ([]*model.Topic, error) {
	var result []*model.Topic

	for _, topic := range f.data {
		result = append(result, topic)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})

	return result, nil
}

type fakeRecommendTopicDatabase struct {
	data *model.RecommendTopics
}

func (f fakeRecommendTopicDatabase) SaveRecommendTopics(recommend model.RecommendTopics) error {
	f.data = &recommend
	return nil
}
