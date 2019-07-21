package handlers

import (
	"path/filepath"

	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type SaveFileReq struct {
	Filename string
	Content  string
}

func (rqm *ReqManager) SaveFile(c echo.Context) error {
	repoPath := filepath.Join(rqm.cfg.BaseDir, rqm.RepoName)
	var sfr SaveFileReq
	if err := c.Bind(&sfr); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// TODO: change this to make this generic and not just save the file in posts
	err := ioutil.WriteFile(filepath.Join(repoPath, "content", "post", sfr.Filename), []byte(sfr.Content), 0644)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, map[string]string{
		"message": "success",
	})
}
