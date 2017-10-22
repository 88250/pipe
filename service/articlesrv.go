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

// Package service is the "business logic" layer, encapsulates transaction.
package service

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
	"github.com/jinzhu/gorm"
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
	adminConsoleArticleListPageSize   = 15
	adminConsoleArticleListWindowSize = 20
)

func (srv *articleService) GetPreviousArticle(id uint, blogID uint) *model.Article {
	ret := &model.Article{}
	if err := db.Where("id < ? AND blog_id = ?", id, blogID).Limit(1).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) GetNextArticle(id uint, blogID uint) *model.Article {
	ret := &model.Article{}
	if err := db.Where("id > ? AND blog_id = ?", id, blogID).Limit(1).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) GetArticleByPath(path string) *model.Article {
	ret := &model.Article{}
	if err := db.Where("path = ?", path).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) AddArticle(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tagStr := normalizeTagStr(article.Tags)
	if "" == tagStr {
		return errors.New("invalid tags [" + article.Tags + "]")
	}
	article.Tags = tagStr

	tx := db.Begin()
	if err := tx.Create(article).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tagArticle(tx, article); nil != err {
		tx.Rollback()

		return err
	}
	author := &model.User{}
	if err := tx.First(author, article.AuthorID).Error; nil != err {
		return err
	}
	author.ArticleCount = author.ArticleCount + 1
	if err := tx.Model(author).Updates(author).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := Statistic.IncArticleCountWithoutTx(tx, article.BlogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := Statistic.IncPublishedArticleCountWithoutTx(tx, article.BlogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *articleService) ConsoleGetArticles(page int, blogID uint) (ret []*model.Article, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleArticleListPageSize
	count := 0
	if err := db.Model(model.Article{}).Select("id, created_at, author_id, title, tags, path, topped, view_count, comment_count").
		Where(model.Article{Status: model.ArticleStatusPublished, BlogID: blogID}).
		Order("topped DESC, id DESC").Count(&count).
		Offset(offset).Limit(adminConsoleArticleListPageSize).
		Find(&ret).Error; nil != err {
		log.Errorf("get articles failed: " + err.Error())
	}

	pageCount := int(math.Ceil(float64(count) / adminConsoleArticleListPageSize))
	pagination = util.NewPagination(page, adminConsoleArticleListPageSize, pageCount, adminConsoleArticleListWindowSize, count)

	return
}

func (srv *articleService) GetArticles(page int, blogID uint) (ret []*model.Article, pagination *util.Pagination) {
	settings := Setting.GetSettings(blogID, model.SettingCategoryPreference, []string{model.SettingNamePreferenceArticleListPageSize, model.SettingNamePreferenceArticleListWindowSize})
	pageSize, err := strconv.Atoi(settings[model.SettingNamePreferenceArticleListPageSize].Value)
	if nil != err {
		log.Errorf("value of setting [%s] is not an integer, actual is [%v]", model.SettingNamePreferenceArticleListPageSize, settings[model.SettingNamePreferenceArticleListPageSize].Value)
		pageSize = adminConsoleArticleListPageSize
	}

	offset := (page - 1) * pageSize
	count := 0
	if err := db.Model(model.Article{}).Select("id, created_at, author_id, title, tags, path, topped, view_count, comment_count").
		Where(model.Article{Status: model.ArticleStatusPublished, BlogID: blogID}).
		Order("topped DESC, id DESC").Count(&count).
		Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		log.Errorf("get articles failed: " + err.Error())
	}

	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))
	windowSize, err := strconv.Atoi(settings[model.SettingNamePreferenceArticleListWindowSize].Value)
	if nil != err {
		log.Errorf("value of setting [%s] is not an integer, actual is [%v]", model.SettingNamePreferenceArticleListWindowSize, settings[model.SettingNamePreferenceArticleListWindowSize].Value)
		windowSize = adminConsoleArticleListWindowSize
	}
	pagination = util.NewPagination(page, pageSize, pageCount, windowSize, count)

	return
}

