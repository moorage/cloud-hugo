package utils


func TestExtractNameFromGitURL(t *testing.T) {
	sshName := utils.ExtractNameFromGitURL("git@github.com:jenkins-x/jx.git")
	HTTPSName := utils.ExtractNameFromGitURL("https://github.com/jenkins-x/jx.git")

	if sshName != "jx" {
		t.Fatalf("ssh based url's name is incorrect")
	}

	if HTTPSName != "jx" {
		t.Fatalf("https based url's name is incorrect")
	}
}