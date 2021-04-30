package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
)

var client *firestore.Client

func NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	if client == nil {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error initializing firebase app: %v", err)
		}

		client, err = app.Firestore(ctx)
		if err != nil {
			return nil, fmt.Errorf("error initializing firestore: %v", err)
		}
	}

	return client, nil
}
