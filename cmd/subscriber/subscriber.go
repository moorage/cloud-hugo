package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/subscr"
)

const subName = "chugo-run-requests-sub"

func main() {
	ctx := context.Background()
	cfg, err := config.NewSubsciberConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := subscr.New(ctx, cfg)

	topic, err := client.InitTopic(cfg.TopicName)
	if err != nil {
		log.Fatalln(err)
	}
	subscription, err := client.InitSubscription(subName, topic)
	if err != nil {
		log.Fatalln(err)
	}

	manager := subscr.NewManager(cfg)
	// Start worker goroutine.
	go subscribe(subscription, manager)

	fs := http.FileServer(http.Dir(cfg.HostingDir))
	http.Handle("/", fs)

	log.Printf("Listening on %s\n", cfg.SubPort)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.SubPort), nil)
}

func subscribe(subscription *pubsub.Subscription, manager *subscr.Manager) {
	ctx := context.Background()
	err := subscription.Receive(ctx, manager.HandleGitMsg)
	if err != nil {
		log.Fatal(err)
	}
}
