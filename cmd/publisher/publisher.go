package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moorage/cloud-hugo/pkg/builder"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/git"
	"github.com/moorage/cloud-hugo/pkg/publisher"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
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
	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()
	cfg, err := config.NewPublisherConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := publisher.New(ctx, cfg)
	gitClient := git.NewClient(cfg.BaseDir)

	log.Println("Initial Website sync")
	if cfg.AccessToken != "" && cfg.UserName != "" {
		log.Println("Cloning with token")
		err := gitClient.CloneOrPullWithAuth(cfg.RepoURL,
			&http.BasicAuth{
				Username: cfg.UserName,
				Password: cfg.AccessToken,
			})

		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err := gitClient.CloneOrPull(cfg.RepoURL)
		if err != nil {
			log.Fatalln(err)
		}
	}

	_, err = client.CreateOrInitTopic(cfg.TopicName)

	if err != nil {
		log.Fatalln(err)
	}

	hugoBuilder := builder.NewHugoBuilder(cfg)

	output, err := hugoBuilder.Build()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(output)

	err = gitClient.CommitAndPush(cfg.RepoURL, cfg.UserName, cfg.UserEmail,
		&http.BasicAuth{
			Username: cfg.UserName,
			Password: cfg.AccessToken,
		})

	if err != nil {
		log.Fatalln(err)
	}

	// the backend
	engine := gin.Default()
	baseRouter(engine)
	log.Println("Listening on 8080")
	err = engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
