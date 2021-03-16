package service

import (
	"fmt"
	"github.com/balloon-chat/topic-recommend/src/model"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

type fakeMessageDatabase struct {
	data map[string][]model.Message
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
	data map[string]*model.Topic
}

func (f fakeTopicDatabase) GetTopicsOrderByCreatedAt() ([]*model.Topic, error) {
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

var (
	messageDB = fakeMessageDatabase{
		data: make(map[string][]model.Message),
	}
	topicDB = fakeTopicDatabase{
		data: make(map[string]*model.Topic),
	}
	recommendDB = fakeRecommendTopicDatabase{
		data: nil,
	}
	service = TopicService{
		messageDatabase:   &messageDB,
		topicDatabase:     &topicDB,
		recommendDatabase: &recommendDB,
	}
)

func TestMain(m *testing.M) {
	messageDB.data = make(map[string][]model.Message)
	topicDB.data = make(map[string]*model.Topic)
	recommendDB.data = nil
	m.Run()
}

func TestTopicService_GetNewestTopics(t *testing.T) {
	for i := 0; i < 10; i++ {
		topicDB.data[strconv.Itoa(i)] = &model.Topic{
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
		topic := &model.Topic{
			Id:        strconv.Itoa(i),
			CreatedAt: rand.Int(),
		}
		topicDB.data[strconv.Itoa(i)] = topic
		messageDB.data[topic.Id] = make([]model.Message, i)
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
