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
	"github.com/gin-gonic/gin"
)

type Status struct {
	*service.PlatformStatus

	Name      string              `json:"name"`
	Nickname  string              `json:"nickname"`
	AvatarURL string              `json:"avatarURL"`
	BlogTitle string              `json:"blogTitle"`
	BlogURL   string              `json:"blogURL"`
	Role      int                 `json:"role"`
	Blogs     []*service.UserBlog `json:"blogs"`
}

func getStatusAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	platformStatus, err := service.Init.Status()
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	data := &Status{
		PlatformStatus: platformStatus,
	}

	session := util.GetSession(c)
	if nil != session {
		user := service.User.GetUser(session.UID)
		if nil == user {
			result.Code = -1
			result.Msg = fmt.Sprintf("not found user [userID=%d]", session.UID)

			return
		}
		data.Name = user.Name
		data.Nickname = user.Nickname
		data.AvatarURL = user.AvatarURL
		data.Role = user.Role

		blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, user.BlogID)
		if nil == blogTitleSetting {
			result.Code = -1
			result.Msg = fmt.Sprintf("not found blog title settings [blogID=%d]", user.BlogID)

			return
		}
		data.BlogTitle = blogTitleSetting.Value

		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, user.BlogID)
		if nil == blogURLSetting {
			result.Code = -1
			result.Msg = fmt.Sprintf("not found blog URL settings [blogID=%d]", user.BlogID)

			return
		}
		data.BlogURL = blogURLSetting.Value
		data.Blogs = service.User.GetUserBlogs(user.ID)
	}

	result.Data = data
}
