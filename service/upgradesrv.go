// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package service

import (
	"strconv"
	"sync"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
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
	if util.Version == currentVer {
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
		fallthrough
	case "1.9.0":
		perform190_191()
		fallthrough
	case "1.9.1":
		perform191_200()
		fallthrough
	case "2.0.0":
		perform200_210()
	default:
		logger.Fatalf("please upgrade to v1.8.7 first")
	}
}

func perform200_210() {
	fromVer := "2.0.0"
	logger.Infof("upgrading from version [" + fromVer + "] to version [" + util.Version + "]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [" + fromVer + "] to version [" + util.Version + "] successfully")
}

func perform191_200() {
	fromVer := "1.9.1"
	logger.Infof("upgrading from version [" + fromVer + "] to version [" + util.Version + "]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [" + fromVer + "] to version [" + util.Version + "] successfully")
}

func perform190_191() {
	fromVer := "1.9.0"
	logger.Infof("upgrading from version [" + fromVer + "] to version [" + util.Version + "]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [" + fromVer + "] to version [" + util.Version + "] successfully")
}

func perform189_190() {
	fromVer := "1.8.9"
	logger.Infof("upgrading from version [" + fromVer + "] to version [" + util.Version + "]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [" + fromVer + "] to version [" + util.Version + "] successfully")
}

func perform188_189() {
	logger.Infof("upgrading from version [1.8.8] to version [1.8.9]....")

	var verSettings []model.Setting
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
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
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}

	rows, err := tx.Model(&model.Setting{}).Select(" blog_id ").Group(" blog_id ").Rows()
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
	if err := db.Model(&model.Setting{}).Where("name = ?", model.SettingNameSystemVer).Find(&verSettings).Error; nil != err {
		logger.Fatalf("load settings failed: %s", err)
	}

	tx := db.Begin()
	for _, setting := range verSettings {
		setting.Value = util.Version
		if err := tx.Save(setting).Error; nil != err {
			tx.Rollback()

			logger.Fatalf("update setting [%+v] failed: %s", setting, err.Error())
		}
	}
	tx.Commit()

	logger.Infof("upgraded from version [1.8.6] to version [1.8.7] successfully")
}
