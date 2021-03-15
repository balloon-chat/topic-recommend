package database

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/balloon-chat/topic-recommend/src/database/firebase"
	"github.com/balloon-chat/topic-recommend/src/model"
)

type MessageDatabase interface {
	GetMessageCountOf(topicId string) (*int, error)
}

type FirebaseMessageDatabase struct {
	ctx         context.Context
	messagesRef *db.Ref
}

func NewFirebaseMessageDatabase(ctx context.Context) (*FirebaseMessageDatabase, error) {
	client, err := firebase.NewFirebaseDatabaseClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseMessageDatabase{
		ctx:         ctx,
		messagesRef: client.NewRef("messages"),
	}, nil
}

func (db FirebaseMessageDatabase) GetMessageCountOf(topicId string) (*int, error) {
	var messages map[string]model.Message
	err := db.messagesRef.Child(topicId).Get(db.ctx, &messages)
	if err != nil {
		return nil, fmt.Errorf("error while getting data: %v", err)
	}

	count := len(messages)
	return &count, nil
}
