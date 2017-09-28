// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
	log "github.com/sirupsen/logrus"
)

var Article = &articleService{
	mutex: &sync.Mutex{},
}

type articleService struct {
	mutex *sync.Mutex
}

// Article pagination arguments of admin console.
const (
	adminConsoleArticleListPageSize    = 15
	adminConsoleArticleListWindowsSize = 20
)

func (srv *articleService) AddArticle(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()

	if err := tx.Create(article).Error; nil != err {
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (srv *articleService) ConsoleGetArticles(page int) (ret []*model.Article, pagination *util.Pagination) {
	if 1 > page {
		page = 1
	}

	offset := (page - 1) * adminConsoleArticleListPageSize
	count := 0
	db.Model(model.Article{}).Select("id, created_at, title, tags, topped, view_count, comment_count").Where(model.Article{Status: model.ArticleStatusPublished}).
		Order("topped DESC, id DESC").Count(&count).
		Offset(offset).Limit(adminConsoleArticleListPageSize).
		Find(&ret)

	pageCount := int(math.Ceil(float64(count) / adminConsoleArticleListPageSize))
	pagination = util.NewPagination(page, adminConsoleArticleListPageSize, pageCount, adminConsoleArticleListWindowsSize, count)

	return
}

func (srv *articleService) ConsoleGetArticle(id uint) *model.Article {
	ret := &model.Article{}
	if nil != db.First(ret, id).Error {
		return nil
	}

	return ret
}

func (srv *articleService) RemoveArticle(id uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	article := &model.Article{
		Model: model.Model{ID: id},
	}

	return db.Delete(article).Error
}

func (srv *articleService) UpdateArticle(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	count := 0
	if db.Model(model.Article{}).First(article).Count(&count); 1 > count {
		return errors.New(fmt.Sprintf("not found article [id=%d] to update", article.ID))
	}

	return db.Model(&model.Article{}).Updates(article).Error
}
