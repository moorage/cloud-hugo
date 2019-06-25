package git

import (
	"os"
)

type GitClient struct {
	BaseDir string
}

func NewClient(baseDir string) *GitClient {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		os.MkdirAll(baseDir, 0777)
	}
	return &GitClient{
		BaseDir: baseDir,
	}
}
