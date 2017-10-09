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

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

func UpdatePreferencesCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)

	prefs := []*model.Setting{}
	if err := c.BindJSON(&prefs); nil != err {
		result.Code = -1
		result.Msg = "parses update preferences request failed"

		return
	}

	for _, pref := range prefs {
		pref.Category = model.SettingCategoryPreference
		pref.BlogID = sessionData.BID
	}

	if err := service.Preference.UpdatePreferences(prefs); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
