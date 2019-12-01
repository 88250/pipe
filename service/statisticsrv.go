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
	"github.com/jinzhu/gorm"
)

// Statistic service.
var Statistic = &statisticService{
	mutex: &sync.Mutex{},
}

type statisticService struct {
	mutex *sync.Mutex
}

func (srv *statisticService) GetAllStatistics(blogID uint64) []*model.Setting {
	var ret []*model.Setting

	if err := db.Where("`category` = ? AND `blog_id` = ?", model.SettingCategoryStatistic, blogID).Find(&ret).Error; nil != err {
		logger.Errorf("get all statistics failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *statisticService) GetStatistic(statisticName string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if err := db.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", statisticName, model.SettingCategoryStatistic, blogID).Find(ret).Error; nil != err {
		logger.Errorf("get statistic failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *statisticService) GetStatistics(blogID uint, statisticNames ...string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	var settings []*model.Setting
	if err := db.Where("`name` IN (?) AND `category` = ? AND `blog_id` = ?", statisticNames, model.SettingCategoryStatistic, blogID).Find(&settings).Error; nil != err {
		logger.Errorf("get statistics failed: " + err.Error())

		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}

func (srv *statisticService) IncViewCount(blogID uint64) error {
	tx := db.Begin()
	if err := srv.IncViewCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncViewCountWithoutTx(tx *gorm.DB, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", model.SettingNameStatisticViewCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Model(setting).Updates(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) IncArticleCount(blogID uint64) error {
	tx := db.Begin()
	if err := srv.IncArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncArticleCountWithoutTx(tx *gorm.DB, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Model(setting).Updates(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecArticleCount(blogID uint64) error {
	tx := db.Begin()
	if err := srv.DecArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecArticleCountWithoutTx(tx *gorm.DB, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count - 1)
	if err := tx.Model(setting).Updates(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) IncCommentCount(blogID uint64) error {
	tx := db.Begin()
	if err := srv.IncCommentCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncCommentCountWithoutTx(tx *gorm.DB, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", model.SettingNameStatisticCommentCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Model(setting).Updates(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecCommentCount(blogID uint64) error {
	tx := db.Begin()
	if err := srv.DecCommentCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecCommentCountWithoutTx(tx *gorm.DB, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("`name` = ? AND `category` = ? AND `blog_id` = ?", model.SettingNameStatisticCommentCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}
	setting.Value = strconv.Itoa(count - 1)
	if err := tx.Model(setting).Updates(setting).Error; nil != err {
		return err
	}

	return nil
}
