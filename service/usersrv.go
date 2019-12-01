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
	"github.com/88250/pipe/util"
)

// User service.
var User = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

// User pagination arguments of admin console.
const (
	adminConsoleUserListPageSize   = 15
	adminConsoleUserListWindowSize = 20
)

func (srv *userService) GetUserByGitHubId(githubId string) *model.User {
	ret := &model.User{}
	if err := db.Where("`github_id` = ?", githubId).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *userService) GetBlogAdmin(blogID uint64) *model.User {
	rel := &model.Correlation{}
	if err := db.Where("`id1` = ? AND `type` = ? AND `blog_id` = ?",
		blogID, model.UserRoleBlogAdmin, blogID).First(rel).Error; nil != err {
		logger.Errorf("can't get blog admin: " + err.Error())

		return nil
	}

	return srv.GetUser(rel.ID2)
}

func (srv *userService) GetPlatformAdmin() *model.User {
	rel := &model.Correlation{}
	if err := db.Where("`id1` = ?", 1).Order("`id2` asc").First(rel).Error; nil != err {
		logger.Errorf("can't get platform admin: " + err.Error())

		return nil
	}

	return srv.GetUser(rel.ID2)
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

func (srv *userService) UpdateUser(user *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(user).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	cache.User.Put(user)

	return nil
}

func (srv *userService) GetUserByName(name string) *model.User {
	ret := &model.User{}
	if err := db.Where("`name` = ?", name).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *userService) GetUser(userID uint64) *model.User {
	ret := cache.User.Get(userID)
	if nil != ret {
		return ret
	}

	ret = &model.User{}
	if err := db.First(ret, userID).Error; nil != err {
		return nil
	}

	cache.User.Put(ret)

	return ret
}

// UserBlog represents user blog.
type UserBlog struct {
	ID               uint64 `json:"id,omitempty"` // blog ID
	Title            string `json:"title"`        // blog title
	URL              string `json:"url"`          // blog URL
	UserID           uint64 `json:"userId,omitempty"`
	UserRole         int    `json:"userRole,omitempty"`
	UserArticleCount int    `json:"userArticleCount"`
}

func (srv *userService) GetBlogUsers(page int, blogID uint64) (ret []*model.User, pagination *util.Pagination) {
	var correlations []*model.Correlation
	offset := (page - 1) * adminConsoleUserListPageSize
	count := 0
	if err := db.Model(&model.Correlation{}).
		Where("`id1` = ? AND `type` = ? AND `blog_id` = ?", blogID, model.CorrelationBlogUser, blogID).
		Count(&count).Offset(offset).Limit(adminConsoleUserListPageSize).Find(&correlations).Error; nil != err {
		logger.Errorf("get users failed: " + err.Error())
	}

	for _, rel := range correlations {
		user := &model.User{}
		if err := db.Where("`id` = ?", rel.ID2).Find(user).Error; nil != err {
			logger.Errorf("get user failed: " + err.Error())

			continue
		}

		ret = append(ret, user)
	}

	pagination = util.NewPagination(page, adminConsoleUserListPageSize, adminConsoleUserListWindowSize, count)

	return
}

func (srv *userService) GetOwnBlog(userID uint64) *UserBlog {
	rel := &model.Correlation{}
	if err := db.Where("`id2` = ? AND `type` = ? AND `int1` = ?",
		userID, model.CorrelationBlogUser, model.UserRoleBlogAdmin).First(rel).Error; nil != err {
		return nil
	}

	blogTitleSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, rel.ID1)
	blogURLSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, rel.ID1)

	return &UserBlog{
		ID:               rel.ID1,
		Title:            blogTitleSetting.Value,
		URL:              blogURLSetting.Value,
		UserID:           userID,
		UserRole:         rel.Int1,
		UserArticleCount: rel.Int2,
	}
}

func (srv *userService) GetRole(userID, blogID uint) int {
	rel := &model.Correlation{}
	if err := db.Where("`id1` = ? AND `id2` = ? AND `type` = ?",
		blogID, userID, model.CorrelationBlogUser).First(rel).Error; nil != err {
		return model.UserRoleNoLogin
	}

	return rel.Int1
}

func (srv *userService) GetUserBlogs(userID uint64) (ret []*UserBlog) {
	var correlations []*model.Correlation
	if err := db.Where("`id2` = ? AND `type` = ?", userID, model.CorrelationBlogUser).
		Find(&correlations).Error; nil != err {
		return
	}

	for _, rel := range correlations {
		blogTitleSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, rel.ID1)
		blogURLSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, rel.ID1)

		userBlog := &UserBlog{
			ID:               rel.ID1,
			Title:            blogTitleSetting.Value,
			URL:              blogURLSetting.Value,
			UserID:           userID,
			UserRole:         rel.Int1,
			UserArticleCount: rel.Int2,
		}
		ret = append(ret, userBlog)
	}

	return ret
}

func (srv *userService) GetUserBlog(userID, blogID uint64) *UserBlog {
	rel := &model.Correlation{}
	if err := db.Where("`id1` = ? AND `id2` = ? AND `type` = ? AND `blog_id` = ?", blogID, userID, model.CorrelationBlogUser, blogID).
		First(&rel).Error; nil != err {
		return nil
	}

	blogTitleSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, rel.ID1)
	blogURLSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, rel.ID1)

	return &UserBlog{
		ID:               rel.ID1,
		Title:            blogTitleSetting.Value,
		URL:              blogURLSetting.Value,
		UserID:           userID,
		UserRole:         rel.Int1,
		UserArticleCount: rel.Int2,
	}
}

func (srv *userService) AddUserToBlog(userID, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if nil != srv.GetUserBlog(userID, blogID) {
		return nil
	}

	blogUser := &model.Correlation{
		ID1:    blogID,
		ID2:    userID,
		Type:   model.CorrelationBlogUser,
		Int1:   model.UserRoleBlogUser,
		Int2:   0,
		BlogID: blogID,
	}
	if err := db.Create(blogUser).Error; nil != err {
		return err
	}

	return nil
}

func (srv *userService) GetTopBlogs(size int) (ret []*UserBlog) {
	var users []*model.User
	if err := db.Model(&model.User{}).Order("`total_article_count` DESC, `id` DESC").Limit(size).
		Find(&users).Error; nil != err {
		return
	}

	for _, user := range users {
		userBlog := srv.GetOwnBlog(user.ID)
		if nil != userBlog && 5 <= userBlog.UserArticleCount {
			ret = append(ret, userBlog)
		}
	}

	if 1 > len(ret) && 1 <= len(users) {
		userBlog := srv.GetOwnBlog(users[0].ID)
		ret = append(ret, userBlog)
	}

	return ret
}
