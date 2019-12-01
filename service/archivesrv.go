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

// Package service is the "business logic" layer, encapsulates transaction.
package service

import (
	"sync"

	"github.com/88250/pipe/model"
	"github.com/jinzhu/gorm"
)

// Archive service.
var Archive = &archiveService{
	mutex: &sync.Mutex{},
}

type archiveService struct {
	mutex *sync.Mutex
}

func (srv *archiveService) GetArchives(blogID uint64) []*model.Archive {
	var ret []*model.Archive
	if err := db.Where("`blog_id` = ? AND `article_count` > 0", blogID).Order("`year` DESC, `month` DESC").Find(&ret).Error; nil != err {
		logger.Error("get archives failed: " + err.Error())
	}

	return ret
}

func (srv *archiveService) UnArchiveArticleWithoutTx(tx *gorm.DB, article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	year := article.CreatedAt.Format("2006")
	month := article.CreatedAt.Format("01")
	archive := &model.Archive{Year: year, Month: month, BlogID: article.BlogID}
	if err := tx.Where("`year` = ? AND `month` = ? AND `blog_id` = ?",
		year, month, article.BlogID).First(archive).Error; nil != err {
		return err
	}
	archive.ArticleCount--
	if archive.ArticleCount < 0 {
		logger.Error("impossible: article count < 0")
		archive.ArticleCount = 0
	}
	if err := tx.Save(archive).Error; nil != err {
		return err
	}
	if err := tx.Where("`id1` = ? AND `id2` = ? AND `type` = ? AND `blog_id` = ?",
		article.ID, archive.ID, model.CorrelationArticleArchive, article.BlogID).
		Delete(&model.Correlation{}).Error; nil != err {
		return err
	}

	return nil
}

func (srv *archiveService) ArchiveArticleWithoutTx(tx *gorm.DB, article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	year := article.CreatedAt.Format("2006")
	month := article.CreatedAt.Format("01")

	archive := &model.Archive{Year: year, Month: month, BlogID: article.BlogID}
	if err := tx.Where("`year` = ? AND `month` = ? AND `blog_id` = ?",
		year, month, article.BlogID).First(archive).Error; nil != err {
		if gorm.ErrRecordNotFound != err {
			return err
		}
	}
	archive.ArticleCount++
	if err := tx.Save(archive).Error; nil != err {
		return err
	}

	articleArchiveRel := &model.Correlation{
		ID1:    article.ID,
		ID2:    archive.ID,
		Type:   model.CorrelationArticleArchive,
		BlogID: article.BlogID,
	}
	if err := tx.Create(articleArchiveRel).Error; nil != err {
		return err
	}

	return nil
}

func (srv *archiveService) GetArchive(year, month string, blogID uint64) *model.Archive {
	ret := &model.Archive{}
	if err := db.Where("`year` = ? AND `month` = ? AND `blog_id` = ?",
		year, month, blogID).First(ret).Error; nil != err {
		return nil
	}

	return ret
}
