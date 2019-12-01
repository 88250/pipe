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

func TestGetAllStatistics(t *testing.T) {
	settings := Statistic.GetAllStatistics(1)
	if 3 != len(settings) {
		t.Errorf("expected is [%d], actual is [%d]", 3, len(settings))
	}
}

func TestGetStatistic(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "99" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "99", setting.Value)
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

	if "99" != settings[model.SettingNameStatisticArticleCount].Value {
		t.Errorf("expected is [%s], actual is [%s]", "99", settings[model.SettingNameStatisticArticleCount].Value)
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

	if "99" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "99", setting.Value)
	}

	if err := Statistic.IncArticleCount(1); nil != err {
		t.Error("Inc article count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if "100" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "100", setting.Value)
	}
}

func TestDecArticleCount(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "100" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "100", setting.Value)
	}

	if err := Statistic.DecArticleCount(1); nil != err {
		t.Error("dec article count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if "99" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "99", setting.Value)
	}
}

func TestIncCommentCount(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticCommentCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "1" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", setting.Value)
	}

	if err := Statistic.IncCommentCount(1); nil != err {
		t.Error("inc article count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticCommentCount, 1)
	if "2" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "2", setting.Value)
	}
}

func TestDecCommentCount(t *testing.T) {
	setting := Statistic.GetStatistic(model.SettingNameStatisticCommentCount, 1)
	if nil == setting {
		t.Errorf("setting is nil")

		return
	}

	if "2" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "2", setting.Value)
	}

	if err := Statistic.DecCommentCount(1); nil != err {
		t.Error("dec comment count failed")

		return
	}

	setting = Statistic.GetStatistic(model.SettingNameStatisticCommentCount, 1)
	if "1" != setting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "1", setting.Value)
	}
}
