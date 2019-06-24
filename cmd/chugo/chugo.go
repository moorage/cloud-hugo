package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/subscr"
)

const subName = "book-worker-sub"

var (
	countMu sync.Mutex
	count   int
)

func main() {
	ctx := context.Background()
	cfg := config.New()

	client := subscr.New(ctx, cfg)

	topic, err := client.InitTopic(cfg.TopicName)
	if err != nil {
		log.Fatalln(err)
	}
	subscription, err := client.InitSubscription(subName, topic)
	if err != nil {
		log.Fatalln(err)
	}
	// Start worker goroutine.
	go subscribe(subscription)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// [END http]
}

func subscribe(subscription *pubsub.Subscription) {
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
