package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
)

// FileInfo provides the file name and that its a dir or not
type FileInfo struct {
	Name  string
	IsDir bool
}

type listDirResp struct {
	Data []FileInfo
}

// ListDir provides a json representation of the current Directory
func (rqm *ReqManager) ListDir(c echo.Context) error {
	repoPath := filepath.Join(rqm.cfg.BaseDir, rqm.RepoName)
	fileInfos, err := ioutil.ReadDir(repoPath)
	if err != nil {
		return err
	}
	var files []FileInfo
	for _, info := range fileInfos {
		// we ignore the .git folder
		if info.Name() == ".git" {
			continue
		}
		files = append(files, FileInfo{
			Name:  info.Name(),
			IsDir: info.IsDir(),
		})
	}
	return c.JSON(http.StatusOK, &listDirResp{
		Data: files,
	})
}
