package service

import (
	"fmt"
	model2 "github.com/balloon-chat/topic-recommend/internal/domain/model"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

type fakeMessageDatabase struct {
	data map[string][]model2.Message
}

func (f fakeMessageDatabase) GetMessageCountOf(topicId string) (*int, error) {
	m, ok := f.data[topicId]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	c := len(m)
	return &c, nil
}

type fakeTopicDatabase struct {
	data map[string]*model2.Topic
}

func (f fakeTopicDatabase) GetTopicsOrderByCreatedAt() ([]*model2.Topic, error) {
	var result []*model2.Topic

	for _, topic := range f.data {
		result = append(result, topic)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})

	return result, nil
}

type fakeRecommendTopicDatabase struct {
	data *model2.RecommendTopics
}

func (f fakeRecommendTopicDatabase) SaveRecommendTopics(recommend model2.RecommendTopics) error {
	f.data = &recommend
	return nil
}

var (
	messageDB = fakeMessageDatabase{
		data: make(map[string][]model2.Message),
	}
	topicDB = fakeTopicDatabase{
		data: make(map[string]*model2.Topic),
	}
	recommendDB = fakeRecommendTopicDatabase{
		data: nil,
	}
	service = TopicService{
		messageRepository:   &messageDB,
		topicRepository:     &topicDB,
		recommendRepository: &recommendDB,
	}
)

func TestMain(m *testing.M) {
	messageDB.data = make(map[string][]model2.Message)
	topicDB.data = make(map[string]*model2.Topic)
	recommendDB.data = nil
	m.Run()
}

func TestTopicService_GetNewestTopics(t *testing.T) {
	for i := 0; i < 10; i++ {
		topicDB.data[strconv.Itoa(i)] = &model2.Topic{
			Id:        strconv.Itoa(i),
			CreatedAt: rand.Int(),
		}
	}
	topics, _ := service.GetNewestTopics()
	for i := 1; i < len(topics); i++ {
		if topics[i-1].CreatedAt < topics[i].CreatedAt {
			t.Fatal("topics must be sorted by createdAt")
		}
	}
}

func TestTopicService_GetPickupTopics(t *testing.T) {
	for i := 0; i < 10; i++ {
		topic := &model2.Topic{
			Id:        strconv.Itoa(i),
			CreatedAt: rand.Int(),
		}
		topicDB.data[strconv.Itoa(i)] = topic
		messageDB.data[topic.Id] = make([]model2.Message, i)
	}

	topics, _ := service.GetPickupTopics()
	for i := 1; i < len(topics); i++ {
		current := messageDB.data[topics[i].Id]
		prev := messageDB.data[topics[i-1].Id]
		if len(prev) < len(current) {
			t.Fatal("topics must be sorted by message count")
		}
	}
}
