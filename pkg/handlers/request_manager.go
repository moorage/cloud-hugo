package handlers

import (
	"cloud.google.com/go/pubsub"
	"github.com/moorage/cloud-hugo/pkg/builder"
	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/git"
	"github.com/moorage/cloud-hugo/pkg/utils"
)

// GitMsg is the message which a publisher sends when a repo needs to be cloned or pulled
type GitMsg struct {
	GitURL string `json:"git_url"`
}

// ReqManager contains all the necessary context and dependencies
type ReqManager struct {
	cfg       *config.PublisherConfig
	gitClient *git.GitClient
	topic     *pubsub.Topic
	builder   builder.Builder
	RepoName  string
}

func NewReqManager(cfg *config.PublisherConfig,
	gitClient *git.GitClient,
	topic *pubsub.Topic,
	builder builder.Builder) *ReqManager {
	return &ReqManager{
		cfg:       cfg,
		gitClient: gitClient,
		builder:   builder,
		topic:     topic,
		RepoName:  utils.ExtractNameFromGitURL(cfg.RepoURL),
	}
}
