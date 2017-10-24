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
	"github.com/b3log/pipe/theme"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UpdateThemeAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	theme := c.Param("id")

	sessionData := util.GetSession(c)

	settings := []*model.Setting{
		&model.Setting{
			Category: model.SettingCategoryTheme,
			Name:     model.SettingNameThemeName,
			Value:    theme,
			BlogID:   sessionData.BID,
		},
	}
	if err := service.Setting.UpdateSettings(model.SettingCategoryTheme, settings); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetThemesAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)

	currentID := theme.Themes[0]
	themeNameSetting := service.Setting.GetSetting(model.SettingCategoryTheme, model.SettingNameThemeName, sessionData.BID)
	if nil == themeNameSetting {
		log.Error("not found theme name setting")
	} else {
		currentID = themeNameSetting.Value
	}

	themes := []*ConsoleTheme{}
	for _, themeName := range theme.Themes {
		consoleTheme := &ConsoleTheme{
			ID:           themeName,
			Title:        themeName,
			PreviewURL:   "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
			ThumbnailURL: "https://img.hacpai.com/?theme=Finding",
		}

		themes = append(themes, consoleTheme)
	}

	result.Data = map[string]interface{}{
		"currentId": currentID,
		"themes":    themes,
	}
}
