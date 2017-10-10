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
	"strconv"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/jinzhu/gorm"
)

var Statistic = &statisticService{
	mutex: &sync.Mutex{},
}

type statisticService struct {
	mutex *sync.Mutex
}

func (srv *statisticService) GetStatistic(statisticName string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if nil != db.Where("name = ? AND category = ? AND blog_id = ?", statisticName, model.SettingCategoryStatistic, blogID).Find(ret).Error {
		return nil
	}

	return ret
}

func (srv *statisticService) GetStatistics(blogID uint, statisticNames ...string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	settings := []*model.Setting{}
	if nil != db.Where("name IN (?) AND category = ? AND blog_id = ?", statisticNames, model.SettingCategoryStatistic, blogID).Find(&settings).Error {
		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}

func (srv *statisticService) IncArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.IncArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncArticleCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.DecArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecArticleCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count - 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) IncPublishedArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.IncPublishedArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncPublishedArticleCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticPublishedArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecPublishedArticleCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.DecPublishedArticleCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecPublishedArticleCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticPublishedArticleCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count - 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) IncCommentCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.IncCommentCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) IncCommentCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticCommentCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}

	setting.Value = strconv.Itoa(count + 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}

func (srv *statisticService) DecCommentCount(blogID uint) error {
	tx := db.Begin()
	if err := srv.DecCommentCountWithoutTx(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *statisticService) DecCommentCountWithoutTx(tx *gorm.DB, blogID uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	setting := &model.Setting{}
	if err := tx.Where("name = ? AND category = ? AND blog_id = ?", model.SettingNameStatisticCommentCount, model.SettingCategoryStatistic, blogID).Find(setting).Error; nil != err {
		return err
	}

	count, err := strconv.Atoi(setting.Value)
	if nil != err {
		return err
	}
	setting.Value = strconv.Itoa(count - 1)
	if err := tx.Save(setting).Error; nil != err {
		return err
	}

	return nil
}
