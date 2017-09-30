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

func TestGetStatistic(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "0" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", setting.Value)
	}
}

func TestGetStatistics(t *testing.T) {
	settings := Statistic.GetStatistics(1, model.SettingNameStatisticArticleCount, model.SettingNameStatisticCommentCount)
	if nil == settings {
		t.Errorf("settings is nil")

		return
	}
	if 1 > len(settings) {
		t.Errorf("settings is empty")

		return
	}

	if "1" != settings[model.SettingNameStatisticCommentCount].Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", settings[model.SettingNameStatisticCommentCount].Value)
	}
}

func TestIncArticleCount(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "0" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", setting.Value)
	}

	if err := Statistic.IncArticleCount(1); nil != err {
		t.Error("Inc article count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if "1" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "2", setting.Value)
	}
}

func TestDecArticleCount(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "1" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", setting.Value)
	}

	if err := Statistic.DecArticleCount(1); nil != err {
		t.Error("Inc article count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if "0" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "2", setting.Value)
	}
}
