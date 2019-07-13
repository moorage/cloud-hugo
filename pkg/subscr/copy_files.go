package subscr

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/moorage/cloud-hugo/pkg/utils"
	"github.com/otiai10/copy"
)

// CopyFilesForHosting copies the files for hosting, it replaces the files as needed
// TODO: opitimize this to only copy files which have changed
func (manager *Manager) CopyFilesForHosting(gitURL string) error {
	name := utils.ExtractNameFromGitURL(gitURL)
	// public folder contains all the built files
	publicAssetDir := filepath.Join(manager.cfg.BaseDir, name, "public")
	if _, err := os.Stat(publicAssetDir); os.IsNotExist(err) {
		return errors.New("public dir doesn't exist")
	}

	return copy.Copy(publicAssetDir, manager.cfg.HostingDir)
}
