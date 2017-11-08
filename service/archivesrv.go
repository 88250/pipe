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
	"sync"

	"github.com/b3log/pipe/model"
	"github.com/jinzhu/gorm"
)

var Archive = &archiveService{
	mutex: &sync.Mutex{},
}

type archiveService struct {
	mutex *sync.Mutex
}

func (srv *archiveService) UnarchiveArticleWithoutTx(tx *gorm.DB, article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	year := article.CreatedAt.Format("2006")
	month := article.CreatedAt.Format("01")

	archive, err := srv.GetArchive(year, month, article.BlogID)
	if nil != err {
		return err
	}
	if nil == archive {
		return nil
	}
	archive.ArticleCount -= 1
	if err := tx.Save(archive).Error; nil != err {
		return err
	}
	if err := tx.Where(&model.Correlation{ID1: article.ID, ID2: archive.ID, Type: model.CorrelationArticleArchive, BlogID: article.BlogID}).
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

	archive, err := srv.GetArchive(year, month, article.BlogID)
	if nil != err {
		return err
	}
	if nil == archive {
		archive = &model.Archive{Year: year, Month: month, BlogID: article.BlogID}
	}
	archive.ArticleCount += 1
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

func (srv *archiveService) GetArchive(year, month string, blogID uint) (*model.Archive, error) {
	ret := &model.Archive{Year: year, Month: month, BlogID: blogID}
	if err := db.Where(ret).First(ret).Error; nil != err {
		if gorm.ErrRecordNotFound != err {
			return nil, err
		}

		return nil, nil
	}

	return ret, nil
}
