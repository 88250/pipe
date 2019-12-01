// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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
	"net/http"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// AddUserAction adds a user.
func AddUserAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add user request failed"

		return
	}

	name := arg["name"].(string)
	user := service.User.GetUserByName(name)
	if nil == user {
		result.Code = util.CodeErr
		result.Msg = "the user should login first"

		return
	}

	session := util.GetSession(c)
	if err := service.User.AddUserToBlog(user.ID, session.BID); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}
}

// GetUsersAction gets users.
func GetUsersAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	var users []*ConsoleUser
	userModels, pagination := service.User.GetBlogUsers(util.GetPage(c), session.BID)
	for _, userModel := range userModels {
		userBlog := service.User.GetUserBlog(userModel.ID, session.BID)
		users = append(users, &ConsoleUser{
			ID:           userModel.ID,
			Name:         userModel.Name,
			Nickname:     userModel.Nickname,
			Role:         userBlog.UserRole,
			URL:          blogURLSetting.Value + util.PathAuthors + "/" + userModel.Name,
			AvatarURL:    userModel.AvatarURL,
			ArticleCount: userBlog.UserArticleCount,
		})
	}

	result.Data = map[string]interface{}{
		"users":      users,
		"pagination": pagination,
	}
}
