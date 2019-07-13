package git

import (
	"path/filepath"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"github.com/moorage/cloud-hugo/pkg/utils"

	git "gopkg.in/src-d/go-git.v4"
)

// Clone clones the repository with the given url to the base dir
func (gc *GitClient) Clone(urlStr string) (*git.Repository, error) {
	return gc.CloneWithAuth(urlStr, nil)
}

// CloneWithAuth clones the repository with the given url to the base dir using auth
func (gc *GitClient) CloneWithAuth(urlStr string, auth http.AuthMethod) (*git.Repository, error) {
	return git.PlainClone(filepath.Join(gc.BaseDir, utils.ExtractNameFromGitURL(urlStr)), false, &git.CloneOptions{
		URL:               urlStr,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})
}
