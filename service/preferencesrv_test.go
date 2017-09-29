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

package service

import (
	"testing"

	"github.com/b3log/solo.go/model"
)

func TestGetPreference(t *testing.T) {
	setting := Preference.GetPreference(model.SettingNamePreferenceBlogTitle, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "Solo.go 示例" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "Solo.go 示例", setting.Value)
	}
}

func TestGetPreferences(t *testing.T) {
	settings := Preference.GetPreferences(1, model.SettingNamePreferenceBlogTitle, model.SettingNamePreferenceBlogSubtitle)
	if nil == settings {
		t.Errorf("settings is nil")

		return
	}
	if 1 > len(settings) {
		t.Errorf("settings is empty")

		return
	}

	if "Solo.go 示例" != settings[model.SettingNamePreferenceBlogTitle].Value {
		t.Errorf("expected is [%s], actual is [%s]", "Solo.go 示例", settings[model.SettingNamePreferenceBlogTitle].Value)
	}
}
