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
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
	"github.com/jinzhu/gorm"
)

// Category service.
var Category = &categoryService{
	mutex: &sync.Mutex{},
}

type categoryService struct {
	mutex *sync.Mutex
}

// Category pagination arguments of admin console.
const (
	adminConsoleCategoryListPageSize   = 15
	adminConsoleCategoryListWindowSize = 20
)

func (srv *categoryService) GetCategoryArticleCount(categoryID, blogID uint64) (ret int) {
	sql := `SELECT
	count(
		DISTINCT (b3_pipe_articles.id)
	) AS c
FROM
	b3_pipe_articles,
	b3_pipe_correlations
WHERE
	b3_pipe_articles.id = b3_pipe_correlations.id1
AND b3_pipe_correlations.type = 1
AND b3_pipe_articles.blog_id = ?
AND b3_pipe_correlations.id2 IN (
	SELECT
		id2
	FROM
		b3_pipe_correlations
	WHERE
		b3_pipe_correlations.id1 = ?
	AND b3_pipe_correlations.type = 0
	AND b3_pipe_correlations.blog_id = ?
)`
	if rows, err := db.DB().Query(sql, blogID, categoryID, blogID); nil != err {
		logger.Errorf("get category article count failed: " + err.Error())
	} else {
		rows.Next()
		if err = rows.Scan(&ret); nil != err {
			logger.Errorf("get category article count failed: " + err.Error())
		}
	}

	return
}

func (srv *categoryService) GetCategoryByPath(path string, blogID uint64) *model.Category {
	path = strings.TrimSpace(path)
	if "" == path || util.IsReservedPath(path) {
		return nil
	}
	path, _ = url.PathUnescape(path)

	ret := &model.Category{}
	if err := db.Where("`path` = ? AND `blog_ID` = ?", path, blogID).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *categoryService) UpdateCategory(category *model.Category) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	count := 0
	if db.Model(&model.Category{}).Where("`id` = ? AND `blog_id` = ?", category.ID, category.BlogID).
		Count(&count); 1 > count {
		return fmt.Errorf("not found category [id=%d] to update", category.ID)
	}

	tagStr := normalizeTagStr(category.Tags)
	category.Tags = tagStr

	if err := normalizeCategoryPath(category); nil != err {
		return err
	}

	tx := db.Begin()
	if err := tx.Model(category).Updates(category).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Where("`id1` = ? AND `type` = ? AND `blog_id` = ?",
		category.ID, model.CorrelationCategoryTag, category.BlogID).Delete(model.Correlation{}).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tagCategory(tx, category); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *categoryService) ConsoleGetCategory(id uint64) *model.Category {
	ret := &model.Category{}
	if err := db.First(ret, id).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *categoryService) AddCategory(category *model.Category) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tagStr := normalizeTagStr(category.Tags)
	category.Tags = tagStr

	if err := normalizeCategoryPath(category); nil != err {
		return err
	}

	tx := db.Begin()
	if err := tx.Create(category).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tagCategory(tx, category); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *categoryService) GetCategoriesByTag(tagTitle string, blogID uint64) (ret []*model.Category) {
	if err := db.Where("`blog_id` = ? AND `tags` LIKE ?", blogID, tagTitle).Find(&ret).Error; nil != err {
		logger.Errorf("get categories failed: " + err.Error())
	}

	return
}

func (srv *categoryService) ConsoleGetCategories(page int, blogID uint64) (ret []*model.Category, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleCategoryListPageSize
	count := 0
	if err := db.Model(&model.Category{}).Order("`number` ASC, `id` DESC").
		Where("`blog_id` = ?", blogID).
		Count(&count).Offset(offset).Limit(adminConsoleCategoryListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get categories failed: " + err.Error())
	}

	pagination = util.NewPagination(page, adminConsoleCategoryListPageSize, adminConsoleCategoryListWindowSize, count)

	return
}

func (srv *categoryService) GetCategories(size int, blogID uint64) (ret []*model.Category) {
	if err := db.Where("`blog_id` = ?", blogID).Order("`number` asc").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get categories failed: " + err.Error())
	}

	return
}

func (srv *categoryService) RemoveCategory(id, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	category := &model.Category{}

	tx := db.Begin()
	if err := tx.Where("`id` = ? AND `blog_id` = ?", id, blogID).Find(category).Error; nil != err {
		return err
	}

	if err := tx.Where("`id1` = ? AND `type` = ? AND `blog_id` = ?",
		category.ID, model.CorrelationCategoryTag, category.BlogID).Delete(model.Correlation{}).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Delete(category).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func normalizeCategoryPath(category *model.Category) error {
	path := strings.TrimSpace(category.Path)
	if "" == path {
		path = "/" + category.Title
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	count := 0
	if db.Model(&model.Category{}).Where("`path` = ? AND `id` != ? AND `blog_id` = ?", path, category.ID, category.BlogID).Count(&count); 0 < count {
		return errors.New("path is reduplicated")
	}

	category.Path = path

	return nil
}

func tagCategory(tx *gorm.DB, category *model.Category) error {
	tags := strings.Split(category.Tags, ",")
	for _, tagTitle := range tags {
		tag := &model.Tag{BlogID: category.BlogID}
		tx.Where("`title` = ? AND `blog_id` = ?", tagTitle, category.BlogID).First(tag)
		if "" == tag.Title {
			tag.Title = tagTitle
			if err := tx.Create(tag).Error; nil != err {
				return err
			}
		}

		rel := &model.Correlation{
			ID1:    category.ID,
			ID2:    tag.ID,
			Type:   model.CorrelationCategoryTag,
			BlogID: category.BlogID,
		}
		if err := tx.Create(rel).Error; nil != err {
			return err
		}
	}

	return nil
}
