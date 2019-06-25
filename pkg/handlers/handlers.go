package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/git"

	"cloud.google.com/go/pubsub"
)

// GitMsg is the message which a publisher sends when a repo needs to be cloned or pulled
type GitMsg struct {
	GitURL string `json:"git_url"`
}

type Manager struct {
	cfg       *config.Config
	gitClient *git.GitClient
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		cfg:       cfg,
		gitClient: git.NewClient(cfg.BaseDir),
	}
}

func (hdlr *Manager) HandleGitMsg(ctx context.Context, msg *pubsub.Message) {
	var gitMsg GitMsg
	if err := json.Unmarshal(msg.Data, &gitMsg); err != nil {
		log.Printf("could not decode message data: %#v", msg)
		msg.Ack()
		return
	}

	log.Printf("[Msg %+v] Processing.", gitMsg)
	err := hdlr.gitClient.CloneOrPull(gitMsg.GitURL)
	if err != nil {
		log.Println("There was an error while processing ", err.Error())
	}
	msg.Ack()
	log.Printf("[Msg] ACK")
}
