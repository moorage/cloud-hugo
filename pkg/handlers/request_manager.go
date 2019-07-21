package handlers

import "github.com/moorage/cloud-hugo/pkg/config"
import "github.com/moorage/cloud-hugo/pkg/utils"

// ReqManager contains all the necessary context and dependencies
type ReqManager struct {
	cfg      *config.PublisherConfig
	RepoName string
}

func NewReqManager(cfg *config.PublisherConfig) *ReqManager {
	return &ReqManager{
		cfg:      cfg,
		RepoName: utils.ExtractNameFromGitURL(cfg.RepoURL),
	}
}
