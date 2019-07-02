package publisher

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"

	"github.com/moorage/cloud-hugo/pkg/config"
	"google.golang.org/api/option"
)

// PubClient holds handling of a single topic
type PubClient struct {
	client *pubsub.Client
	cfg    *config.PublisherConfig
	ctx    context.Context
	topic  *pubsub.Topic
}

// New creates a publisher client
func New(ctx context.Context, cfg *config.PublisherConfig) *PubClient {
	client, err := pubsub.NewClient(ctx, cfg.ProjectID, option.WithCredentialsFile(cfg.CredFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &PubClient{
		client: client,
		ctx:    ctx,
	}
}

// CreateOrInitTopic checks if a topic exists and it does then returns it else creates it
func (pcl *PubClient) CreateOrInitTopic(topicName string) (*pubsub.Topic, error) {
	if pcl.topic != nil {
		return pcl.topic, nil
	}
	log.Println("Initializing topic " + topicName)
	topic := pcl.client.Topic(topicName)

	exists, err := topic.Exists(pcl.ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		topic, err := pcl.client.CreateTopic(pcl.ctx, topicName)
		if err != nil {
			return nil, err
		}
		pcl.topic = topic
		return topic, nil
	}
	pcl.topic = topic
	return topic, nil
}
