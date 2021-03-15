package database

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/balloon-chat/topic-recommend/src/database/firebase"
	"github.com/balloon-chat/topic-recommend/src/model"
)

type RecommendTopicDatabase interface {
	SaveRecommendTopics(recommend model.RecommendTopics) error
}

type FirebaseRecommendTopicDatabase struct {
	ctx          context.Context
	recommendRef *db.Ref
}

func NewFirebaseRecommendTopicDatabase(ctx context.Context) (*FirebaseRecommendTopicDatabase, error) {
	client, err := firebase.NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseRecommendTopicDatabase{
		ctx:          ctx,
		recommendRef: client.NewRef("recommend"),
	}, nil
}

func (db FirebaseRecommendTopicDatabase) SaveRecommendTopics(recommend model.RecommendTopics) error {
	err := db.recommendRef.Set(db.ctx, recommend)
	if err != nil {
		return fmt.Errorf("error while writing data: %v", err)
	}
	return nil
}
