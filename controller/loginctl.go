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
	"fmt"
	"net/http"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type loginRequest struct {
	NameOrEmail    string `json:"nameOrEmail" binding:"required"`
	PasswordHashed string `json:"passwordHashed" binding:"required"`
}

func loginAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	reqData := &loginRequest{}
	if err := c.BindJSON(reqData); nil != err {
		result.Code = -1
		result.Msg = "parses login request failed"

		return
	}

	user := service.User.GetUserByNameOrEmail(reqData.NameOrEmail)
	if nil == user {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	if user.Password != reqData.PasswordHashed {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, user.BlogID)
	if nil == blogTitleSetting {
		result.Code = -1
		result.Msg = fmt.Sprintf("not found blog title settings [blogID=%d]", user.BlogID)

		return
	}

	pathSetting := service.Setting.GetSetting(model.SettingCategorySystem, model.SettingNameSystemPath, user.BlogID)
	if nil == pathSetting {
		result.Code = -1
		result.Msg = fmt.Sprintf("not found path settings [blogID=%d]", user.BlogID)

		return
	}

	data := map[string]interface{}{}
	data["name"] = user.Name
	data["nickname"] = user.Nickname
	data["blogTitle"] = blogTitleSetting.Value
	data["blogPath"] = util.PathBlogs + pathSetting.Value
	data["role"] = user.Role
	blogs := service.User.GetUserBlogs(user.ID)
	if 1 > len(blogs) {
		result.Code = -1
		result.Msg = fmt.Sprint("not found blog [userID=%d]", user.ID)
	}
	data["blogs"] = blogs
	result.Data = data

	sessionData := &util.SessionData{
		UID:   user.ID,
		UName: user.Name,
		URole: user.Role,
		BID:   user.BlogID,
		BPath: util.PathBlogs + pathSetting.Value,
	}
	if err := sessionData.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}
}

func logoutAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Clear()
	if err := session.Save(); nil != err {
		log.Errorf("saves session failed: " + err.Error())
	}
}