func (srv *articleService) GetMostViewArticles(size int, blogID uint) (ret []*model.Article) {
	if err := db.Model(model.Article{}).Select("id, created_at, author_id, title, path").
		Where(model.Article{Status: model.ArticleStatusPublished, BlogID: blogID}).
		Order("view_count DESC, id DESC").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get most view articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) GetMostCommentArticles(size int, blogID uint) (ret []*model.Article) {
	if err := db.Model(model.Article{}).Select("id, created_at, author_id, title, path").
		Where(model.Article{Status: model.ArticleStatusPublished, BlogID: blogID}).
		Order("comment_count DESC, id DESC").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get most comment articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) GetRandomArticles(size int, blogID uint) (ret []*model.Article) {
	if err := db.Model(model.Article{}).Select("id, created_at, author_id, title, path").
		Where(model.Article{Status: model.ArticleStatusPublished, BlogID: blogID}).
		Order("RANDOM()").Limit(size).Find(&ret).Error; nil != err {
		log.Errorf("get random articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) ConsoleGetArticle(id uint) *model.Article {
	ret := &model.Article{}
	if err := db.First(ret, id).Error; nil != err {
		log.Errorf("get article failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *articleService) RemoveArticle(id uint) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	article := &model.Article{}

	tx := db.Begin()
	if err := tx.First(article, id).Error; nil != err {
		return err
	}
	author := &model.User{}
	if err := tx.First(author, article.AuthorID).Error; nil != err {
		return err
	}
	author.ArticleCount = author.ArticleCount - 1
	author.PublishedArticleCount = author.PublishedArticleCount - 1
	if err := tx.Model(author).Updates(author).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Delete(article).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Where("id1 = ? AND type = ?", article.ID, model.CorrelationArticleTag).Delete(model.Correlation{}).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := Statistic.DecArticleCountWithoutTx(tx, author.BlogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := Statistic.DecPublishedArticleCountWithoutTx(tx, author.BlogID); nil != err {
		tx.Rollback()

		return err
	}
	comments := []*model.Comment{}
	if err := tx.Model(&model.Comment{}).Where("article_id = ?", id).Find(&comments).Error; nil != err {
		tx.Rollback()

		return err
	}
	if 0 < len(comments) {
		if err := tx.Where("article_id = ?", id).Delete(&model.Comment{}).Error; nil != err {
			tx.Rollback()

			return err
		}
		for _, _ = range comments {
			Statistic.DecCommentCountWithoutTx(tx, author.BlogID)
		}
	}
	tx.Commit()

	return nil
}

func (srv *articleService) UpdateArticle(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	count := 0
	if db.Model(&model.Article{}).Where("id = ?", article.ID).Count(&count); 1 > count {
		return errors.New(fmt.Sprintf("not found article [id=%d] to update", article.ID))
	}

	tx := db.Begin()
	if err := tx.Model(article).Updates(article).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func normalizeTagStr(tagStr string) string {
	reg := regexp.MustCompile(`\s+`)
	tagStr = reg.ReplaceAllString(tagStr, "")
	tagStr = strings.Replace(tagStr, "，", ",", -1)
	tagStr = strings.Replace(tagStr, "、", ",", -1)
	tagStr = strings.Replace(tagStr, "；", ",", -1)
	tagStr = strings.Replace(tagStr, ";", ",", -1)

	reg = regexp.MustCompile(`[\\u4e00-\\u9fa5,\\w,&,\\+,-,\\.]+`)
	tags := strings.Split(tagStr, ",")
	retTags := []string{}
	for _, tag := range tags {
		if contains(retTags, tag) {
			continue
		}

		if !reg.MatchString(tag) {
			continue
		}

		retTags = append(retTags, tag)
	}

	return strings.Join(retTags, ",")
}

func tagArticle(tx *gorm.DB, article *model.Article) error {
	tags := strings.Split(article.Tags, ",")
	for _, tagTitle := range tags {
		tag := &model.Tag{BlogID: article.BlogID}
		tx.Where("title = ?", tagTitle).First(tag)
		if "" == tag.Title {
			tag.Title = tagTitle
			tag.ArticleCount = 1
			tag.PublishedArticleCount = 1
			if err := tx.Create(tag).Error; nil != err {
				return err
			}
		} else {
			tag.ArticleCount = tag.ArticleCount + 1
			tag.PublishedArticleCount = tag.PublishedArticleCount + 1
			if err := tx.Model(tag).Updates(tag).Error; nil != err {
				return err
			}
		}

		rel := &model.Correlation{
			ID1:    article.ID,
			ID2:    tag.ID,
			Type:   model.CorrelationArticleTag,
			BlogID: article.BlogID,
		}
		if err := tx.Create(rel).Error; nil != err {
			return err
		}
	}

	return nil
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}

	return false
}
