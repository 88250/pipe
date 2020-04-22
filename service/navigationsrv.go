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
	"fmt"
	"sync"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
)

// Navigation service.
var Navigation = &navigationService{
	mutex: &sync.Mutex{},
}

type navigationService struct {
	mutex *sync.Mutex
}

// Navigation pagination arguments of admin console.
const (
	adminConsoleNavigationListPageSize   = 15
	adminConsoleNavigationListWindowSize = 20
)

func (srv *navigationService) AddNavigation(navigation *model.Navigation) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(navigation).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *navigationService) RemoveNavigation(id, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	navigation := &model.Navigation{}

	tx := db.Begin()
	if err := tx.Where("id = ? AND blog_id = ?", id, blogID).Find(navigation).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Delete(navigation).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *navigationService) UpdateNavigation(navigation *model.Navigation) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	count := 0
	if db.Model(&model.Navigation{}).Where("id = ? AND blog_id = ?", navigation.ID, navigation.BlogID).
		Count(&count); 1 > count {
		return fmt.Errorf("not found navigation [id=%d] to update", navigation.ID)
	}

	tx := db.Begin()
	if err := tx.Model(navigation).Updates(map[string]interface{}{
		"Title":      navigation.Title,
		"URL":        navigation.URL,
		"IconURL":    navigation.IconURL,
		"OpenMethod": navigation.OpenMethod}).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *navigationService) ConsoleGetNavigations(page int, blogID uint64) (ret []*model.Navigation, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleNavigationListPageSize
	count := 0
	if err := db.Model(&model.Navigation{}).Order("number ASC, id DESC").
		Where("blog_id = ?", blogID).
		Count(&count).Offset(offset).Limit(adminConsoleNavigationListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get navigations failed: " + err.Error())
	}

	pagination = util.NewPagination(page, adminConsoleNavigationListPageSize, adminConsoleNavigationListWindowSize, count)

	return
}

func (srv *navigationService) GetNavigations(blogID uint64) (ret []*model.Navigation) {
	if err := db.Model(&model.Navigation{}).Order("number ASC, id DESC").
		Where("blog_id = ?", blogID).Find(&ret).Error; nil != err {
		logger.Errorf("get navigations failed: " + err.Error())
	}

	return
}

func (srv *navigationService) ConsoleGetNavigation(id uint64) *model.Navigation {
	ret := &model.Navigation{}
	if err := db.First(ret, id).Error; nil != err {
		return nil
	}

	return ret
}
