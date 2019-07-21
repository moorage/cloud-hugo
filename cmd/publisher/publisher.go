package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/moorage/cloud-hugo/pkg/builder"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/git"
	"github.com/moorage/cloud-hugo/pkg/handlers"
	"github.com/moorage/cloud-hugo/pkg/publisher"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func baseRouter(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	v1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
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

	rqm := handlers.NewReqManager(cfg)

	// the backend
	e := echo.New()

	v1 := baseRouter(e)
	v1.POST("/save_file", rqm.SaveFile)
	e.Static("/", "./frontend/dist/")
	log.Println("Listening on 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
