package usecase

import (
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"math/rand"
	"strconv"
	"testing"
)

func TestTopicUseCase_GetPickupTopics(t *testing.T) {
	u := NewTopicUsecase(messageDB, recommendDB, topicDB)
	for i := 0; i < 10; i++ {
		topic := &model.Topic{
			Id:        strconv.Itoa(i),
			CreatedAt: rand.Int(),
		}
		topicDB.data[strconv.Itoa(i)] = topic
		messageDB.data[topic.Id] = make([]model.Message, i)
	}

	topics, _ := u.GetPickupTopics()
	for i := 1; i < len(topics); i++ {
		current := messageDB.data[topics[i].Id]
		prev := messageDB.data[topics[i-1].Id]
		if len(prev) < len(current) {
			t.Fatal("topics must be sorted by message count")
		}
	}
}
