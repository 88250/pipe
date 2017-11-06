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
	"net/http"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUsersAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	users := []*ConsoleUser{}
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)
	if nil == blogURLSetting {
		log.Errorf("not found blog URL setting [blogID=%d]", session.BID)

		return
	}
	userModels := service.User.GetBlogUsers(session.BID)
	for _, userModel := range userModels {
		users = append(users, &ConsoleUser{
			ID:        userModel.ID,
			Name:      userModel.Name,
			Nickname:  userModel.Nickname,
			Role:      userModel.Role,
			URL:       blogURLSetting.Value + util.PathAuthors + "/" + userModel.Name,
			AvatarURL: userModel.AvatarURL,
		})
	}

	result.Data = users
}
