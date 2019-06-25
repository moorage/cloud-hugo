package git

type GitClient struct {
	BaseDir string
}

func NewClient(baseDir string) *GitClient {
	return &GitClient{
		BaseDir: baseDir,
	}
}
