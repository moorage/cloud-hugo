package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/git"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"cloud.google.com/go/pubsub"
)

// GitMsg is the message which a publisher sends when a repo needs to be cloned or pulled
type GitMsg struct {
	GitURL string `json:"git_url"`
}

// Manager provides handlers for handling all kinds of message which can be sent by the publisher
type Manager struct {
	cfg       *config.SubscriberConfig
	gitClient *git.GitClient
}

// NewManager initializes a Manager
func NewManager(cfg *config.SubscriberConfig) *Manager {
	return &Manager{
		cfg:       cfg,
		gitClient: git.NewClient(cfg.BaseDir),
	}
}

// HandleGitMsg  handles the git message sent by the publisher
func (hdlr *Manager) HandleGitMsg(ctx context.Context, msg *pubsub.Message) {
	var gitMsg GitMsg
	if err := json.Unmarshal(msg.Data, &gitMsg); err != nil {
		log.Printf("could not decode message data: %#v", msg)
		msg.Ack()
		return
	}

	log.Printf("[Msg %+v] Processing.", gitMsg)
	if hdlr.cfg.AccessToken != "" || hdlr.cfg.UserName != "" {
		hdlr.gitClient.CloneOrPullWithAuth(gitMsg.GitURL,
			&http.BasicAuth{
				Username: hdlr.cfg.UserName, // yes, this can be anything except an empty string
				Password: hdlr.cfg.AccessToken,
			})
	} else {
		err := hdlr.gitClient.CloneOrPull(gitMsg.GitURL)
		if err != nil {
			log.Println("There was an error while processing ", err.Error())
		}
	}

	msg.Ack()
	log.Printf("[Msg] ACK")
}
