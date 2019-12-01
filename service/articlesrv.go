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
	"crypto/tls"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
	"github.com/jinzhu/gorm"
	"github.com/parnurzeal/gorequest"
)

// Article service.
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

func (srv *articleService) GetPlatMostViewArticles(size int) (ret []*model.Article) {
	if err := db.Model(&model.Article{}).Select("`id`, `created_at`, `author_id`, `title`, `path`, `view_count`, `comment_count`, `blog_id`").
		Where("`status` = ?", model.ArticleStatusOK).
		Order("`view_count` DESC, `created_at` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get platform most view articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) GetUnpushedArticles() (ret []*model.Article) {
	if err := db.Where("`pushed_at` <= ?", model.ZeroPushTime).Find(&ret).Error; nil != err {
		return
	}

	return
}

func (srv *articleService) GetArchiveArticles(archiveID uint64, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	pageSize, windowSize := getPageWindowSize(blogID)
	offset := (page - 1) * pageSize
	count := 0

	var rels []*model.Correlation
	if err := db.Where("`id2` = ? AND `type` = ? AND `blog_id` = ?", archiveID, model.CorrelationArticleArchive, blogID).
		Find(&rels).Error; nil != err {
		return
	}

	var articleIDs []uint64
	for _, articleTagRel := range rels {
		articleIDs = append(articleIDs, articleTagRel.ID1)
	}

	if err := db.Model(&model.Article{}).
		Where("`id` IN (?) AND `status` = ? AND `blog_id` = ?", articleIDs, model.ArticleStatusOK, blogID).
		Order("`topped` DESC, `created_at` DESC").Count(&count).
		Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		logger.Errorf("get archive articles failed: " + err.Error())
	}

	pagination = util.NewPagination(page, pageSize, windowSize, count)

	return
}

func (srv *articleService) GetPreviousArticle(id uint64, blogID uint64) *model.Article {
	ret := &model.Article{}
	if err := db.Where("`id` < ? AND `blog_id` = ?", id, blogID).Order("`created_at` DESC").Limit(1).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) GetNextArticle(id uint64, blogID uint64) *model.Article {
	ret := &model.Article{}
	if err := db.Where("`id` > ? AND `blog_id` = ?", id, blogID).Limit(1).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) GetArticleByPath(path string, blogID uint64) *model.Article {
	path = strings.TrimSpace(path)
	if "" == path || util.IsReservedPath(path) {
		return nil
	}
	path, _ = url.PathUnescape(path)

	ret := &model.Article{}
	if err := db.Where("`path` = ? AND `blog_id` = ?", path, blogID).Find(ret).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) AddArticle(article *model.Article) (err error) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if article.CreatedAt.IsZero() {
		article.CreatedAt = time.Now()
	}
	if article.CreatedAt != article.PushedAt {
		article.PushedAt = model.ZeroPushTime
	}
	if err := normalizeArticle(article); nil != err {
		return err
	}

	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if err = tx.Create(article).Error; nil != err {
		return
	}
	if err = tagArticle(tx, article); nil != err {
		return
	}
	if err = Archive.ArchiveArticleWithoutTx(tx, article); nil != err {
		return
	}
	author := &model.User{}
	if err = tx.First(author, article.AuthorID).Error; nil != err {
		return
	}
	author.TotalArticleCount += 1
	if err = tx.Model(author).Updates(author).Error; nil != err {
		return
	}
	blogUserRel := &model.Correlation{}
	if err = tx.Where("`id1` = ? AND `id2` = ? AND `type` = ? AND `blog_id` = ?",
		article.BlogID, author.ID, model.CorrelationBlogUser, article.BlogID).First(blogUserRel).Error; nil != err {
		return
	}
	blogUserRel.Int2 += 1
	if err = tx.Model(blogUserRel).Updates(blogUserRel).Error; nil != err {
		return
	}
	if err = Statistic.IncArticleCountWithoutTx(tx, article.BlogID); nil != err {
		return
	}

	return nil // triger commit in the defer
}

