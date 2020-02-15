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
