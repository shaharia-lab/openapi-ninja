// Package storage provides the functionality to interact with any data storage
package storage

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"github.com/shahariaazam/openapi-ninja/pkg/config"
)

// NewFireStore build client to communicate with Firestore
func NewFireStore(cfg config.Config) {
	ctx := context.Background()

	// Replace "my-project-id" with your own project ID
	projectID := cfg.GoogleCloudProject

	// Create a new Firestore client using the default credentials mechanism
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()
}
