// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
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
	"html/template"
	"time"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/dustin/go-humanize"
	"github.com/parnurzeal/gorequest"
)

// RecommendArticles saves all recommend articles.
var RecommendArticles []*model.ThemeArticle

// CommunityRecommendArticles saves all community recommend articles.
var CommunityRecommendArticles []*model.ThemeArticle

func refreshRecommendArticlesPeriodically() {
	go refreshRecommendArticles()
	go refreshCommunityRecommendArticlesPeriodically()

	go func() {
		for range time.Tick(time.Minute * 30) {
			refreshRecommendArticles()
			refreshCommunityRecommendArticlesPeriodically()
		}
	}()
}

func refreshRecommendArticles() {
	defer util.Recover()

	size := 7
	articles := service.Article.GetPlatMostViewArticles(size)
	size = len(articles)
	indics := util.RandInts(0, size, size)
	images := util.RandImages(size)
	indics = indics[:len(images)]
	var recommendations []*model.ThemeArticle
	for i, index := range indics {
		article := articles[index]
		authorModel := service.User.GetUser(article.AuthorID)
		if nil == authorModel {
			logger.Errorf("not found author of article [id=%d, authorID=%d]", article.ID, article.AuthorID)

			continue
		}

		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, article.BlogID)
		blogURL := blogURLSetting.Value
		author := &model.ThemeAuthor{
			Name:      authorModel.Name,
			URL:       blogURL + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}
		themeArticle := &model.ThemeArticle{
			Title:     article.Title,
			URL:       blogURL + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
			CommentCount: article.CommentCount,
			ViewCount:    article.ViewCount,
			ThumbnailURL: util.ImageSize(images[i], 280, 90),
		}
		recommendations = append(recommendations, themeArticle)
	}

	RecommendArticles = recommendations
}

func refreshCommunityRecommendArticlesPeriodically() {
	go refreshCommunityRecommendArticles()

	go func() {
		for range time.Tick(time.Minute * 30) {
			refreshCommunityRecommendArticles()
		}
	}()
}

func refreshCommunityRecommendArticles() {
	defer util.Recover()

	result := util.NewResult()
	_, _, errs := gorequest.New().Get(util.HacPaiURL+"/apis/recommend/articles").
		Set("user-agent", model.UserAgent).Timeout(30*time.Second).
		Retry(3, 5*time.Second).EndStruct(result)
	if nil != errs {
		logger.Errorf("get recommend articles: %s", errs)

		return
	}
	if 0 != result.Code {
		return
	}

	size := 30
	entries := result.Data.([]interface{})
	if size > len(entries) {
		size = len(entries)
	}

	indics := util.RandInts(0, len(entries), size)
	images := util.RandImages(size)
	indics = indics[:len(images)]
	var recommendations []*model.ThemeArticle
	for i, index := range indics {
		article := entries[index].(map[string]interface{})
		author := &model.ThemeAuthor{
			Name:      article["articleAuthorName"].(string),
			URL:       "https://hacpai.com/member/" + article["articleAuthorName"].(string),
			AvatarURL: article["articleAuthorThumbnailURL"].(string),
		}

		recommendations = append(recommendations, &model.ThemeArticle{
			Author:       author,
			Abstract:     template.HTML(article["articlePreviewContent"].(string)),
			CreatedAt:    time.Unix(int64(article["articleCreateTime"].(float64)/1000), 0).Format("2006-01-02"),
			Title:        article["articleTitle"].(string),
			URL:          article["articlePermalink"].(string),
			CommentCount: int(article["articleCommentCount"].(float64)),
			ViewCount:    int(article["articleViewCount"].(float64)),
			ThumbnailURL: util.ImageSize(images[i], 280, 90),
		})
	}

	CommunityRecommendArticles = recommendations
}
