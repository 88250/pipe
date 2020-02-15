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
