package git

import (
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	git "gopkg.in/src-d/go-git.v4"
)

// Clone clones the repository with the given url to the base dir
func (gc *GitClient) Clone(urlStr string) (*git.Repository, error) {
	return gc.CloneWithAuth(urlStr, nil)
}

// CloneWithAuth clones the repository with the given url to the base dir using auth
func (gc *GitClient) CloneWithAuth(urlStr string, auth http.AuthMethod) (*git.Repository, error) {
	return git.PlainClone(filepath.Join(gc.BaseDir, ExtractNameFromGitURL(urlStr)), false, &git.CloneOptions{
		URL:               urlStr,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})
}

func ExtractNameFromGitURL(urlStr string) string {
	url := strings.Split(urlStr, "/")
	gitPath := url[len(url)-1]
	return strings.TrimSuffix(gitPath, ".git")
}
