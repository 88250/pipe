// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package console

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/88250/gulu"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// ExportMarkdownAction exports articles as markdown zip file.
func ExportMarkdownAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if 0 == session.UID {
		result.Code = util.CodeErr
		result.Msg = "please login before export"

		return
	}

	tempDir := os.TempDir()
	logger.Trace("temp dir path is [" + tempDir + "]")
	zipFilePath := filepath.Join(tempDir, session.UName+"-export-md.zip")
	zipFile, err := gulu.Zip.Create(zipFilePath)
	if nil != err {
		logger.Errorf("create zip file [" + zipFilePath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "create zip file failed"

		return
	}

	c.Header("Content-Disposition", "attachment; filename="+session.UName+"-export-md.zip")
	c.Header("Content-Type", "application/zip")

	mdFiles := service.Export.ExportMarkdowns(session.BID)
	if 1 > len(mdFiles) {
		zipFile.Close()
		file, err := os.Open(zipFilePath)
		if nil != err {
			logger.Errorf("open zip file [" + zipFilePath + " failed: " + err.Error())
			result.Code = util.CodeErr
			result.Msg = "open zip file failed"

			return
		}
		defer file.Close()

		io.Copy(c.Writer, file)

		return
	}

	zipPath := filepath.Join(tempDir, session.UName+"-export-md")
	if err = os.RemoveAll(zipPath); nil != err {
		logger.Errorf("remove temp dir [" + zipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "remove temp dir failed"

		return
	}
	if err = os.Mkdir(zipPath, 0755); nil != err {
		logger.Errorf("make temp dir [" + zipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "make temp dir failed"

		return
	}
	for _, mdFile := range mdFiles {
		filename := filepath.Join(zipPath, mdFile.Name+".md")
		if err := ioutil.WriteFile(filename, []byte(mdFile.Content), 0644); nil != err {
			logger.Errorf("write file [" + filename + "] failed: " + err.Error())
		}
	}

	zipFile.AddDirectory(session.UName+"-export-md", zipPath)
	if err := zipFile.Close(); nil != err {
		logger.Errorf("zip failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "zip failed"

		return
	}
	file, err := os.Open(zipFilePath)
	if nil != err {
		logger.Errorf("open zip file [" + zipFilePath + " failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "open zip file failed"

		return
	}
	defer file.Close()

	io.Copy(c.Writer, file)
}
