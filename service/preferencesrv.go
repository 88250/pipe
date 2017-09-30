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
	"sync"

	"github.com/b3log/solo.go/model"
)

var Preference = &preferenceService{
	mutex: &sync.Mutex{},
}

type preferenceService struct {
	mutex *sync.Mutex
}

func (srv *preferenceService) GetPreference(preferenceName string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if nil != db.Where("name = ? AND category = ? AND blog_id = ?", preferenceName, model.SettingCategoryPreference, blogID).Find(ret).Error {
		return nil
	}

	return ret
}

func (srv *preferenceService) GetPreferences(blogID uint, preferenceNames ...string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	settings := []*model.Setting{}
	if nil != db.Where("name IN (?) AND category = ? AND blog_id = ?", preferenceNames, model.SettingCategoryPreference, blogID).Find(&settings).Error {
		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}
