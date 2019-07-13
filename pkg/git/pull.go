package git

import (
	"log"
	"path/filepath"

	"github.com/moorage/cloud-hugo/pkg/utils"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func (gc *GitClient) Pull(urlStr string) error {
	// We instanciate a new repository targeting the given path (the .git folder)
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
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
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err == git.NoErrAlreadyUpToDate {
		// if err is that the repo is up-to-date we ignore it
		return nil
	}
	return err
}

func (gc *GitClient) PullWithAuth(urlStr string, auth http.AuthMethod) error {
	// We instance\iate a new repository targeting the given path (the .git folder)
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
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
	err = w.Pull(&git.PullOptions{RemoteName: "origin", Auth: auth})

	if err == git.NoErrAlreadyUpToDate {
		// if err is that the repo is up-to-date we ignore it
		return nil
	}
	return err
}
