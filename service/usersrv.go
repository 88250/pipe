// Pipe - A small and beautiful blogging platform written in golang.
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

	"github.com/b3log/pipe/model"
)

var User = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

func (srv *userService) AddUser(user *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(user).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *userService) GetUserByName(name string) *model.User {
	ret := &model.User{}
	if err := db.Where("name = ?", name).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *userService) GetUser(userID uint) *model.User {
	ret := &model.User{}
	if err := db.First(ret, userID).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *userService) GetUserByNameOrEmail(nameOrEmail string) *model.User {
	ret := &model.User{}
	if err := db.Where("name = ? OR email = ?", nameOrEmail, nameOrEmail).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

type UserBlog struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	UserID   uint   `json:"userId"`
	UserRole int    `json:"userRole"`
}

func (srv *userService) GetBlogUsers(blogID uint) (ret []*model.User) {
	correlations := []*model.Correlation{}
	if err := db.Where("id1 = ? AND type = ?", blogID, model.CorrelationBlogUser).Find(&correlations).Error; nil != err {
		return
	}

	for _, rel := range correlations {
		user := &model.User{}
		if err := db.Where("id = ?", rel.ID2).Find(user).Error; nil != err {
			return
		}

		ret = append(ret, user)
	}

	return
}

func (srv *userService) GetUserBlogs(userID uint) (ret []*UserBlog) {
	correlations := []*model.Correlation{}
	if err := db.Where("id2 = ? AND type = ?", userID, model.CorrelationBlogUser).Find(&correlations).Error; nil != err {
		return
	}

	user := srv.GetUser(userID)
	if nil == user {
		return
	}

	for _, rel := range correlations {
		blogTitleSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, rel.ID1)
		if nil == blogTitleSetting {
			continue
		}

		blogURLSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, rel.ID1)
		if nil == blogURLSetting {
			continue
		}

		blog := &UserBlog{
			ID:       rel.ID1,
			Title:    blogTitleSetting.Value,
			URL:      blogURLSetting.Value,
			UserID:   userID,
			UserRole: model.UserRoleBlogUser,
		}
		if user.BlogID == blog.ID {
			blog.UserRole = model.UserRoleBlogAdmin
		}

		ret = append(ret, blog)
	}

	return ret
}
