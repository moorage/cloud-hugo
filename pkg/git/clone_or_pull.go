package git

import (
	"log"
	"os"
	"path/filepath"

	"github.com/moorage/cloud-hugo/pkg/utils"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func (gc *GitClient) CloneOrPull(urlStr string) error {
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		log.Printf("folder %s doesn't exist hence cloning", repoPath)
		_, err = gc.Clone(urlStr)
		if err != nil {
			return err
		}
	}
	return gc.Pull(urlStr)
}

func (gc *GitClient) CloneOrPullWithAuth(urlStr string, auth http.AuthMethod) error {
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		log.Printf("folder %s doesn't exist hence cloning", repoPath)
		_, err = gc.CloneWithAuth(urlStr, auth)
		if err != nil {
			return err
		}
	}
	return gc.PullWithAuth(urlStr, auth)
}
