package subscr

import (
	"log"

	"context"
	"fmt"

	"cloud.google.com/go/pubsub"

	"github.com/moorage/cloud-hugo/pkg/config"
	"google.golang.org/api/option"
)

// SubClient holds handling of a single topic
type SubClient struct {
	client *pubsub.Client
	cfg    *config.SubscriberConfig
	ctx    context.Context
	topic  *pubsub.Topic
}

func New(ctx context.Context, cfg *config.SubscriberConfig) *SubClient {
	client, err := pubsub.NewClient(ctx, cfg.ProjectID, option.WithCredentialsFile(cfg.CredFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// TODO: provide a options parameter to all functions to provide a different context
	return &SubClient{
		client: client,
		ctx:    ctx,
	}
}

// InitTopic checks if a topic exists and it does then returns it else error out because
// creating a topic is publisher's
func (sc *SubClient) InitTopic(topicName string) (*pubsub.Topic, error) {
	log.Println("Initializing topic " + topicName)
	topic := sc.client.Topic(topicName)

	exists, err := topic.Exists(sc.ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("The %s topic ID doesn't exist. The topic needs to be created by the publisher", topicName)
	}
	sc.topic = topic
	return topic, nil
}

func (sc *SubClient) InitSubscription(subName string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	// Create topic subscription if it does not yet exist.
	var subscription *pubsub.Subscription
	log.Println("Initializing subscription " + subName)
	subscription = sc.client.Subscription(subName)
	exists, err := subscription.Exists(sc.ctx)
	if err != nil {
		return nil, fmt.Errorf("Error checking for subscription: %v", err)
	}
	if !exists {
		log.Println("Creating subscription " + subName + " as it doesn't exist")

		if subscription, err = sc.client.CreateSubscription(sc.ctx, subName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
			return nil, fmt.Errorf("Failed to create subscription: %v", err)
		}
	}

	return subscription, err

}

func (sc *SubClient) GetCurrentTopic() *pubsub.Topic {
	return sc.topic
}
