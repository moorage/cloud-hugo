package main

import (
	"context"
	"log"

	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/publisher"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewSubsciberConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := publisher.New(ctx, cfg)

	topic, err := client.CreateOrInitTopic(cfg.TopicName)
}
