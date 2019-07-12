package utils

import "strings"

func ExtractNameFromGitURL(urlStr string) string {
	url := strings.Split(urlStr, "/")
	gitPath := url[len(url)-1]
	return strings.TrimSuffix(gitPath, ".git")
}
