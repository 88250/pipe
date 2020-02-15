// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package cron

import (
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/dustin/go-humanize"
)

// RecommendArticles saves all recommend articles.
var RecommendArticles []*model.ThemeArticle

func refreshRecommendArticlesPeriodically() {
	go refreshRecommendArticles()

	go func() {
		for range time.Tick(time.Minute * 30) {
			refreshRecommendArticles()
		}
	}()
}

func refreshRecommendArticles() {
	defer gulu.Panic.Recover(nil)

	size := 7
	articles := service.Article.GetPlatMostViewArticles(size)
	size = len(articles)
	indics := gulu.Rand.Ints(0, size, size)
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
			Title:        article.Title,
			URL:          blogURL + article.Path,
			CreatedAt:    humanize.Time(article.CreatedAt),
			Author:       author,
			CommentCount: article.CommentCount,
			ViewCount:    article.ViewCount,
			ThumbnailURL: util.ImageSize(images[i], 280, 90),
		}
		recommendations = append(recommendations, themeArticle)
	}

	RecommendArticles = recommendations
}
