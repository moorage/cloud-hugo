package git
import (
	"testing"
)

func TestExtractNameFromGitURL(t *testing.T) {
	sshName := ExtractNameFromGitURL("git@github.com:jenkins-x/jx.git")
	HTTPSName := ExtractNameFromGitURL("https://github.com/jenkins-x/jx.git")

	if sshName != "jx" {
		t.Fatalf("ssh based url's name is incorrect")
	}

	if HTTPSName != "jx" {
		t.Fatalf("https based url's name is incorrect")
	}
}