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
	"errors"
	"math"
	"strings"
	"sync"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

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

func (srv *categoryService) AddCategory(category *model.Category) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tagStr := normalizeTagStr(category.Tags)
	if "" == tagStr {
		return errors.New("invalid tags [" + category.Tags + "]")
	}
	category.Tags = tagStr

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

func (srv *categoryService) ConsoleGetCategories(page int, blogID uint) (ret []*model.Category, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleCategoryListPageSize
	count := 0
	if err := db.Model(model.Category{}).Order("number ASC, id DESC").
		Where(model.Category{BlogID: blogID}).
		Count(&count).Offset(offset).Limit(adminConsoleCategoryListPageSize).Find(&ret).Error; nil != err {
		log.Errorf("get categories failed: " + err.Error())
	}

	pageCount := int(math.Ceil(float64(count) / adminConsoleCategoryListPageSize))
	pagination = util.NewPagination(page, adminConsoleCategoryListPageSize, pageCount, adminConsoleCategoryListWindowSize, count)

	return
}

func (srv *categoryService) GetCategories(size int, blogID uint) (ret []*model.Category) {
	if err := db.Where(model.Category{BlogID: blogID}).Order("number asc").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get categories failed: " + err.Error())
	}

	return
}

func (srv *categoryService) RemoveCategory(id uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	category := &model.Category{}

	tx := db.Begin()
	if err := tx.First(category, id).Error; nil != err {
		return err
	}

	if err := tx.Where("id1 = ? AND type = ?", category.ID, model.CorrelationCategoryTag).Delete(model.Correlation{}).Error; nil != err {
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

func tagCategory(tx *gorm.DB, category *model.Category) error {
	tags := strings.Split(category.Tags, ",")
	for _, tagTitle := range tags {
		tag := &model.Tag{BlogID: category.BlogID}
		tx.Where("title = ?", tagTitle).First(tag)
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
