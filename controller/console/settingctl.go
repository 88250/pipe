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

// Package console defines console controllers.
package console

import (
	"net/http"
	"strconv"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetBasicSettingsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)
	basicSettings := service.Setting.GetAllSettings(sessionData.BID, model.SettingCategoryBasic)
	data := map[string]interface{}{}
	for _, setting := range basicSettings {
		if model.SettingNameBasicCommentable == setting.Name {
			v, err := strconv.ParseBool(setting.Value)
			if nil != err {
				log.Errorf("value of basic setting [name=%s] must be \"true\" or \"false\"", setting.Name)
				data[setting.Name] = true
			} else {
				data[setting.Name] = v
			}
		} else {
			data[setting.Name] = setting.Value
		}
	}
	result.Data = data
}

func UpdateSettingsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update preferences request failed"

		return
	}

	sessionData := util.GetSession(c)
	prefs := []*model.Setting{}
	for k, v := range args {
		var value interface{}
		switch v.(type) {
		case float64:
			value = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		default:
			value = v.(string)
		}

		pref := &model.Setting{
			Category: model.SettingCategoryPreference,
			BlogID:   sessionData.BID,
			Name:     k,
			Value:    value.(string),
		}

		prefs = append(prefs, pref)
	}

	if err := service.Setting.UpdatePreferences(prefs); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
