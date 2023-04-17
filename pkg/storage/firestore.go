package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func FireStore() {
	ctx := context.Background()

	// Replace "my-project-id" with your own project ID
	projectID := "my-project-id"

	// Create a new Firestore client using the default credentials mechanism
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()
}
