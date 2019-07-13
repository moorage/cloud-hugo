package builder

import (
	"os/exec"
	"path/filepath"

	"github.com/moorage/cloud-hugo/pkg/config"
	"github.com/moorage/cloud-hugo/pkg/utils"
)

// Builder is the interface implemented by all builders
type Builder interface {
	Build() (string, error)
}

// HugoBuilder is used to build hugo static sites
type HugoBuilder struct {
	// the hugo builder is going to be building on the publisher side
	cfg *config.PublisherConfig
}

func NewHugoBuilder(cfg *config.PublisherConfig) *HugoBuilder {
	return &HugoBuilder{
		cfg: cfg,
	}
}

func (hb *HugoBuilder) Build() (string, error) {
	cmd := exec.Command("hugo")
	name := utils.ExtractNameFromGitURL(hb.cfg.RepoURL)
	cmd.Dir = filepath.Join(hb.cfg.BaseDir, name)

	output, err := cmd.CombinedOutput()
	return string(output), err
}
