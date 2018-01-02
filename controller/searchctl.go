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

package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/vinta/pangu"
)

func searchAction(c *gin.Context) {
	blogID := getBlogID(c)
	key := c.Query("key")
	page := util.GetPage(c)
	articleModels, pagination := service.Article.GetArticles(key, page, blogID)
	var articles []*model.ThemeArticle
	for _, articleModel := range articleModels {
		var themeTags []*model.ThemeTag
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &model.ThemeTag{
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

		mdResult := util.Markdown(articleModel.Content)
		article := &model.ThemeArticle{
			Title:    pangu.SpacingText(articleModel.Title),
			Abstract: template.HTML(mdResult.AbstractText),
			URL:      getBlogURL(c) + articleModel.Path,
			Tags:     themeTags,
		}

		articles = append(articles, article)
	}

	dataModel := getDataModel(c)
	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	dataModel["Key"] = key
	c.HTML(http.StatusOK, "search.html", dataModel)
}
