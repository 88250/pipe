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

package controller

import (
	"context"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/api.v7/storage"
	"github.com/satori/go.uuid"
)

var ut = &uploadToken{}

type uploadToken struct {
	token   string
	expired time.Time
}

func uploadAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if nil == session {
		result.Code = -1
		result.Msg = "please login before upload"

		return
	}

	refreshUploadToken()

	header, err := c.FormFile("file")
	if nil != err {
		msg := "parse upload file header failed"
		logger.Errorf(msg + ": " + err.Error())

		result.Code = -1
		result.Msg = msg

		return
	}

	ext := filepath.Ext(header.Filename)
	if "" == ext {
		typ := header.Header.Get("Content-Type")
		exts, _ := mime.ExtensionsByType(typ)
		if 0 < len(exts) {
			ext = exts[0]
		} else {
			ext = "." + strings.Split(typ, "/")[1]
		}
	}

	platformAdmin := service.User.GetPlatformAdmin()
	blogID := getBlogID(c)
	blogAdmin := service.User.GetBlogAdmin(blogID)
	key := "pipe/" + platformAdmin.Name + "/" + blogAdmin.Name + "/" + session.UName + "/" + strings.Replace(uuid.NewV4().String(), "-", "", -1) + ext

	logger.Info(key)

	return
	if err := storage.NewFormUploader(nil).Put(context.Background(), nil, ut.token, key, c.Request.Body, header.Size, nil); nil != err {
		msg := "upload file to storage failed"
		logger.Errorf(msg + ": " + err.Error())

		result.Code = -1
		result.Msg = msg

		return
	}
}

func refreshUploadToken() {
	now := time.Now()
	dur, _ := time.ParseDuration("30m")
	if now.Sub(ut.expired) <= dur {
		return
	}

	uploadTokenResult := util.NewResult()
	if _, _, errs := gorequest.New().Get(util.HacPaiURL+"/apis/qiniu/ut").Timeout(15*time.Second).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).EndStruct(uploadTokenResult); nil != errs {
		logger.Errorf("can't refresh upload token: %s", errs[0])

		return
	}

	if 0 != uploadTokenResult.Code {
		logger.Errorf("can't refresh upload token, get upload token result is [%+v]", uploadTokenResult)

		return
	}

	ut.token = uploadTokenResult.Data.(string)
	ut.expired = now
}
