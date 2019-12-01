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

package cron

import (
	"crypto/tls"
	"net/url"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/parnurzeal/gorequest"
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
	defer gulu.Panic.Recover(nil)

	server, _ := url.Parse(model.Conf.Server)
	if !util.IsDomain(server.Hostname()) {
		return
	}

	articles := service.Article.GetUnpushedArticles()
	for _, article := range articles {
		service.Article.ConsolePushArticle(article)
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
	defer gulu.Panic.Recover(nil)

	server, _ := url.Parse(model.Conf.Server)
	if !util.IsDomain(server.Hostname()) {
		return
	}

	comments := service.Comment.GetUnpushedComments()
	for _, comment := range comments {
		article := service.Article.ConsoleGetArticle(comment.ArticleID)
		articleAuthor := service.User.GetUser(article.AuthorID)
		b3Key := articleAuthor.B3Key
		b3Name := articleAuthor.Name
		if "" == b3Key {
			continue
		}

		author := service.User.GetUser(comment.AuthorID)
		blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, comment.BlogID)
		requestJSON := map[string]interface{}{
			"comment": map[string]interface{}{
				"id":         comment.ID,
				"articleId":  comment.ArticleID,
				"content":    comment.Content,
				"authorName": author.Name,
			},
			"client": map[string]interface{}{
				"name":      "Pipe",
				"ver":       model.Version,
				"title":     blogTitleSetting.Value,
				"host":      model.Conf.Server,
				"userName":  b3Name,
				"userB3Key": b3Key,
			},
		}
		result := &map[string]interface{}{}
		_, _, errs := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			Post("https://rhythm.b3log.org/api/comment").SendMap(requestJSON).
			Set("user-agent", model.UserAgent).Timeout(30*time.Second).
			Retry(3, 5*time.Second).EndStruct(result)
		if nil != errs {
			logger.Errorf("push a comment to Rhy failed: " + errs[0].Error())
		} else {
			logger.Infof("push a comment to Rhy result: %+v", result)
		}
		comment.PushedAt = comment.UpdatedAt

		service.Comment.UpdatePushedAt(comment)
	}
}
