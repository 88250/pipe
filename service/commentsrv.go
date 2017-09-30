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
	"math"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
)

var Comment = &commentService{
	mutex: &sync.Mutex{},
}

type commentService struct {
	mutex *sync.Mutex
}

// Comment pagination arguments of admin console.
const (
	adminConsoleCommentListPageSize    = 15
	adminConsoleCommentListWindowsSize = 20
)

func (srv *commentService) ConsoleGetComments(page int, blogID uint) (ret []*model.Comment, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleCommentListPageSize
	count := 0
	db.Model(model.Comment{}).Order("id DESC").
		Where(model.Comment{BlogID: blogID}).
		Count(&count).Offset(offset).Limit(adminConsoleCommentListPageSize).Find(&ret)

	pageCount := int(math.Ceil(float64(count) / adminConsoleCommentListPageSize))
	pagination = util.NewPagination(page, adminConsoleCommentListPageSize, pageCount, adminConsoleCommentListWindowsSize, count)

	return
}

func (srv *commentService) RemoveComment(id uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment := &model.Comment{
		Model: model.Model{ID: id},
	}

	tx := db.Begin()
	if err := db.Delete(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}
