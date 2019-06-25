package git

import (
	"log"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
)

func (gc *GitClient) Pull(urlStr string) error {
	// We instance\iate a new repository targeting the given path (the .git folder)
	repoPath := filepath.Join(gc.BaseDir, ExtractNameFromGitURL(urlStr))
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	// Pull the latest changes from the origin remote and merge into the current branch
	log.Println("git pull", urlStr)
	return w.Pull(&git.PullOptions{RemoteName: "origin"})
}
