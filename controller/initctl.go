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

// Package controller is the "controller" layer.
package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func showInitPageAction(c *gin.Context) {
	t, err := template.ParseFiles("console/dist/init/index.html")
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

	platformAdmin := &model.User{
		Name:      session.UName,
		B3Key:     session.UB3Key,
		AvatarURL: session.UAvatar,
	}

	if err := service.Init.InitPlatform(platformAdmin); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, platformAdmin.BlogID)
	if nil == blogURLSetting {
		result.Code = -1
		result.Msg = fmt.Sprintf("not found blog URL settings [blogID=%d]", platformAdmin.BlogID)

		return
	}
	sessionData := &util.SessionData{
		UID:     platformAdmin.ID,
		UName:   platformAdmin.Name,
		UB3Key:  platformAdmin.B3Key,
		URole:   platformAdmin.Role,
		UAvatar: platformAdmin.AvatarURL,
		BID:     platformAdmin.BlogID,
		BURL:    blogURLSetting.Value,
	}
	if err := sessionData.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}
}
