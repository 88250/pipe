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

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
)

// Comment service.
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

func (srv *commentService) UpdatePushedAt(comment *model.Comment) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment.PushedAt = comment.UpdatedAt
	if err := db.Model(comment).UpdateColumns(comment).Error; nil != err {
		return err
	}

	return nil
}

func (srv *commentService) GetUnpushedComments() (ret []*model.Comment) {
	if err := db.Where("`pushed_at` <= ?", model.ZeroPushTime).Find(&ret).Error; nil != err {
		return
	}

	return
}

func (srv *commentService) GetComment(commentID uint64) *model.Comment {
	ret := &model.Comment{}
	if err := db.First(ret, commentID).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *commentService) GetCommentPage(articleID, commentID uint64, blogID uint64) int {
	count := 0
	if err := db.Model(&model.Comment{}).Where("`article_id` = ? AND `id` < ? AND `blog_id` = ?", articleID, commentID, blogID).
		Count(&count).Error; nil != err {
		return 1
	}

	return (count / adminConsoleCommentListPageSize) + 1
}

func (srv *commentService) GetRepliesCount(parentCommentID uint64, blogID uint64) int {
	ret := 0
	if err := db.Model(&model.Comment{}).Where("`parent_comment_id` = ? AND `blog_id` = ?", parentCommentID, blogID).Count(&ret).Error; nil != err {
		logger.Errorf("count comment [id=%d]'s replies failed: "+err.Error(), parentCommentID)
	}

	return ret
}

func (srv *commentService) GetReplies(parentCommentID uint64, blogID uint64) (ret []*model.Comment) {
	if err := db.Where("`parent_comment_id` = ? AND `blog_id` = ?", parentCommentID, blogID).Find(&ret).Error; nil != err {
		logger.Errorf("get comment [id=%d]'s replies failed: "+err.Error(), parentCommentID)
	}

	return
}

func (srv *commentService) ConsoleGetComments(keyword string, page int, blogID uint64) (ret []*model.Comment, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleCommentListPageSize
	count := 0

	where := "`blog_id` = ?"
	whereArgs := []interface{}{blogID}
	if "" != keyword {
		where += " AND `content` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	if err := db.Model(&model.Comment{}).
		Where(where, whereArgs...).Order("`created_at` DESC").
		Count(&count).Offset(offset).Limit(adminConsoleCommentListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get comments failed: " + err.Error())
	}

	pagination = util.NewPagination(page, adminConsoleCommentListPageSize, adminConsoleCommentListWindowSize, count)

	return
}

func (srv *commentService) GetRecentComments(size int, blogID uint64) (ret []*model.Comment) {
	if err := db.Model(&model.Comment{}).Select("`id`, `created_at`, `content`, `author_id`, `article_id`, `author_name`, `author_avatar_url`, `author_url`").
		Where("`blog_id` = ?", blogID).
		Order("`created_at` DESC, `id` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get recent comments failed: " + err.Error())
	}

	return
}

func (srv *commentService) GetArticleComments(articleID uint64, page int, blogID uint64) (ret []*model.Comment, pagination *util.Pagination) {
	offset := (page - 1) * themeCommentListPageSize
	count := 0
	if err := db.Model(&model.Comment{}).Order("`id` ASC").
		Where("`article_id` = ? AND `blog_id` = ?", articleID, blogID).
		Count(&count).Offset(offset).Limit(themeCommentListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get comments failed: " + err.Error())
	}

	pagination = util.NewPagination(page, themeCommentListPageSize, themeCommentListWindowSize, count)

	return
}

func (srv *commentService) AddComment(comment *model.Comment) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment.ID = util.CurrentMillisecond()
	comment.PushedAt = model.ZeroPushTime
	tx := db.Begin()
	if err := tx.Create(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	article := &model.Article{}
	if err := tx.First(article, comment.ArticleID).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Model(article).Update("comment_count", article.CommentCount+1).Error; nil != err {
		tx.Rollback()

		return err
	}
	Statistic.IncCommentCountWithoutTx(tx, comment.BlogID)
	tx.Commit()

	return nil
}

func (srv *commentService) RemoveComment(id, blogID uint64) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	comment := &model.Comment{}

	tx := db.Begin()
	if err := tx.Where("`id` = ? AND `blog_id` = ?", id, blogID).Find(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Delete(comment).Error; nil != err {
		tx.Rollback()

		return err
	}
	article := &model.Article{}
	if err := tx.First(article, comment.ArticleID).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Model(article).Update("comment_count", article.CommentCount-1).Error; nil != err {
		tx.Rollback()

		return err
	}
	Statistic.DecCommentCountWithoutTx(tx, comment.BlogID)
	tx.Commit()

	return nil
}
