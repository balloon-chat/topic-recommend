package firebase

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
)

type FirebaseRecommendTopicRepository struct {
	ctx          context.Context
	recommendRef *db.Ref
}

func NewFirebaseRecommendTopicRepository(ctx context.Context) (*FirebaseRecommendTopicRepository, error) {
	client, err := NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseRecommendTopicRepository{
		ctx:          ctx,
		recommendRef: client.NewRef("recommend"),
	}, nil
}

func (db FirebaseRecommendTopicRepository) SaveRecommendTopics(recommend model.RecommendTopics) error {
	err := db.recommendRef.Set(db.ctx, recommend)
	if err != nil {
		return fmt.Errorf("error while writing data: %v", err)
	}
	return nil
}
