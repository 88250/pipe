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

package service

import (
	"testing"

	"github.com/88250/pipe/model"
)

func TestGetSetting(t *testing.T) {
	setting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "pipe 的博客" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "pipe 的博客", setting.Value)
	}
}

func TestGetAllSettings(t *testing.T) {
	settings := Setting.GetAllSettings(1)
	settingsCount := 27
	if settingsCount != len(settings) {
		t.Errorf("expected is [%d], actual is [%d]", settingsCount, len(settings))
	}
}

func TestGetCategorySettings(t *testing.T) {
	basicSettings := Setting.GetCategorySettings(model.SettingCategoryBasic, 1)
	if 11 != len(basicSettings) {
		t.Errorf("expected is [%d], actual is [%d]", 10, len(basicSettings))
	}
}

func TestGetSettings(t *testing.T) {
	settings := Setting.GetSettings(model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle}, 1)
	if nil == settings {
		t.Errorf("settings is nil")

		return
	}
	if 1 > len(settings) {
		t.Errorf("settings is empty")

		return
	}

	if "pipe 的博客" != settings[model.SettingNameBasicBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "pipe 的博客", settings[model.SettingNameBasicBlogTitle].Value)
	}
}

func TestUpdateSettings(t *testing.T) {
	settings := Setting.GetSettings(model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle}, 1)
	settings[model.SettingNameBasicBlogTitle].Value = "更新后的标题"
	var basics []*model.Setting
	for _, setting := range settings {
		basics = append(basics, setting)
	}
	if err := Setting.UpdateSettings(model.SettingCategoryBasic, basics, 1); nil != err {
		t.Errorf("updates settings failed: " + err.Error())

		return
	}

	settings = Setting.GetSettings(model.SettingCategoryBasic,
		[]string{model.SettingNameBasicBlogTitle, model.SettingNameBasicBlogSubtitle}, 1)
	if "更新后的标题" != settings[model.SettingNameBasicBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "更新后的标题", settings[model.SettingNameBasicBlogTitle].Value)
	}
}
