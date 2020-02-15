// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