func (srv *articleService) ConsoleGetArticles(keyword string, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	offset := (page - 1) * adminConsoleArticleListPageSize
	count := 0

	where := "`status` = ? AND `blog_id` = ?"
	whereArgs := []interface{}{model.ArticleStatusOK, blogID}
	if "" != keyword {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	if err := db.Model(&model.Article{}).Select("`id`, `created_at`, `author_id`, `title`, `tags`, `path`, `topped`, `view_count`, `comment_count`").
		Where(where, whereArgs...).
		Order("`topped` DESC, `created_at` DESC").Count(&count).
		Offset(offset).Limit(adminConsoleArticleListPageSize).Find(&ret).Error; nil != err {
		logger.Errorf("get articles failed: " + err.Error())
	}

	pagination = util.NewPagination(page, adminConsoleArticleListPageSize, adminConsoleArticleListWindowSize, count)

	return
}

func (srv *articleService) GetArticles(keyword string, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	pageSize, windowSize := getPageWindowSize(blogID)
	offset := (page - 1) * pageSize
	count := 0

	where := "`status` = ? AND `blog_id` = ?"
	whereArgs := []interface{}{model.ArticleStatusOK, blogID}
	if "" != keyword {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	if err := db.Model(&model.Article{}).Select("`id`, `created_at`, `author_id`, `title`, `abstract`, `content`, `tags`, `path`, `topped`, `view_count`, `comment_count`").
		Where(where, whereArgs...).
		Order("`topped` DESC, `created_at` DESC").Count(&count).
		Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		logger.Errorf("get articles failed: " + err.Error())
	}

	pagination = util.NewPagination(page, pageSize, windowSize, count)

	return
}

func (srv *articleService) GetCategoryArticles(categoryID uint64, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	pageSize, windowSize := getPageWindowSize(blogID)
	offset := (page - 1) * pageSize

	var rels []*model.Correlation
	if err := db.Model(&model.Correlation{}).Where("`id1` = ? AND `type` = ? AND `blog_id` = ?", categoryID, model.CorrelationCategoryTag, blogID).
		Find(&rels).Error; nil != err {
		return
	}

	var tagIDs []uint64
	for _, categoryTagRel := range rels {
		tagIDs = append(tagIDs, categoryTagRel.ID2)
	}

	count := 0
	rels = []*model.Correlation{}
	if err := db.Model(&model.Correlation{}).Where("`id2` IN (?) AND `type` = ? AND `blog_id` = ?", tagIDs, model.CorrelationArticleTag, blogID).
		Order("`id` DESC").Count(&count).Offset(offset).Limit(pageSize).
		Find(&rels).Error; nil != err {
		return
	}

	pagination = util.NewPagination(page, pageSize, windowSize, count)

	var articleIDs []uint64
	for _, articleTagRel := range rels {
		articleIDs = append(articleIDs, articleTagRel.ID1)
	}

	if err := db.Where("`id` IN (?) AND `blog_id` = ?", articleIDs, blogID).Find(&ret).Error; nil != err {
		return
	}

	return
}

func (srv *articleService) GetTagArticles(tagID uint64, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	pageSize, windowSize := getPageWindowSize(blogID)
	offset := (page - 1) * pageSize
	count := 0

	var rels []*model.Correlation
	if err := db.Where("`id2` = ? AND `type` = ? AND `blog_id` = ?", tagID, model.CorrelationArticleTag, blogID).
		Find(&rels).Error; nil != err {
		return
	}

	var articleIDs []uint64
	for _, articleTagRel := range rels {
		articleIDs = append(articleIDs, articleTagRel.ID1)
	}

	if err := db.Model(&model.Article{}).
		Where("`id` IN (?) AND `status` = ? AND `blog_id` = ?", articleIDs, model.ArticleStatusOK, blogID).
		Order("`topped` DESC, `created_at` DESC").Count(&count).Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		logger.Errorf("get tag articles failed: " + err.Error())
	}

	pagination = util.NewPagination(page, pageSize, windowSize, count)

	return
}

