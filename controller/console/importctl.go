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

// ImportMarkdownAction imports markdown zip file as articles.
func ImportMarkdownAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if 0 == session.UID {
		result.Code = util.CodeErr
		result.Msg = "please login before import"

		return
	}

	form, err := c.MultipartForm()
	if nil != err {
		msg := "parse upload file header failed"
		logger.Errorf(msg + ": " + err.Error())
		result.Code = util.CodeErr
		result.Msg = msg

		return
	}

	file := form.File["file"][0]
	f, err := file.Open()
	if nil != err {
		msg := "open upload file failed"
		logger.Errorf(msg + ": " + err.Error())
		result.Code = util.CodeErr
		result.Msg = msg

		return
	}
	defer f.Close()

	tempDir := os.TempDir()
	logger.Trace("temp dir path is [" + tempDir + "]")
	zipFilePath := filepath.Join(tempDir, session.UName+"-import-md.zip")
	zipFile, err := os.Create(zipFilePath)
	if nil != err {
		logger.Errorf("create temp file [" + zipFilePath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "create temp file failed"

		return
	}
	_, err = io.Copy(zipFile, f)
	if nil != err {
		logger.Errorf("write temp file [" + zipFilePath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "write temp file failed"

		return
	}
	defer zipFile.Close()

	unzipPath := filepath.Join(tempDir, session.UName+"-import-md")
	if err = os.RemoveAll(unzipPath); nil != err {
		logger.Errorf("remove temp dir [" + unzipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "remove temp dir failed"

		return
	}
	if err = os.Mkdir(unzipPath, 0755); nil != err {
		logger.Errorf("make temp dir [" + unzipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "make temp dir failed"

		return
	}
	if err = gulu.Zip.Unzip(zipFilePath, unzipPath); nil != err {
		logger.Errorf("unzip [" + zipFilePath + "] to [" + unzipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "unzip failed"

		return
	}

	logger.Info("importing markdowns [zipFilePath=" + zipFilePath + ", unzipPath=" + unzipPath + "]")

	var filePaths []string
	err = filepath.Walk(unzipPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			filePaths = append(filePaths, path)
		}

		return err
	})
	if nil != err {
		logger.Errorf("read dir [" + unzipPath + "] failed: " + err.Error())
		result.Code = util.CodeErr
		result.Msg = "read dir failed"

		return
	}

	var mdFiles []*service.MarkdownFile
	const (
		bom0 = 0xef
		bom1 = 0xbb
		bom2 = 0xbf
	)
	for _, filePath := range filePaths {
		data, err := ioutil.ReadFile(filePath)
		if nil != err {
			logger.Errorf("read file [" + filePath + "] failed")

			continue
		}

		if len(data) >= 3 && data[0] == bom0 && data[1] == bom1 && data[2] == bom2 {
			data = data[3:]
		}

		mdFile := &service.MarkdownFile{
			Name:    filepath.Base(filePath),
			Path:    filePath,
			Content: string(data),
		}

		mdFiles = append(mdFiles, mdFile)
	}

	service.Import.ImportMarkdowns(mdFiles, session.UID, session.BID)
}
