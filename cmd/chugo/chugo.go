// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample pubsub_worker demonstrates the use of the Cloud Pub/Sub API to communicate between two modules.
// See https://cloud.google.com/go/getting-started/using-pub-sub
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
