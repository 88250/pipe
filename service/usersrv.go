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
	log "github.com/sirupsen/logrus"
)

var User = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

func (srv *userService) GetUser(userID uint) *model.User {
	ret := &model.User{}
	if nil != db.First(ret, userID).Error {
		return nil
	}

	return ret
}

func (srv *userService) GetUserByNameOrEmail(nameOrEmail string) *model.User {
	ret := &model.User{}
	if nil != db.Where("name = ? OR email = ?", nameOrEmail, nameOrEmail).Find(ret).Error {
		return nil
	}

	return ret
}

type Blog struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (srv *userService) GetUserBlogs(userID uint) []*Blog {
	correlations := []*model.Correlation{}
	if nil != db.Where("id2 = ? AND type = ?", userID, model.CorrelationBlogUser).Find(&correlations).Error {
		return nil
	}

	ret := []*Blog{}
	for _, rel := range correlations {
		prefs := Preference.GetPreferences(rel.ID1, model.SettingNamePreferenceBlogTitle, model.SettingNamePreferencePath)
		if nil == prefs {
			log.Errorf("not found blog setting [blogID=%d]", rel.ID1)

			continue
		}

		blog := &Blog{
			ID:    rel.ID1,
			Title: prefs[model.SettingNamePreferenceBlogTitle].Value,
			Path:  prefs[model.SettingNamePreferencePath].Value,
		}

		ret = append(ret, blog)
	}

	return ret
}
