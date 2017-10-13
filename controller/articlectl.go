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

// Package controller is the "controller" layer.
package controller

import (
	"net/http"

	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

func showArticleAction(c *gin.Context) {
	dataModel := DataModel{}

	fillCommon(&dataModel)
	c.HTML(http.StatusOK, "index.html", dataModel)
}

func fillCommon(dataModel *DataModel) {
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
	(*dataModel)["setting"] = settingMap
	(*dataModel)["title"] = settingMap["basicBlogTitle"]
	(*dataModel)["metaKeywords"] = settingMap["basicMetaKeywords"]
	(*dataModel)["metaDescription"] = settingMap["basicMetaDescription"]
	(*dataModel)["conf"] = util.Conf

	navigations := service.Navigation.GetNavigations(1)
	(*dataModel)["navigations"] = navigations
}
