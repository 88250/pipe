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
	"math"
	"sync"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	log "github.com/sirupsen/logrus"
)

var Comment = &commentService{
	mutex: &sync.Mutex{},
}

type commentService struct {
	mutex *sync.Mutex
}

// Comment pagination arguments of admin console.
const (
	adminConsoleCommentListPageSize   = 15
	adminConsoleCommentListWindowSize = 20
)

// Comment pagination arguments of theme.
const (
	themeCommentListPageSize   = 15
	themeCommentListWindowSize = 20
)

func (srv *commentService) GetComment(commentID uint) *model.Comment {
	ret := &model.Comment{}
	if err := db.First(ret, commentID).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *commentService) GetCommentPage(articleID, commentID uint, blogID uint) int {
	count := 0
	if err := db.Model(&model.Comment{}).Where("article_id = ? AND id < ? AND blog_id = ?", articleID, commentID, blogID).
		Count(&count).Error; nil != err {
		return 1
	}

	return (count / adminConsoleCommentListPageSize) + 1
}

func (srv *commentService) GetRepliesCount(parentCommentID uint, blogID uint) int {
	ret := 0
	if err := db.Model(&model.Comment{}).Where(&model.Comment{ParentCommentID: parentCommentID, BlogID: blogID}).Count(&ret).Error; nil != err {
		log.Errorf("count comment [id=%d]'s replies failed: "+err.Error(), parentCommentID)
	}

	return ret
}

func (srv *commentService) GetReplies(parentCommentID uint, blogID uint) (ret []*model.Comment) {
	if err := db.Where(&model.Comment{ParentCommentID: parentCommentID, BlogID: blogID}).Find(&ret).Error; nil != err {
		log.Errorf("get comment [id=%d]'s replies failed: "+err.Error(), parentCommentID)
	}

	return
}

func (srv *commentService) ConsoleGetComments(page int, blogID uint) (ret []*model.Comment, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleCommentListPageSize
	count := 0
	if err := db.Model(model.Comment{}).Order("id ASC").
		Where(&model.Comment{BlogID: blogID}).
		Count(&count).Offset(offset).Limit(adminConsoleCommentListPageSize).Find(&ret).Error; nil != err {
		log.Errorf("get comments failed: " + err.Error())
	}

	pageCount := int(math.Ceil(float64(count) / adminConsoleCommentListPageSize))
	pagination = util.NewPagination(page, adminConsoleCommentListPageSize, pageCount, adminConsoleCommentListWindowSize, count)

	return
}

func (srv *commentService) GetRecentComments(size int, blogID uint) (ret []*model.Comment) {
	if err := db.Model(model.Comment{}).Select("id, created_at, content").
		Where(model.Comment{BlogID: blogID}).
		Order("created_at DESC, id DESC").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get recent comments failed: " + err.Error())
	}

	return
}

func (srv *commentService) GetArticleComments(articleID uint, page int, blogID uint) (ret []*model.Comment, pagination *util.Pagination) {
	offset := (page - 1) * themeCommentListPageSize
	count := 0
	if err := db.Model(model.Comment{}).Order("id ASC").
		Where(model.Comment{ArticleID: articleID, BlogID: blogID}).
		Count(&count).Offset(offset).Limit(themeCommentListPageSize).Find(&ret).Error; nil != err {
		log.Errorf("get comments failed: " + err.Error())
	}

	pageCount := int(math.Ceil(float64(count) / themeCommentListPageSize))
	pagination = util.NewPagination(page, themeCommentListPageSize, pageCount, themeCommentListWindowSize, count)

	return
}

func (srv *commentService) AddComment(comment *model.Comment) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment.ID = util.CurrentMillisecond()

	tx := db.Begin()
	if err := tx.Create(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	Statistic.IncCommentCountWithoutTx(tx, comment.BlogID)
	tx.Commit()

	return nil
}

func (srv *commentService) RemoveComment(id uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment := &model.Comment{}

	tx := db.Begin()
	if err := tx.First(comment, id).Error; nil != err {
		return err
	}
	if err := tx.Delete(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	Statistic.DecCommentCountWithoutTx(tx, comment.BlogID)
	tx.Commit()

	return nil
}
