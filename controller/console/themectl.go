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
	"github.com/88250/pipe/theme"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// UpdateThemeAction updates theme.
func UpdateThemeAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	theme := c.Param("id")
	session := util.GetSession(c)

	settings := []*model.Setting{
		{
			Category: model.SettingCategoryTheme,
			Name:     model.SettingNameThemeName,
			Value:    theme,
			BlogID:   session.BID,
		},
	}
	if err := service.Setting.UpdateSettings(model.SettingCategoryTheme, settings, session.BID); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}

// GetThemesAction gets themes.
func GetThemesAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	currentID := theme.Themes[0]
	themeNameSetting := service.Setting.GetSetting(model.SettingCategoryTheme, model.SettingNameThemeName, session.BID)
	if nil == themeNameSetting {
		logger.Errorf("not found theme name setting")
	} else {
		currentID = themeNameSetting.Value
	}

	var themes []*ConsoleTheme
	for _, themeName := range theme.Themes {
		consoleTheme := &ConsoleTheme{
			Name:         themeName,
			ThumbnailURL: model.Conf.Server + "/theme/x/" + themeName + "/thumbnail.jpg",
		}

		themes = append(themes, consoleTheme)
	}

	result.Data = map[string]interface{}{
		"currentId": currentID,
		"themes":    themes,
	}
}
