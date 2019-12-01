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
	"sync"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
)

// Tag service.
var Tag = &tagService{
	mutex: &sync.Mutex{},
}

type tagService struct {
	mutex *sync.Mutex
}

const (
	adminConsoleTagListPageSize   = 15
	adminConsoleTagListWindowSize = 20
)

func (srv *tagService) ConsoleGetTags(keyword string, page int, blogID uint64) (ret []*model.Tag, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleTagListPageSize
	count := 0

	where := "`blog_id` = ?"
	whereArgs := []interface{}{blogID}
	if "" != keyword {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	if err := db.Model(&model.Tag{}).Order("`id` DESC").
		Where(where, whereArgs...).
		Count(&count).Offset(offset).Limit(adminConsoleTagListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get tags failed: " + err.Error())
	}

	pagination = util.NewPagination(page, adminConsoleTagListPageSize, adminConsoleTagListWindowSize, count)

	return
}

func (srv *tagService) GetTags(size int, blogID uint64) (ret []*model.Tag) {
	if err := db.Where("`blog_id` = ?", blogID).Order("`article_count` DESC, `id` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get tags failed: " + err.Error())
	}

	return
}

func (srv *tagService) GetTagByTitle(title string, blogID uint64) *model.Tag {
	ret := &model.Tag{}
	if err := db.Where("`title` = ? AND `blog_id` = ?", title, blogID).First(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *tagService) RemoveTag(id, blogID uint64) (err error) {
	tag := &model.Tag{}
	if err := db.Where("`id` = ? AND `blog_id` = ?", id, blogID).Find(tag).Error; nil != err {
		return err
	}

	if 0 < tag.ArticleCount {
		return errors.New("can not remove tags that have articles")
	}

	tagTitle := tag.Title
	categories := Category.GetCategoriesByTag(tagTitle, blogID)
	if 0 < len(categories) {
		return errors.New("can not remove tags in a category")
	}

	if err = db.Delete(&tag).Error; nil != err {
		logger.Errorf("delete tag [" + tagTitle + "] failed: " + err.Error())

		return
	}

	return nil
}
