package git

import (
	"log"
	"os"
	"path/filepath"
)

func (gc *GitClient) CloneOrPull(urlStr string) error {
	repoPath := filepath.Join(gc.BaseDir, ExtractNameFromGitURL(urlStr))
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		log.Printf("folder %s doesn't exist hence cloning", repoPath)
		_, err = gc.Clone(urlStr)
		if err != nil {
			return err
		}
	}
	return gc.Pull(urlStr)
}
