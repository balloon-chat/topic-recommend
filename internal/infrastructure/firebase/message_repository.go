package firebase

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
)

type FirebaseMessageRepository struct {
	ctx         context.Context
	messagesRef *db.Ref
}

func NewFirebaseMessageRepository(ctx context.Context) (*FirebaseMessageRepository, error) {
	client, err := NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseMessageRepository{
		ctx:         ctx,
		messagesRef: client.NewRef("messages"),
	}, nil
}

func (db FirebaseMessageRepository) GetMessageCountOf(topicId string) (*int, error) {
	var messages map[string]model.Message
	err := db.messagesRef.Child(topicId).Get(db.ctx, &messages)
	if err != nil {
		return nil, fmt.Errorf("error while getting data: %v", err)
	}

	count := len(messages)
	return &count, nil
}
