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
	"html/template"
	"net/http"

	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func indexCtl(c *gin.Context) {
	theme := "yilia"
	t, err := template.ParseFiles("theme/" + theme + "/index.html")
	if nil != err {
		log.Error("loads theme [" + theme + "] failed: " + err.Error())
		c.String(http.StatusNotFound, "loads theme failed")

		return
	}

	dataModel := map[string]interface{}{}
	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, 1)
	dataModel["i18n"] = i18n.GetMessages(localeSetting.Value)
	dataModel["hi"] = "Index"
	basicSettings := service.Setting.GetAllSettings(1, model.SettingCategoryBasic)
	basicSettingMap := map[string]string{}
	for _, basicSetting := range basicSettings {
		basicSettingMap[basicSetting.Name] = basicSetting.Value
	}
	dataModel["setting"] = basicSettingMap

	t.Execute(c.Writer, dataModel)
}
