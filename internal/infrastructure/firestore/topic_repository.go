package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/balloon-chat/topic-recommend/internal/domain/model"
	"google.golang.org/api/iterator"
	"time"
)

const (
	topicCollectionKey = "topics"
	createdAtKey       = "createdAt"
	isPrivateKey       = "isPrivate"
)

type FirestoreTopicRepository struct {
	ctx              context.Context
	topicsCollection *firestore.CollectionRef
}

func NewFirestoreTopicRepository(ctx context.Context) (*FirestoreTopicRepository, error) {
	client, err := NewFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FirestoreTopicRepository{
		ctx:              ctx,
		topicsCollection: client.Collection(topicCollectionKey),
	}, nil
}

func (db *FirestoreTopicRepository) GetTopicsOrderByCreatedAt() ([]*model.Topic, error) {
	lastWeek := time.Now().Add(-7 * time.Hour * 24).Unix()
	query := db.topicsCollection.Where(createdAtKey, ">=", lastWeek).Where(isPrivateKey, "==", false).OrderBy("createdAt", firestore.Desc)
	docItr := query.Documents(db.ctx)

	var topics []*model.Topic
	for {
		doc, err := docItr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while getting data from firestore: %v", err)
		}

		var topic model.Topic
		err = doc.DataTo(&topic)
		if err != nil {
			return nil, fmt.Errorf("error while converting data: %v", err)
		}

		topics = append(topics, &topic)
	}

	return topics, nil
}
