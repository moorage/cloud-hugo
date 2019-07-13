package git

import (
	"log"
	"path/filepath"
	"time"

	"github.com/moorage/cloud-hugo/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// CommitAndPush opens the repository provided in the urlStr and then adds all the files
// then commits and pushes
func (gc *GitClient) CommitAndPush(urlStr, authorName, authorEmail string) error {
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.AddGlob("*")
	if err != nil {
		return err
	}
	commit, err := w.Commit("added changes", &git.CommitOptions{
		Author: &object.Signature{
			Name:  authorName,
			Email: authorEmail,
			When:  time.Now(),
		},
	})

	obj, err := r.CommitObject(commit)
	if err != nil {
		return nil
	}
	log.Println("Commit successful - ", obj.String())
	return r.Push(&git.PushOptions{})
}