func (srv *articleService) GetAuthorArticles(authorID uint64, page int, blogID uint64) (ret []*model.Article, pagination *util.Pagination) {
	pageSize, windowSize := getPageWindowSize(blogID)
	offset := (page - 1) * pageSize
	count := 0

	if err := db.Model(&model.Article{}).
		Where("`author_id` = ? AND `status` = ? AND `blog_id` = ?", authorID, model.ArticleStatusOK, blogID).
		Order("`topped` DESC, `created_at` DESC").Count(&count).
		Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		logger.Errorf("get author articles failed: " + err.Error())
	}

	pagination = util.NewPagination(page, pageSize, windowSize, count)

	return
}

func (srv *articleService) GetMostViewArticles(size int, blogID uint64) (ret []*model.Article) {
	if err := db.Model(&model.Article{}).Select("`id`, `created_at`, `author_id`, `title`, `path`").
		Where("`status` = ? AND `blog_id` = ?", model.ArticleStatusOK, blogID).
		Order("`view_count` DESC, `created_at` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get most view articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) GetMostCommentArticles(size int, blogID uint64) (ret []*model.Article) {
	if err := db.Model(&model.Article{}).Select("`id`, `created_at`, `author_id`, `title`, `path`").
		Where("`status` = ? AND `blog_id` = ?", model.ArticleStatusOK, blogID).
		Order("`comment_count` DESC, `id` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get most comment articles failed: " + err.Error())
	}

	return
}

func (srv *articleService) ConsolePushArticle(article *model.Article) {
	if nil == article {
		return
	}

	author := User.GetUser(article.AuthorID)
	b3Key := author.B3Key
	b3Name := author.Name
	if "" == b3Key {
		return
	}

	blogTitleSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, article.BlogID)
	blogURLSetting := Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, article.BlogID)
	requestJSON := map[string]interface{}{
		"article": map[string]interface{}{
			"id":        article.ID,
			"title":     article.Title,
			"permalink": article.Path,
			"tags":      article.Tags,
			"content":   article.Content,
		},
		"client": map[string]interface{}{
			"title":     blogTitleSetting.Value,
			"host":      blogURLSetting.Value,
			"name":      "Pipe",
			"ver":       model.Version,
			"userName":  b3Name,
			"userB3Key": b3Key,
		},
	}
	result := gulu.Ret.NewResult()
	_, _, errs := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Post("https://rhythm.b3log.org/api/article").SendMap(requestJSON).
		Set("user-agent", model.UserAgent).Timeout(30*time.Second).
		Retry(3, 5*time.Second).EndStruct(result)
	if nil != errs {
		logger.Debugf("push an article to Rhy failed: " + errs[0].Error())
	} else {
		logger.Infof("push an article to Rhy result: %+v", result)
	}

	article.PushedAt = article.UpdatedAt
	Article.UpdatePushedAt(article)
}

func (srv *articleService) ConsoleGetArticle(id uint64) *model.Article {
	ret := &model.Article{}
	if err := db.First(ret, id).Error; nil != err {
		return nil
	}

	return ret
}

