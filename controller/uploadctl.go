// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"bytes"
	"context"
	"io/ioutil"
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
	domain  string
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

	form, err := c.MultipartForm()
	if nil != err {
		msg := "parse upload file header failed"
		logger.Errorf(msg + ": " + err.Error())

		result.Code = -1
		result.Msg = msg

		return
	}

	platformAdmin := service.User.GetPlatformAdmin()
	blogAdminName := session.UName
	if strings.Contains(c.Request.URL.Path, util.PathBlogs) {
		blogID := getBlogID(c)
		blogAdmin := service.User.GetBlogAdmin(blogID)
		blogAdminName = blogAdmin.Name
	}
	files := form.File["file[]"]

	errFiles := []string{}
	succMap := map[string]string{}
	for _, file := range files {
		ext := filepath.Ext(file.Filename)
		if "" == ext {
			typ := file.Header.Get("Content-Type")
			exts, _ := mime.ExtensionsByType(typ)
			if 0 < len(exts) {
				ext = exts[0]
			} else {
				ext = "." + strings.Split(typ, "/")[1]
			}
		}

		f, err := file.Open()
		if nil != err {
			errFiles = append(errFiles, file.Filename)

			continue
		}
		defer f.Close()

		data, err := ioutil.ReadAll(f)
		if nil != err {
			errFiles = append(errFiles, file.Filename)

			continue
		}

		key := "pipe/" + platformAdmin.Name + "/" + blogAdminName + "/" + session.UName + "/" + strings.Replace(uuid.NewV4().String(), "-", "", -1) + ext

		uploadRet := &storage.PutRet{}
		if err := storage.NewFormUploader(nil).Put(context.Background(), uploadRet, ut.token, key, bytes.NewReader(data), int64(len(data)), nil); nil != err {
			logger.Errorf("upload file to storage failed: " + err.Error())
			errFiles = append(errFiles, file.Filename)

			continue
		}

		succMap[file.Filename] = ut.domain + "/" + uploadRet.Key
	}

	data := map[string]interface{}{}
	data["succMap"] = succMap
	data["errFiles"] = errFiles
	result.Data = data
}

func fetchUploadAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if nil == session {
		result.Code = -1
		result.Msg = "please login before upload"

		return
	}

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses fetch upload request failed"

		return
	}
	fileURL := arg["url"].(string)
	timeout := time.Duration(5 * time.Second)
	resp, fileData, errs := gorequest.New().Get(fileURL).Timeout(timeout).
		Retry(3, 1*time.Second, http.StatusInternalServerError).EndBytes()
	if nil != errs {
		result.Code = -1
		result.Msg = "get data failed"

		return
	}

	typ := resp.Header.Get("content-type")
	if !strings.Contains(typ, "image") {
		client := http.Client{
			Timeout: timeout,
		}
		res, err := client.Get(fileURL)
		if nil != err {
			result.Code = -1
			result.Msg = "get data failed"

			return
		}
		defer res.Body.Close()
		fileData, err = ioutil.ReadAll(res.Body)
		if nil != err {
			result.Code = -1
			result.Msg = "get data failed"

			return
		}
		typ = res.Header.Get("content-type")
	}

	refreshUploadToken()

	platformAdmin := service.User.GetPlatformAdmin()
	blogID := getBlogID(c)
	blogAdmin := service.User.GetBlogAdmin(blogID)

	exts, _ := mime.ExtensionsByType(typ)
	ext := ""
	if 0 < len(exts) {
		ext = exts[0]
	} else {
		ext = "." + strings.Split(typ, "/")[1]
	}

	key := "pipe/" + platformAdmin.Name + "/" + blogAdmin.Name + "/" + session.UName + "/e/" + strings.Replace(uuid.NewV4().String(), "-", "", -1) + ext

	uploadRet := &storage.PutRet{}
	if err := storage.NewFormUploader(nil).Put(context.Background(), uploadRet, ut.token, key, bytes.NewReader(fileData), int64(len(fileData)), nil); nil != err {
		msg := "upload file to storage failed: " + err.Error()
		logger.Errorf(msg)

		result.Code = -1
		result.Msg = msg

		return
	}

	data := map[string]interface{}{}
	data["url"] = ut.domain + "/" + uploadRet.Key
	data["originalURL"] = fileURL
	result.Data = data
}

func refreshUploadToken() {
	now := time.Now()
	dur, _ := time.ParseDuration("30m")
	if now.Sub(ut.expired) <= dur {
		return
	}

	uploadTokenResult := util.NewResult()
	if _, _, errs := gorequest.New().Get(util.HacPaiURL+"/apis/qiniu/ut").Timeout(15*time.Second).
		Retry(3, time.Second, http.StatusInternalServerError).EndStruct(uploadTokenResult); nil != errs {
		logger.Errorf("can't refresh upload token: %s", errs[0])

		return
	}

	if 0 != uploadTokenResult.Code {
		logger.Errorf("can't refresh upload token, get upload token result is [%+v]", uploadTokenResult)

		return
	}

	data := uploadTokenResult.Data.(map[string]interface{})
	ut.token = data["token"].(string)
	ut.domain = data["domain"].(string)
	ut.expired = now
}
