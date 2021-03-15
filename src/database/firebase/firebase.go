package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"fmt"
	"os"
)

var client *db.Client

const (
	FirebaseDatabaseUrlKey = "FIREBASE_DATABASE_URL"
)

func NewFirebaseDatabaseClient(ctx context.Context) (*db.Client, error) {
	dbUrl := os.Getenv(FirebaseDatabaseUrlKey)
	if len(dbUrl) == 0 {
		return nil, fmt.Errorf("set environment %s", FirebaseDatabaseUrlKey)
	}
	if client == nil {
		conf := &firebase.Config{
			DatabaseURL: dbUrl,
		}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			return nil, fmt.Errorf("error initializing firebase app: %v", err)
		}

		client, err = app.Database(ctx)
		if err != nil {
			return nil, fmt.Errorf("error initializing firebase realtime database: %v", err)
		}
	}

	return client, nil
}
