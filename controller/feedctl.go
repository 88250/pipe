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

package controller

import (
	"strconv"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
)

func outputAtomAction(c *gin.Context) {
	feed := generateFeed(c)

	feed.WriteAtom(c.Writer)
}

func outputRSSAction(c *gin.Context) {
	feed := generateFeed(c)

	feed.WriteRss(c.Writer)
}

func generateFeed(c *gin.Context) *feeds.Feed {
	blogID := getBlogID(c)

	feedOutputModeSetting := service.Setting.GetSetting(model.SettingCategoryFeed, model.SettingNameFeedOutputMode, blogID)
	blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, blogID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, blogID)
	blogSubtitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogSubtitle, blogID)
	ret := &feeds.Feed{
		Title:       blogTitleSetting.Value,
		Link:        &feeds.Link{Href: blogURLSetting.Value},
		Description: blogSubtitleSetting.Value,
	}

	var items []*feeds.Item
	articles, _ := service.Article.GetArticles("", 1, blogID)
	for _, article := range articles {
		mdResult := util.Markdown(article.Content)
		description := mdResult.AbstractText
		if strconv.Itoa(model.SettingFeedOutputModeValueFull) == feedOutputModeSetting.Value {
			description = mdResult.ContentHTML
		}
		user := service.User.GetUser(article.AuthorID)
		items = append(items, &feeds.Item{
			Title:       article.Title,
			Link:        &feeds.Link{Href: blogURLSetting.Value + article.Path},
			Description: description,
			Author:      &feeds.Author{Name: user.Name},
			Created:     article.CreatedAt,
		})
	}
	ret.Items = items

	return ret
}
