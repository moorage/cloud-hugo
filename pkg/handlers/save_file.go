package handlers

import (
	"context"
	"encoding/json"
	"log"
	"path/filepath"

	"io/ioutil"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/labstack/echo"
	ghttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type SaveFileReq struct {
	Filename string
	Content  string
}

func (rqm *ReqManager) SaveFile(c echo.Context) error {
	repoPath := filepath.Join(rqm.cfg.BaseDir, rqm.RepoName)
	var sfr SaveFileReq
	if err := c.Bind(&sfr); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// TODO: change this to make this generic and not just save the file in posts
	err := ioutil.WriteFile(filepath.Join(repoPath, "content", "post", sfr.Filename), []byte(sfr.Content), 0644)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = rqm.pubsubWorflow()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}

func (rqm *ReqManager) pubsubWorflow() error {
	output, err := rqm.builder.Build()
	if err != nil {
		return err
	}

	log.Println(output)

	err = rqm.gitClient.CommitAndPush(rqm.cfg.RepoURL, rqm.cfg.UserName, rqm.cfg.UserEmail,
		&ghttp.BasicAuth{
			Username: rqm.cfg.UserName,
			Password: rqm.cfg.AccessToken,
		})

	d, err := json.Marshal(GitMsg{
		GitURL: rqm.cfg.RepoURL,
	})
	_, err = rqm.topic.Publish(context.Background(), &pubsub.Message{Data: d})
	return err
}
