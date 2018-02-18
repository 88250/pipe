// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func GetBasicSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	settings := service.Setting.GetCategorySettings(model.SettingCategoryBasic, session.BID)
	data := map[string]interface{}{}
	for _, setting := range settings {
		if model.SettingNameBasicCommentable == setting.Name {
			v, err := strconv.ParseBool(setting.Value)
			if nil != err {
				logger.Errorf("value of basic setting [name=%s] must be \"true\" or \"false\"", setting.Name)
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

func UpdateBasicSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update basic settings request failed"

		return
	}

	session := util.GetSession(c)
	var basics []*model.Setting
	for k, v := range args {
		var value interface{}
		switch v.(type) {
		case bool:
			value = strconv.FormatBool(v.(bool))
		default:
			value = strings.TrimSpace(v.(string))
		}

		if model.SettingNameBasicBlogURL == k {
			blogURL := value.(string)
			if !strings.Contains(blogURL, "://") {
				blogURL = "http://" + blogURL
			}

			url, err := url.Parse(blogURL)
			if nil != err {
				result.Code = -1
				result.Msg = "invalid URL format"

				return
			}

			blogURL = url.Scheme + "://" + url.Host
			if "" != url.Path {
				blogURL +=  path.Clean(url.Path)
			}
			value = blogURL
		}

		basic := &model.Setting{
			Category: model.SettingCategoryBasic,
			BlogID:   session.BID,
			Name:     k,
			Value:    value.(string),
		}
		basics = append(basics, basic)
	}

	if err := service.Setting.UpdateSettings(model.SettingCategoryBasic, basics, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetPreferenceSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	settings := service.Setting.GetCategorySettings(model.SettingCategoryPreference, session.BID)
	data := map[string]interface{}{}
	for _, setting := range settings {
		if model.SettingNamePreferenceArticleListStyle != setting.Name {
			v, err := strconv.ParseInt(setting.Value, 10, 64)
			if nil != err {
				logger.Errorf("value of preference setting [name=%s] must be an integer", setting.Name)
				data[setting.Name] = 10
			} else {
				data[setting.Name] = v
			}
		} else {
			data[setting.Name] = setting.Value
		}
	}

	result.Data = data
}

func UpdatePreferenceSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update preference settings request failed"

		return
	}

	session := util.GetSession(c)
	var prefs []*model.Setting
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
			BlogID:   session.BID,
			Name:     k,
			Value:    value.(string),
		}
		prefs = append(prefs, pref)
	}

	if err := service.Setting.UpdateSettings(model.SettingCategoryPreference, prefs, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetSignSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	signSetting := service.Setting.GetSetting(model.SettingCategorySign, model.SettingNameArticleSign, session.BID)
	result.Data = signSetting.Value
}

func UpdateSignSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update sign settings request failed"

		return
	}

	session := util.GetSession(c)
	var signs []*model.Setting
	sign := &model.Setting{
		Category: model.SettingCategorySign,
		BlogID:   session.BID,
		Name:     model.SettingNameArticleSign,
		Value:    args["sign"].(string),
	}
	signs = append(signs, sign)

	if err := service.Setting.UpdateSettings(model.SettingCategorySign, signs, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetI18nSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	settings := service.Setting.GetCategorySettings(model.SettingCategoryI18n, session.BID)
	data := map[string]interface{}{}
	for _, setting := range settings {
		data[setting.Name] = setting.Value
	}
	result.Data = data
}

func UpdateI18nSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update i18n settings request failed"

		return
	}

	session := util.GetSession(c)
	var i18ns []*model.Setting
	for k, v := range args {
		i18n := &model.Setting{
			Category: model.SettingCategoryI18n,
			BlogID:   session.BID,
			Name:     k,
			Value:    v.(string),
		}
		i18ns = append(i18ns, i18n)
	}

	if err := service.Setting.UpdateSettings(model.SettingCategoryI18n, i18ns, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetFeedSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	settings := service.Setting.GetCategorySettings(model.SettingCategoryFeed, session.BID)
	data := map[string]interface{}{}
	for _, setting := range settings {
		if model.SettingNameFeedOutputMode == setting.Name {
			v, err := strconv.ParseInt(setting.Value, 10, 64)
			if nil != err {
				logger.Errorf("value of feed setting [name=%s] must be an integer", setting.Name)
				data[setting.Name] = 20
			} else {
				data[setting.Name] = v
			}
		} else {
			data[setting.Name] = setting.Value
		}
	}
	result.Data = data
}

func UpdateFeedSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update feed settings request failed"

		return
	}

	session := util.GetSession(c)
	var feeds []*model.Setting
	for k, v := range args {
		var value interface{}
		switch v.(type) {
		case float64:
			value = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		default:
			value = v.(string)
		}

		feed := &model.Setting{
			Category: model.SettingCategoryFeed,
			BlogID:   session.BID,
			Name:     k,
			Value:    value.(string),
		}
		feeds = append(feeds, feed)
	}

	if err := service.Setting.UpdateSettings(model.SettingCategoryFeed, feeds, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetThirdStatisticSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	baiduStatisticSetting := service.Setting.GetSetting(model.SettingCategoryThirdStatistic, model.SettingNameThirdStatisticBaidu, session.BID)
	data := map[string]string{
		model.SettingNameThirdStatisticBaidu: baiduStatisticSetting.Value,
	}
	result.Data = data
}

func UpdateThirdStatisticSettingsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Code = -1
		result.Msg = "parses update third statistic settings request failed"

		return
	}

	session := util.GetSession(c)
	var thridStatistics []*model.Setting
	baiduStatistic := &model.Setting{
		Category: model.SettingCategoryThirdStatistic,
		BlogID:   session.BID,
		Name:     model.SettingNameThirdStatisticBaidu,
		Value:    args["thirdStatisticBaidu"].(string),
	}
	thridStatistics = append(thridStatistics, baiduStatistic)

	if err := service.Setting.UpdateSettings(model.SettingCategoryThirdStatistic, thridStatistics, session.BID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
