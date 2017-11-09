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

package controller

import (
	"net/http"
	"strings"

	"github.com/b3log/pipe/i18n"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func showArchivesAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)
	locale := getLocale(c)
	themeArchives := []*ThemeArchive{}
	archiveModels := service.Archive.GetArchives(blogAdmin.BlogID)
	for _, archiveModel := range archiveModels {
		archive := &ThemeArchive{
			Title:        i18n.GetMessagef(locale, "archiveYearMonth", archiveModel.Year, archiveModel.Month),
			URL:          getBlogURL(c) + util.PathArchives + "/" + archiveModel.Year + "/" + archiveModel.Month,
			ArticleCount: archiveModel.ArticleCount,
		}
		themeArchives = append(themeArchives, archive)
	}

	dataModel["Archives"] = themeArchives
	c.HTML(http.StatusOK, getTheme(c)+"/archives.html", dataModel)
}

func showArchiveArticlesAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)
	locale := getLocale(c)
	session := util.GetSession(c)
	date := strings.SplitAfter(c.Request.URL.Path, util.PathArchives+"/")[1]
	year := strings.Split(date, "/")[0]
	month := strings.Split(date, "/")[1]
	archiveModel := service.Archive.GetArchive(year, month, blogAdmin.BlogID)
	if nil == archiveModel {
		c.Status(http.StatusNotFound)

		return
	}
	articleModels, pagination := service.Article.GetArchiveArticles(archiveModel.ID, util.GetPage(c), blogAdmin.BlogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
			}
			themeTags = append(themeTags, themeTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		if nil == authorModel {
			logger.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

			continue
		}

		author := &ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		mdResult := util.Markdown(articleModel.Content)
		article := &ThemeArticle{
			ID:           articleModel.ID,
			Abstract:     mdResult.AbstractText,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         themeTags,
			URL:          getBlogURL(c) + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL: mdResult.ThumbURL,
			Editable:     session.UID == authorModel.ID,
		}

		articles = append(articles, article)
	}
	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	dataModel["Archive"] = &ThemeArchive{
		Title:        i18n.GetMessagef(locale, "archiveYearMonth", archiveModel.Year, archiveModel.Month),
		URL:          getBlogURL(c) + util.PathArchives + "/" + archiveModel.Year + "/" + archiveModel.Month,
		ArticleCount: archiveModel.ArticleCount,
	}

	c.HTML(http.StatusOK, getTheme(c)+"/archive-articles.html", dataModel)
}
