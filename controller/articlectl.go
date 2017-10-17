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

// Package controller is the "controller" layer.
package controller

import (
	"net/http"
	"strings"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ThemeListArticle struct {
	ID           uint
	Abstract     string
	Author       *ThemeAuthor
	CreatedAt    string
	Title        string
	Tags         []*ThemeTag
	URL          string
	Topped       bool
	ViewCount    int
	CommentCount int
	ThumbnailURL string
	Content  string
}

type ThemeTag struct {
	Title string
	URL   string
}

type ThemeAuthor struct {
	Name      string
	AvatarURL string
	URL       string
}

func showArticlesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	page := c.GetInt("p")
	if 1 > page {
		page = 1
	}

	blogAdminVal, _ := c.Get("blogAdmin")
	blogAdmin := blogAdminVal.(*model.User)

	articleModels, pagination := service.Article.GetArticles(page, blogAdmin.BlogID)
	articles := []*ThemeListArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   dataModel["Setting"].(map[string]string)[model.SettingNameSystemPath] + util.PathTags + "/" + tagStr,
			}
			themeTags = append(themeTags, themeTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		if nil == authorModel {
			log.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

			continue
		}

		author := &ThemeAuthor{
			Name:      authorModel.Name,
			URL: "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}

		article := &ThemeListArticle{
			ID:           articleModel.ID,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         themeTags,
			URL:          dataModel["Setting"].(map[string]string)[model.SettingNameSystemPath] + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL:  "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}

		articles = append(articles, article)
	}

	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	c.HTML(http.StatusOK, "index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	page := c.GetInt("p")
    	if 1 > page {
    		page = 1
    	}

    	blogAdminVal, _ := c.Get("blogAdmin")
    	blogAdmin := blogAdminVal.(*model.User)

    	articleModels, pagination := service.Article.GetArticles(page, blogAdmin.BlogID)
    	articles := []*ThemeListArticle{}
    		themeTags := []*ThemeTag{}
    	for _, articleModel := range articleModels {
    		tagStrs := strings.Split(articleModel.Tags, ",")
    		for _, tagStr := range tagStrs {
    			themeTag := &ThemeTag{
    				Title: tagStr,
    				URL:   dataModel["Setting"].(map[string]string)[model.SettingNameSystemPath] + util.PathTags + "/" + tagStr,
    			}
    			themeTags = append(themeTags, themeTag)
    		}

    		authorModel := service.User.GetUser(articleModel.AuthorID)
    		if nil == authorModel {
    			log.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

    			continue
    		}

    		author := &ThemeAuthor{
    			Name:      authorModel.Name,
    			URL: "http://localhost:5879/blogs/solo/vanessa",
    			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
    		}

    		article := &ThemeListArticle{
    			ID:           articleModel.ID,
    			Author:       author,
    			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
    			Title:        articleModel.Title,
    			Tags:         themeTags,
    			URL:          dataModel["Setting"].(map[string]string)[model.SettingNameSystemPath] + articleModel.Path,
    			Topped:       articleModel.Topped,
    			ViewCount:    articleModel.ViewCount,
    			CommentCount: articleModel.CommentCount,
    			ThumbnailURL:  "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
    		}

    		articles = append(articles, article)
    	}

    	dataModel["Article"] =  &ThemeListArticle{
                                   			Author:       &ThemeAuthor{
                                                              			Name:      "Vanessa",
                                                              			URL: "http://localhost:5879/blogs/solo/vanessa",
                                                              			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
                                                              		},
                                   			CreatedAt:    "2015-12-12",
                                   			Title:        "Title",
                                   			Tags:         themeTags,
                                   			URL:          "url",
                                   			Topped:       true,
                                   			ViewCount:    1,
                                   			CommentCount: 1,
                                   			Content: "sfasdfsf",
                                   		}
    	dataModel["Comments"] = articles
    	dataModel["Pagination"] = pagination

	dataModel["RelevantArticles"] = articles
	dataModel["ExternalRelevantArticles"] = articles
	c.HTML(http.StatusOK, "article.html", dataModel)
}
