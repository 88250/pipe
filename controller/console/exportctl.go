// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
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
	"github.com/gin-gonic/gin"
	"github.com/b3log/pipe/util"
	"net/http"
	"github.com/b3log/pipe/service"
	"os"
	"path/filepath"
)

func ExportMarkdownAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if nil == session {
		result.Code = -1
		result.Msg = "please login before export"

		return
	}

	mdFiles:=service.Export.ExportMarkdowns(session.BID)

	tempDir := os.TempDir()
	logger.Trace("temp dir path is [" + tempDir + "]")
	zipFilePath := filepath.Join(tempDir, session.UName+"-export-md.zip")
	zipFile, err := os.Create(zipFilePath)
	if nil != err {
		logger.Errorf("create temp file [" + zipFilePath + "] failed: " + err.Error())
		result.Code = -1
		result.Msg = "create temp file failed"

		return
	}
	util.Zip.Create()

}
