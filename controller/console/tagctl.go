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

package console

import (
	"net/http"
	"strings"

	"github.com/b3log/solo.go/service"
	"github.com/b3log/wide/util"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

func GetTagsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	articleModels, pagination := service.Article.ConsoleGetArticles(c.GetInt("p"))

	articles := []ConsoleArticle{}
	for _, articleModel := range articleModels {
		tagPermalinks := []*TagPermalink{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			tagPermalink := &TagPermalink{
				Title:     tagStr,
				Permalink: "context/" + tagStr,
			}
			tagPermalinks = append(tagPermalinks, tagPermalink)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		if nil == authorModel {
			log.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

			continue
		}

		author := &Author{
			Name:      authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		article := ConsoleArticle{
			ID:           articleModel.ID,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         tagPermalinks,
			Permalink:    articleModel.Permalink,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
		}

		articles = append(articles, article)
	}

	data := map[string]interface{}{}
	data["articles"] = articles
	data["pagination"] = pagination
	result.Data = data
}
