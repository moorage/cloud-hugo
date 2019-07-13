package subscr

import "github.com/otiai10/copy"



// CopyFilesForHosting copies the files for hosting, it replaces the files as needed
// TODO: opitimize this to only copy files which have changed
func CopyFilesForHosting(baseDir string, hostingDir string) error {
	return copy.Copy(baseDir, hostingDir)
}
