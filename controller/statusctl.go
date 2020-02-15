// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"net/http"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Status represents platform status and blog info.
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
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	platformStatus, err := service.Init.Status()
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	data := &Status{
		PlatformStatus: platformStatus,
	}

	session := util.GetSession(c)
	if 0 != session.UID {
		user := service.User.GetUser(session.UID)
		if nil == user {
			session := sessions.Default(c)
			session.Options(sessions.Options{
				Path:   "/",
				MaxAge: -1,
			})
			session.Clear()
			if err := session.Save(); nil != err {
				logger.Errorf("saves session failed: " + err.Error())
			}

			return
		}

		data.Name = user.Name
		data.Nickname = user.Nickname
		data.AvatarURL = user.AvatarURL
		data.Role = model.UserRoleBlogAdmin

		if model.UserRoleNoLogin != session.URole && platformStatus.Inited {
			ownBlog := service.User.GetOwnBlog(user.ID)
			if nil != ownBlog {
				data.BlogTitle = ownBlog.Title
				data.BlogURL = ownBlog.URL
			}
			data.Blogs = service.User.GetUserBlogs(user.ID)
		}
	}

	result.Data = data
}
