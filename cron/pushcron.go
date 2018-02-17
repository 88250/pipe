// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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

package cron

import (
	"net/url"
	"time"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
)

func pushArticlesPeriodically() {
	go pushArticles()

	go func() {
		for range time.Tick(time.Second * 30) {
			pushArticles()
		}
	}()
}

func pushArticles() {
	defer util.Recover()

	server, _ := url.Parse(util.Conf.Server)
	if !util.IsDomain(server.Hostname()) {
		return
	}

	articles := service.Article.GetUnpushedArticles()
	for _, article := range articles {
		author := service.User.GetUser(article.AuthorID)
		b3Key := author.B3Key
		b3Name := author.Name
		if "" == b3Key && !strings.Contains(util.Conf.Server, "pipe.b3log.org") {
			pa := service.User.GetPlatformAdmin()
			b3Key = pa.B3Key
			b3Name = pa.Name
		}
		if "" == b3Key {
			continue
		}

		blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, article.BlogID)
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, article.BlogID)
		requestJSON := map[string]interface{}{
			"article": map[string]interface{}{
				"id":        article.ID,
				"title":     article.Title,
				"permalink": article.Path,
				"tags":      article.Tags,
				"content":   article.Content,
			},
			"client": map[string]interface{}{
				"name":  "Pipe",
				"ver":   util.Version,
				"title": blogTitleSetting.Value,
				"host":  blogURLSetting.Value,
				"email": b3Name,
				"key":   b3Key,
			},
		}
		result := &map[string]interface{}{}
		_, _, errs := gorequest.New().Post("https://rhythm.b3log.org/api/article").SendMap(requestJSON).
			Set("user-agent", util.UserAgent).Timeout(30*time.Second).
			Retry(3, 5*time.Second, http.StatusInternalServerError).EndStruct(result)
		if nil != errs {
			logger.Errorf("push article to Rhythm failed: " + errs[0].Error())
		}

		article.PushedAt = article.UpdatedAt
		service.Article.UpdatePushedAt(article)
	}
}

func pushCommentsPeriodically() {
	go pushComments()

	go func() {
		for range time.Tick(time.Second * 30) {
			pushComments()
		}
	}()
}

func pushComments() {
	defer util.Recover()

	server, _ := url.Parse(util.Conf.Server)
	if !util.IsDomain(server.Hostname()) {
		return
	}

	comments := service.Comment.GetUnpushedComments()
	for _, comment := range comments {
		author := service.User.GetUser(comment.AuthorID)
		article := service.Article.ConsoleGetArticle(comment.ArticleID)
		articleAuthor := service.User.GetUser(article.AuthorID)
		b3Key := articleAuthor.B3Key
		b3Name := articleAuthor.Name
		if "" == b3Key {
			pa := service.User.GetPlatformAdmin()
			b3Key = pa.B3Key
			b3Name = pa.Name
		}
		if "" == b3Key {
			continue
		}

		blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, comment.BlogID)
		requestJSON := map[string]interface{}{
			"comment": map[string]interface{}{
				"id":          comment.ID,
				"articleId":   comment.ArticleID,
				"content":     comment.Content,
				"authorName":  author.Name,
				"authorEmail": "",
			},
			"client": map[string]interface{}{
				"title": blogTitleSetting.Value,
				"host":  util.Conf.Server,
				"email": b3Name,
				"key":   b3Key,
			},
		}
		result := &map[string]interface{}{}
		_, _, errs := gorequest.New().Post("https://rhythm.b3log.org/api/comment").SendMap(requestJSON).
			Set("user-agent", util.UserAgent).Timeout(30*time.Second).
			Retry(3, 5*time.Second, http.StatusInternalServerError).EndStruct(result)
		if nil != errs {
			logger.Errorf("push comment to Rhythm failed: " + errs[0].Error())
		}

		comment.PushedAt = comment.UpdatedAt
		service.Comment.UpdatePushedAt(comment)
	}
}
