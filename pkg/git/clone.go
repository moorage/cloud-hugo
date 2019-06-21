package git

import (
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
)

// Clone clones the repository with the given url to the base dir
func (gc *GitClient) Clone(name, url string) (*git.Repository, error) {
	return git.PlainClone(filepath.Join(gc.BaseDir, name), false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

}
