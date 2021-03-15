package database

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/balloon-chat/topic-recommend/src/database/firebase"
	"github.com/balloon-chat/topic-recommend/src/model"
	"time"
)

type TopicDatabase interface {
	GetTopicsOrderByCreatedAt() ([]*model.Topic, error)
}

type FirebaseTopicDatabase struct {
	ctx       context.Context
	topicsRef *db.Ref
}

func NewFirebaseTopicDatabase(ctx context.Context) (*FirebaseTopicDatabase, error) {
	client, err := firebase.NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseTopicDatabase{
		ctx:       ctx,
		topicsRef: client.NewRef("topics"),
	}, nil
}

func (db *FirebaseTopicDatabase) GetTopicsOrderByCreatedAt() ([]*model.Topic, error) {
	lastWeek := time.Now().Add(-7 * time.Hour * 24).Unix()
	results, err := db.topicsRef.OrderByChild("createdAt").StartAt(lastWeek * 1000).LimitToFirst(100).GetOrdered(db.ctx)
	if err != nil {
		return nil, fmt.Errorf("error while querying topics: %v", err)
	}

	var topics []*model.Topic
	for _, r := range results {
		var topic model.Topic
		if err := r.Unmarshal(&topic); err != nil {
			return nil, fmt.Errorf("error while unmarshaling result: %v", err)
		}
		topics = append(topics, &topic)
	}

	return topics, nil
}
