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
	"strconv"
	"sync"

	"github.com/88250/pipe/model"
)

// Upgrade service.
var Upgrade = &upgradeService{
	mutex: &sync.Mutex{},
}

type upgradeService struct {
	mutex *sync.Mutex
}

func (srv *upgradeService) Perform() {
	if !Init.Inited() {
		return
	}
	sysVerSetting := Setting.GetSetting(model.SettingCategorySystem, model.SettingNameSystemVer, 1)
	if nil == sysVerSetting {
		logger.Fatalf("system state is error, please contact developer: https://github.com/88250/pipe/issues/new")
	}

	currentVer := sysVerSetting.Value
	if model.Version == currentVer {
		return
	}

	switch currentVer {
	case "1.8.6":
		perform186_187()
		fallthrough
	case "1.8.7":
		perform187_188()
		fallthrough
	case "1.8.8":
		perform188_189()
		fallthrough
	case "1.8.9":
		perform189_190()
	default:
		logger.Fatalf("please upgrade to v1.8.7 first")
	}
}

func perform189_190() {
	fromVer := "1.8.9"
	logger.Infof("upgrading from version [" + fromVer + "] to version [" + model.Version + "]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("`name`= ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = model.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [" + fromVer + "] to version [" + model.Version + "] successfully")
}

func perform188_189() {
	logger.Infof("upgrading from version [1.8.8] to version [1.8.9]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("`name`= ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = model.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [1.8.8] to version [1.8.9] successfully")
}

func perform187_188() {
	logger.Infof("upgrading from version [1.8.7] to version [1.8.8]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("`name`= ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = model.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}

	rows, err := tx.Model(&model.Setting{}).Select("`blog_id`").Group("`blog_id`").Rows()
	if nil != err {
		tx.Rollback()

		logger.Fatalf("update settings failed: %s", err.Error())
	}

	blogIDs := []uint64{}
	for rows.Next() {
		var blogID uint64
		err := rows.Scan(&blogID)
		if nil != err {
			tx.Rollback()

			logger.Fatalf("update settings failed: %s", err.Error())
		}
		blogIDs = append(blogIDs, blogID)
	}
	rows.Close()

	for _, blogID := range blogIDs {
		if err := tx.Create(&model.Setting{
			Category: model.SettingCategoryPreference,
			Name:     model.SettingNamePreferenceRecommendArticleListSize,
			Value:    strconv.Itoa(model.SettingPreferenceRecommendArticleListSizeDefault),
			BlogID:   blogID}).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update settings failed: %s", err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [1.8.7] to version [1.8.8] successfully")
}

func perform186_187() {
	logger.Infof("upgrading from version [1.8.6] to version [1.8.7]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("`name`= ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = model.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [1.8.6] to version [1.8.7] successfully")
}
