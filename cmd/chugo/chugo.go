package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/handlers"

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

	manager := handlers.NewManager(cfg)
	// Start worker goroutine.
	go subscribe(subscription, manager)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// [END http]
}

func subscribe(subscription *pubsub.Subscription, manager *handlers.Manager) {
	ctx := context.Background()
	err := subscription.Receive(ctx, manager.HandleGitMsg)
	if err != nil {
		log.Fatal(err)
	}
}
