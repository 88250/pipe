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
	ret := &model.User{GithubId: githubId}
	if err := db.Where(ret).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *userService) GetBlogAdmin(blogID uint64) *model.User {
	rel := &model.Correlation{ID1: blogID, Type: model.UserRoleBlogAdmin, BlogID: blogID}
	if err := db.Where(rel,
		blogID, model.UserRoleBlogAdmin, blogID).First(rel).Error; nil != err {
		logger.Errorf("can't get blog admin: " + err.Error())

		return nil
	}

	return srv.GetUser(rel.ID2)
}

func (srv *userService) GetPlatformAdmin() *model.User {
	rel := &model.Correlation{ID1: 1}
	if err := db.Where(rel).Order("id2 ASC").First(rel).Error; nil != err {
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
	ret := &model.User{Name: name}
	if err := db.Where(ret).First(ret).Error; nil != err {
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
	rel := &model.Correlation{ID1: blogID, Type: model.CorrelationBlogUser, BlogID: blogID}
	if err := db.Model(rel).
		Where(rel, blogID, model.CorrelationBlogUser, blogID).
		Count(&count).Offset(offset).Limit(adminConsoleUserListPageSize).Find(&correlations).Error; nil != err {
		logger.Errorf("get users failed: " + err.Error())
	}

	for _, rel := range correlations {
		user := &model.User{}
		if err := db.Where(rel.ID2).Find(user).Error; nil != err {
			logger.Errorf("get user failed: " + err.Error())

			continue
		}

		ret = append(ret, user)
	}

	pagination = util.NewPagination(page, adminConsoleUserListPageSize, adminConsoleUserListWindowSize, count)

	return
}

func (srv *userService) GetOwnBlog(userID uint64) *UserBlog {
	rel := &model.Correlation{ID2: userID, Type: model.CorrelationBlogUser, Int1: model.UserRoleBlogAdmin}
	if err := db.Where(rel,
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

func (srv *userService) GetRole(userID, blogID uint64) int {
	rel := &model.Correlation{ID1: blogID, ID2: userID, Type: model.CorrelationBlogUser}
	if err := db.Where(rel).First(rel).Error; nil != err {
		return model.UserRoleNoLogin
	}

	return rel.Int1
}

func (srv *userService) GetUserBlogs(userID uint64) (ret []*UserBlog) {
	var correlations []*model.Correlation
	if err := db.Where(&model.Correlation{ID2: userID, Type: model.CorrelationBlogUser}).
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
	rel := &model.Correlation{ID1: blogID, ID2: userID, Type: model.CorrelationBlogUser, BlogID: blogID}
	if err := db.Where(rel).
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
	if err := db.Model(&model.User{}).Order("total_article_count DESC, id DESC").Limit(size).
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
