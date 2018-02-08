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

package service

import (
	"sync"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
)

var Upgrade = &upgradeService{
	mutex: &sync.Mutex{},
}

type upgradeService struct {
	mutex *sync.Mutex
}

const (
	fromVer = "1.2.0"
	toVer   = util.Version
)

func (srv *upgradeService) Perform() {
	if !Init.Inited() {
		return
	}
	sysVerSetting := Setting.GetSetting(model.SettingCategorySystem, model.SettingNameSystemVer, 1)
	if nil == sysVerSetting {
		logger.Fatalf("system state is error, please contact developer: https://github.com/b3log/pipe/issues/new")
	}

	currentVer := sysVerSetting.Value
	if util.Version == currentVer {
		return
	}

	if fromVer == currentVer {
		perform()

		return
	}

	logger.Fatalf("attempt to skip more than one version to upgrade. Expected: %s, Actually: %s", fromVer, currentVer)
}

func perform() {
	logger.Infof("upgrading from version [%s] to version [%s]....", fromVer, toVer)

	var allSettings []model.Setting
	if err := db.Find(&allSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	var updateSettings []model.Setting
	for _, setting := range allSettings {
		if model.SettingNameSystemVer == setting.Name {
			setting.Value = util.Version
			updateSettings = append(updateSettings, setting)

			continue
		}
	}

	tx := db.Begin()
	for _, setting := range updateSettings {
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err)
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [%s] to version [%s] successfully :-)", fromVer, toVer)
}