func (srv *articleService) RemoveArticle(id, blogID uint64) (err error) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	article := &model.Article{}

	tx := db.Begin()
	defer func() {
		if nil == err {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if err = tx.Where("`id` = ? AND `blog_id` = ?", id, blogID).Find(article).Error; nil != err {
		return
	}
	author := &model.User{}
	if err = tx.First(author, article.AuthorID).Error; nil != err {
		return
	}
	author.TotalArticleCount -= 1
	if err = tx.Model(author).Updates(author).Error; nil != err {
		return
	}
	blogUserRel := &model.Correlation{}
	if err = tx.Where("`id1` = ? AND `id2` = ? AND `type` = ? AND `blog_id` = ?",
		article.BlogID, author.ID, model.CorrelationBlogUser, article.BlogID).First(blogUserRel).Error; nil != err {
		return
	}
	blogUserRel.Int2 -= 1
	if err = tx.Model(blogUserRel).Updates(blogUserRel).Error; nil != err {
		return
	}
	if err = tx.Delete(article).Error; nil != err {
		return
	}
	if err = removeTagArticleRels(tx, article); nil != err {
		return
	}
	if err = Archive.UnArchiveArticleWithoutTx(tx, article); nil != err {
		return
	}
	if err = Statistic.DecArticleCountWithoutTx(tx, article.BlogID); nil != err {
		return
	}
	var comments []*model.Comment
	if err = tx.Model(&model.Comment{}).Where("`article_id` = ? AND `blog_id` = ?", id, article.BlogID).Find(&comments).Error; nil != err {
		return
	}
	if 0 < len(comments) {
		if err = tx.Where("`article_id` = ? AND `blog_id` = ?", id, article.BlogID).Delete(&model.Comment{}).Error; nil != err {
			return
		}
		for range comments {
			Statistic.DecCommentCountWithoutTx(tx, article.BlogID)
		}
	}
	return nil // trigger commit in the defer
}

func (srv *articleService) UpdatePushedAt(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	article.PushedAt = article.UpdatedAt
	if err := db.Model(article).UpdateColumns(article).Error; nil != err {
		return err
	}

	return nil
}

func (srv *articleService) UpdateArticle(article *model.Article) (err error) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	oldArticle := &model.Article{}
	if err = db.Model(&model.Article{}).Where("`id` = ? AND `blog_id` = ?", article.ID, article.BlogID).
		Find(oldArticle).Error; nil != err {
		return
	}

	oldArticle.Title = strings.TrimSpace(article.Title)
	oldArticle.Abstract = strings.TrimSpace(article.Abstract)
	oldArticle.Content = strings.TrimSpace(article.Content)
	oldArticle.Commentable = article.Commentable
	oldArticle.Topped = article.Topped
	now := time.Now()
	oldArticle.UpdatedAt = now

	if article.PushedAt.IsZero() {
		oldArticle.PushedAt = model.ZeroPushTime
	}

	tagStr := normalizeTagStr(article.Tags)
	oldArticle.Tags = tagStr

	if err = normalizeArticlePath(article); nil != err {
		return
	}
	oldArticle.Path = article.Path

	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if oldArticle.CreatedAt.Format("200601") != article.CreatedAt.Format("200601") {
		// https://github.com/b3log/pipe/issues/106
		if err = Archive.UnArchiveArticleWithoutTx(tx, oldArticle); nil != err {
			return
		}
		oldArticle.CreatedAt = article.CreatedAt
		if err = Archive.ArchiveArticleWithoutTx(tx, oldArticle); nil != err {
			return
		}
	}
	if err = tx.Save(oldArticle).Error; nil != err {
		return
	}
	if err = removeTagArticleRels(tx, article); nil != err {
		return
	}
	if err = tagArticle(tx, article); nil != err {
		return
	}

	return nil
}

func (srv *articleService) IncArticleViewCount(article *model.Article) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	article.ViewCount = article.ViewCount + 1
	if err := db.Model(&model.Article{}).Where("`id` = ?", article.ID).Select("view_count").Updates(article).Error; nil != err {
		return err
	}

	return nil
}

func normalizeArticle(article *model.Article) error {
	title := strings.TrimSpace(article.Title)
	if "" == title {
		return errors.New("title can not be empty")
	}
	count := 0
	if err := db.Model(&model.Article{}).Where("`title` = ?", title).Count(&count).Error; nil != err {
		return err
	}
	if 0 < count {
		return errors.New("title [" + title + "] is reduplicated")
	}

	content := strings.TrimSpace(article.Content)
	if "" == content {
		return errors.New("content can not be empty")
	}
	article.Content = content

	if util.IsReservedPath(article.Path) {
		return errors.New("invalid path [" + article.Path + "]")
	}

	tagStr := normalizeTagStr(article.Tags)
	article.Tags = tagStr

	if 1 > article.ID {
		article.ID = util.CurrentMillisecond()
	}

	if err := normalizeArticlePath(article); nil != err {
		return err
	}

	return nil
}

