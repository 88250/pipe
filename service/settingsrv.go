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
	"sync"

	"github.com/88250/pipe/cache"
	"github.com/88250/pipe/model"
)

// Setting service.
var Setting = &settingService{
	mutex: &sync.Mutex{},
}

type settingService struct {
	mutex *sync.Mutex
}

func (srv *settingService) GetSetting(category, name string, blogID uint64) *model.Setting {
	ret := cache.Setting.Get(category, name, blogID)
	if nil != ret {
		return ret
	}

	ret = &model.Setting{}
	if err := db.Where("`category` = ? AND `name` = ? AND `blog_id` = ?", category, name, blogID).Find(ret).Error; nil != err {
		return nil
	}

	cache.Setting.Put(ret)

	return ret
}

func (srv *settingService) GetCategorySettings(category string, blogID uint64) []*model.Setting {
	var ret []*model.Setting

	if err := db.Where("`category` = ? AND `blog_id` = ?", category, blogID).Find(&ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *settingService) GetAllSettings(blogID uint64) []*model.Setting {
	var ret []*model.Setting

	if err := db.Where("`category` != ? AND `blog_id` = ?", model.SettingCategoryStatistic, blogID).Find(&ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *settingService) GetSettings(category string, names []string, blogID uint64) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	var settings []*model.Setting
	if err := db.Where("`category` = ? AND `name` IN (?) AND `blog_id` = ?", category, names, blogID).Find(&settings).Error; nil != err {
		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}

func (srv *settingService) AddSetting(setting *model.Setting) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if nil != srv.GetSetting(setting.Category, setting.Name, setting.BlogID) {
		return nil
	}

	tx := db.Begin()
	if err := tx.Create(setting).Error; nil != err {
		return err
	}
	tx.Commit()

	return nil
}

func (srv *settingService) UpdateSettings(category string, settings []*model.Setting, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	for _, setting := range settings {
		if err := tx.Model(&model.Setting{}).Where("`category` = ? AND `name` = ? AND `blog_id` = ?",
			category, setting.Name, blogID).Select("value").Updates(map[string]interface{}{"value": setting.Value}).Error; nil != err {
			tx.Rollback()

			return err
		}

		cache.Setting.Put(setting)
	}
	tx.Commit()

	return nil
}
