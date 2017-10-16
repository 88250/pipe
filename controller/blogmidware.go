// Solo.go - A small and beautiful blogging platform written in golang.
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
	"net/http"
	"strconv"
	"time"

	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func resolveBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		if "" == username {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}
		blogAdmin := service.User.GetUserByName(username)
		if nil == blogAdmin {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}
		c.Set("blogAdmin", blogAdmin)

		fillCommon(c, &DataModel{})

		c.Next()
	}
}

type DataModel map[string]interface{}

func fillCommon(c *gin.Context, dataModel *DataModel) {
	if "dev" == util.Conf.RuntimeMode {
		i18n.Load()
	}
	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, 1)
	(*dataModel)["i18n"] = i18n.GetMessages(localeSetting.Value)

	settings := service.Setting.GetAllSettings(1)
	settingMap := map[string]string{}
	for _, setting := range settings {
		settingMap[setting.Name] = setting.Value
	}
	settingMap[model.SettingNameSystemPath] = util.PathBlogs + settingMap[model.SettingNameSystemPath]
	(*dataModel)["setting"] = settingMap
	statistics := service.Statistic.GetAllStatistics(1)
	statisticMap := map[string]int{}
	for _, statistic := range statistics {
		count, err := strconv.Atoi(statistic.Value)
		if nil != err {
			log.Errorf("statistic [%s] should be an integer, actual is [%v]", statistic.Name, statistic.Value)
		}
		statisticMap[statistic.Name] = count
	}
	(*dataModel)["statistic"] = statisticMap
	(*dataModel)["title"] = settingMap["basicBlogTitle"]
	(*dataModel)["metaKeywords"] = settingMap["basicMetaKeywords"]
	(*dataModel)["metaDescription"] = settingMap["basicMetaDescription"]
	(*dataModel)["conf"] = util.Conf
	(*dataModel)["year"] = time.Now().Year()

	(*dataModel)["username"] = ""
	session := util.GetSession(c)
	if nil != session {
		(*dataModel)["username"] = session.UName
	}
	(*dataModel)["userCount"] = len(service.User.GetBlogUsers(1))

	navigations := service.Navigation.GetNavigations(1)
	(*dataModel)["navigations"] = navigations

	c.Set("dataModel", dataModel)
}
