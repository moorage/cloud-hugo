package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

const subName = "book-worker-sub"

var (
	countMu sync.Mutex
	count   int

	subscription *pubsub.Subscription
)

func main() {
	ctx := context.Background()
	projectID := "test-project"
	topicID := "my-topic"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	topic := client.Topic(topicID)

	exists, err := topic.Exists(ctx)

	if err != nil {
		log.Fatalln("Failed to get the topic")
	}
	if !exists {
		log.Fatalf("The %s topic ID doesn't exist. The topic needs to be created by the publisher", topicID)
	}

	// Create topic subscription if it does not yet exist.
	subscription = client.Subscription(subName)
	exists, err = subscription.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking for subscription: %v", err)
	}
	if !exists {
		if _, err = client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
		}
	}

	// Start worker goroutine.
	go subscribe()

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// [END http]
}

func subscribe() {
	ctx := context.Background()
	err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var id int64
		if err := json.Unmarshal(msg.Data, &id); err != nil {
			log.Printf("could not decode message data: %#v", msg)
			msg.Ack()
			return
		}

		log.Printf("[ID %d] Processing.", id)

		countMu.Lock()
		count++
		countMu.Unlock()

		msg.Ack()
		log.Printf("[ID %d] ACK", id)
	})
	if err != nil {
		log.Fatal(err)
	}
}
