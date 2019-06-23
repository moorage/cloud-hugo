package subscr

import (
	"log"

	"context"

	"cloud.google.com/go/pubsub"

	"github.com/moorage/cloud-hugo/pkg/config"
	"google.golang.org/api/option"
)

type SubClient struct {
	*pubsub.Client
}

func New(cfg *config.Config) *SubClient {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cfg.ProjectID, option.WithCredentialsFile(cfg.CredFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &SubClient{
		Client: client,
	}
}
