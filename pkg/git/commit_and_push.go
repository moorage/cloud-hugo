package git

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/moorage/cloud-hugo/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// CommitAndPush opens the repository provided in the urlStr and then adds all the files
// then commits and pushes
func (gc *GitClient) CommitAndPush(urlStr, authorName, authorEmail string, auth http.AuthMethod) error {
	repoPath := filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr))
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}
	status, err := w.Status()

	if err != nil {
		return err
	}

	// no changes are present locally
	if status.IsClean() {
		log.Println("No changes locally")
		return nil
	}

	for file, fstat := range status {
		fmt.Println(file)
		// if the file is anything but unmodified we add it
		if fstat.Worktree != git.Unmodified {
			_, err := w.Add(file)
			if err != nil {
				// TODO: usually we should ignore it but for now we error out
				return err
			}
		}

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
	return r.Push(&git.PushOptions{
		Auth: auth,
	})
}