func normalizeTagStr(tagStr string) string {
	reg := regexp.MustCompile(`\s+`)
	ret := reg.ReplaceAllString(tagStr, "")
	ret = strings.Replace(ret, "，", ",", -1)
	ret = strings.Replace(ret, "、", ",", -1)
	ret = strings.Replace(ret, "；", ",", -1)
	ret = strings.Replace(ret, ";", ",", -1)

	reg = regexp.MustCompile(`[\\u4e00-\\u9fa5,\\w,&,\\+,-,\\.]+`)
	tags := strings.Split(ret, ",")
	var retTags []string
	for _, tag := range tags {
		if contains(retTags, tag) {
			continue
		}

		if !reg.MatchString(tag) {
			continue
		}

		retTags = append(retTags, tag)
	}

	if "" == ret {
		return "待分类"
	}

	return ret
}

func removeTagArticleRels(tx *gorm.DB, article *model.Article) error {
	var rels []*model.Correlation
	if err := tx.Where("`id1` = ? AND `type` = ? AND `blog_id` = ?",
		article.ID, model.CorrelationArticleTag, article.BlogID).Find(&rels).Error; nil != err {
		return err
	}
	for _, rel := range rels {
		tag := &model.Tag{}
		if err := tx.Where("`id` = ? AND `blog_id` = ?", rel.ID2, article.BlogID).First(tag).Error; nil != err {
			continue
		}
		tag.ArticleCount = tag.ArticleCount - 1
		if err := tx.Save(tag).Error; nil != err {
			continue
		}
	}

	if err := tx.Where("`id1` = ? AND `type` = ? AND `blog_id` = ?", article.ID, model.CorrelationArticleTag, article.BlogID).
		Delete(&model.Correlation{}).Error; nil != err {
		return err
	}

	return nil
}

func tagArticle(tx *gorm.DB, article *model.Article) error {
	tags := strings.Split(article.Tags, ",")
	for _, tagTitle := range tags {
		tag := &model.Tag{BlogID: article.BlogID}
		tx.Where("`title` = ? AND `blog_id` = ?", tagTitle, article.BlogID).First(tag)
		if "" == tag.Title {
			tag.Title = tagTitle
			tag.ArticleCount = 1
			tag.BlogID = article.BlogID
			if err := tx.Create(tag).Error; nil != err {
				return err
			}
		} else {
			tag.ArticleCount = tag.ArticleCount + 1
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

func normalizeArticlePath(article *model.Article) error {
	path := strings.TrimSpace(article.Path)
	if "" == path {
		path = util.PathArticles + article.CreatedAt.Format("/2006/01/02/") +
			fmt.Sprintf("%d", article.ID)
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	count := 0
	if db.Model(&model.Article{}).Where("`path` = ? AND `id` != ? AND `blog_id` = ?", path, article.ID, article.BlogID).Count(&count); 0 < count {
		return errors.New("path [" + path + "] is reduplicated")
	}

	article.Path = path

	return nil
}

func getPageWindowSize(blogID uint64) (pageSize, windowSize int) {
	pageSizeSetting := Setting.GetSetting(model.SettingCategoryPreference, model.SettingNamePreferenceArticleListPageSize, blogID)
	pageSize, err := strconv.Atoi(pageSizeSetting.Value)
	if nil != err {
		logger.Errorf("value of setting [%s] is not an integer, actual is [%v]", model.SettingNamePreferenceArticleListPageSize, pageSizeSetting.Value)
		pageSize = adminConsoleArticleListPageSize
	}

	windowSizeSetting := Setting.GetSetting(model.SettingCategoryPreference, model.SettingNamePreferenceArticleListWindowSize, blogID)
	windowSize, err = strconv.Atoi(windowSizeSetting.Value)
	if nil != err {
		logger.Errorf("value of setting [%s] is not an integer, actual is [%v]", model.SettingNamePreferenceArticleListWindowSize, windowSizeSetting.Value)
		windowSize = adminConsoleArticleListWindowSize
	}

	return
}
