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
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
)

func outputAtomAction(c *gin.Context) {
	blogAdmin := getBlogAdmin(c)

	blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, blogAdmin.BlogID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, blogAdmin.BlogID)
	blogSubtitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogSubtitle, blogAdmin.BlogID)
	feed := &feeds.Feed{
		Title:       blogTitleSetting.Value,
		Link:        &feeds.Link{Href: blogURLSetting.Value},
		Description: blogSubtitleSetting.Value,
		Author:      &feeds.Author{Name: blogAdmin.Name},
	}

	items := []*feeds.Item{}
	articles, _ := service.Article.GetArticles(1, blogAdmin.BlogID)
	for _, article := range articles {
		user := service.User.GetUser(article.AuthorID)
		items = append(items, &feeds.Item{
			Title:       article.Title,
			Link:        &feeds.Link{Href: blogURLSetting.Value + article.Path},
			Description: util.MarkdownAbstract(article.Content),
			Author:      &feeds.Author{Name: user.Name},
			Created:     article.CreatedAt,
		})
	}
	feed.Items = items

	feed.WriteAtom(c.Writer)
}
