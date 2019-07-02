package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/publisher"
)

func baseRouter(engine *gin.Engine) *gin.RouterGroup {

	v1 := engine.Group("/api/v1")
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "OK",
		})
	})
	return v1
}

func main() {
	ctx := context.Background()
	cfg, err := config.NewPublisherConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := publisher.New(ctx, cfg)

	_, err = client.CreateOrInitTopic(cfg.TopicName)

	if err != nil {
		log.Fatalln(err)
	}

	engine := gin.Default()
	baseRouter(engine)
	engine.Run()
}
