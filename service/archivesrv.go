// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
	if err := db.Where("blog_id = ? AND article_count > 0", blogID).Order("year DESC, month DESC").Find(&ret).Error; nil != err {
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
	if err := tx.Where(" year  = ? AND  month  = ? AND  blog_id  = ?",
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

	rel := &model.Correlation{ID1: article.ID, ID2: archive.ID, Type: model.CorrelationArticleArchive, BlogID: article.BlogID}
	if err := tx.Where(rel).Delete(rel).Error; nil != err {
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
	if err := tx.Where(archive).First(archive).Error; nil != err {
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
	ret := &model.Archive{Year: year, Month: month, BlogID: blogID}
	if err := db.Where(ret).First(ret).Error; nil != err {
		return nil
	}

	return ret
}
