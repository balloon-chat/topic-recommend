package firebase

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	model2 "github.com/balloon-chat/topic-recommend/internal/domain/model"
)

type RecommendTopicDatabase interface {
	SaveRecommendTopics(recommend model2.RecommendTopics) error
}

type FirebaseRecommendTopicDatabase struct {
	ctx          context.Context
	recommendRef *db.Ref
}

func NewFirebaseRecommendTopicDatabase(ctx context.Context) (*FirebaseRecommendTopicDatabase, error) {
	client, err := NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseRecommendTopicDatabase{
		ctx:          ctx,
		recommendRef: client.NewRef("recommend"),
	}, nil
}

func (db FirebaseRecommendTopicDatabase) SaveRecommendTopics(recommend model2.RecommendTopics) error {
	err := db.recommendRef.Set(db.ctx, recommend)
	if err != nil {
		return fmt.Errorf("error while writing data: %v", err)
	}
	return nil
}
