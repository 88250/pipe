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

// Package controller is the "controller" layer.
package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func showInitPageAction(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(util.Conf.StaticRoot, "console/dist/init/index.html")))
	if nil != err {
		logger.Errorf("load init page failed: " + err.Error())
		c.String(http.StatusNotFound, "load init page failed")

		return
	}

	t.Execute(c.Writer, nil)
}

func initAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	if nil == session {
		result.Code = -1
		result.Msg = "session is nil"

		return
	}

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses init request failed"

		return
	}
	b3key := strings.TrimSpace(arg["b3key"].(string))
	if "" == b3key {
		result.Code = -1
		result.Msg = "B3 key cant' be empty"

		return
	}
	if 20 < len(b3key) {
		result.Code = -1
		result.Msg = "B3 key should less then 20 characters"

		return
	}

	checkResult := util.NewResult()
	request := gorequest.New()
	_, _, errs := request.Post(util.HacPaiURL+"/apis/check-b3key").Send(map[string]interface{}{
		"userName":  session.UName,
		"userB3Key": b3key,
	}).Set("user-agent", util.UserAgent).Timeout(30*time.Second).
		Retry(3, 5*time.Second, http.StatusInternalServerError).EndStruct(checkResult)
	if nil != errs {
		logger.Errorf("check b3 key failed: %s", errs)
		result.Code = -1
		result.Msg = "check b3 key failed"

		return
	}

	if 0 != checkResult.Code {
		result.Code = -1
		result.Msg = "B3 key is not match"

		return
	}

	platformAdmin := &model.User{
		Name:      session.UName,
		B3Key:     b3key,
		AvatarURL: session.UAvatar,
	}

	if err := service.Init.InitPlatform(platformAdmin); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, 1)
	if nil == blogURLSetting {
		result.Code = -1
		result.Msg = fmt.Sprintf("not found blog URL settings [blogID=%d]", 1)

		return
	}
	sessionData := &util.SessionData{
		UID:     platformAdmin.ID,
		UName:   platformAdmin.Name,
		UB3Key:  platformAdmin.B3Key,
		URole:   model.UserRoleBlogAdmin,
		UAvatar: platformAdmin.AvatarURL,
		BID:     1,
		BURL:    blogURLSetting.Value,
	}
	if err := sessionData.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}
}
